// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcore

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	headerOperationLocation = "Operation-Location"
	headerLocation          = "Location"
)

// returns true if the LRO response contains a valid HTTP status code
func lroStatusCodeValid(resp *Response) bool {
	return resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusCreated, http.StatusNoContent)
}

// ErrorUnmarshaller is the func to invoke when the endpoint returns an error response that requires unmarshalling.
type ErrorUnmarshaller func(*Response) error

// NewLROPoller creates a poller based on the provided initial response.
// NOTE: this is only meant for internal use in generated code.
func NewLROPoller(resp *Response, pl Pipeline, eu ErrorUnmarshaller) (*LROPoller, error) {
	// this is a back-stop in case the swagger is incorrect (i.e. missing one or more status codes for success).
	// ideally the codegen should return an error if the initial response failed and not even create a poller.
	if !lroStatusCodeValid(resp) {
		return nil, errors.New("the operation failed or was cancelled")
	}
	opLoc := resp.Header.Get(headerOperationLocation)
	loc := resp.Header.Get(headerLocation)
	// in the case of both headers, always prefer the operation-location header
	if opLoc != "" {
		return &LROPoller{
			lro:  &opPoller{req: resp.Request, pollURL: opLoc, locURL: loc},
			pl:   pl,
			eu:   eu,
			resp: resp,
		}, nil
	}
	if loc != "" {
		return &LROPoller{
			lro:  &locPoller{pollURL: loc, status: resp.StatusCode},
			pl:   pl,
			eu:   eu,
			resp: resp,
		}, nil
	}
	return &LROPoller{lro: &nopPoller{}}, nil
}

// LROPoller defines the methods that will be called internally in the generated code for long-running operations.
// NOTE: this is only meant for internal use in generated code.
type LROPoller struct {
	lro  lroPoller
	pl   Pipeline
	eu   ErrorUnmarshaller
	resp *Response
	err  error
}

// Done returns true if the LRO has reached a terminal state.
func (l *LROPoller) Done() bool {
	if l.err != nil {
		return true
	}
	return l.lro.Done()
}

// Poll sends a polling request to the polling endpoint and returns the response or an error.
func (l *LROPoller) Poll(ctx context.Context) (*http.Response, error) {
	if l.Done() {
		// the LRO has reached a terminal state, don't poll again
		if l.resp != nil {
			return l.resp.Response, nil
		}
		return nil, l.err
	}
	req, err := NewRequest(ctx, http.MethodGet, l.lro.PollURL())
	if err != nil {
		return nil, err
	}
	resp, err := l.pl.Do(req)
	if err != nil {
		// don't update the poller for failed requests
		return nil, err
	}
	if !lroStatusCodeValid(resp) {
		// the LRO failed.  unmarshall the error and update state
		l.err = l.eu(resp)
		l.resp = nil
		return nil, l.err
	}
	if err = l.lro.Update(resp); err != nil {
		return nil, err
	}
	return resp.Response, nil
}

// ResumeToken returns a token string that can be used to resume a poller that has not yet reached a terminal state.
func (l *LROPoller) ResumeToken() (string, error) {
	if l.Done() {
		return "", errors.New("cannot create a ResumeToken from a poller in a terminal state")
	}
	panic("NYI")
}

// FinalResponse will perform a final GET if request and return the final http response for the polling
// operation and unmarshall the content of the payload into the respType interface that is provided.
func (l *LROPoller) FinalResponse(ctx context.Context, respType interface{}) (*http.Response, error) {
	if !l.Done() {
		return nil, errors.New("cannot return a final response from a poller in a non-terminal state")
	}
	// if there's nothing to unmarshall into just return the final response
	if respType == nil {
		return l.resp.Response, nil
	}
	u, err := l.lro.FinalGetURL(l.resp)
	if err != nil {
		return nil, err
	}
	if u != "" {
		req, err := NewRequest(ctx, http.MethodGet, u)
		if err != nil {
			return nil, err
		}
		resp, err := l.pl.Do(req)
		if err != nil {
			return nil, err
		}
		if !lroStatusCodeValid(resp) {
			// TODO: unmarshall?
			return nil, errors.New("TODO")
		}
		l.resp = resp
	}
	body, err := ioutil.ReadAll(l.resp.Body)
	l.resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, respType); err != nil {
		return nil, err
	}
	return l.resp.Response, nil
}

// PollUntilDone will handle the entire span of the polling operation until a terminal state is reached,
// then return the final http response for the polling operation and unmarshal the content of the payload
// into the respType interface that is provided.
func (l *LROPoller) PollUntilDone(ctx context.Context, freq time.Duration, respType interface{}) (*http.Response, error) {
	logPollUntilDoneExit := func(v interface{}) {
		Log().Writef(LogLongRunningOperation, "END PollUntilDone() for %T: %v", l.lro, v)
	}
	Log().Writef(LogLongRunningOperation, "BEGIN PollUntilDone() for %T", l.lro)
	if l.resp != nil {
		// initial check for a retry-after header existing on the initial response
		if retryAfter := RetryAfter(l.resp.Response); retryAfter > 0 {
			Log().Writef(LogLongRunningOperation, "initial Retry-After delay for %s", retryAfter.String())
			if err := delay(ctx, retryAfter); err != nil {
				logPollUntilDoneExit(err)
				return nil, err
			}
		}
	}
	// begin polling the endpoint until a terminal state is reached
	for {
		resp, err := l.Poll(ctx)
		if err != nil {
			logPollUntilDoneExit(err)
			return nil, err
		}
		if l.Done() {
			logPollUntilDoneExit(l.lro.Status())
			break
		}
		d := freq
		if retryAfter := RetryAfter(resp); retryAfter > 0 {
			Log().Writef(LogLongRunningOperation, "Retry-After delay for %s", retryAfter.String())
			d = retryAfter
		} else {
			Log().Writef(LogLongRunningOperation, "delay for %s", d.String())
		}
		if err = delay(ctx, d); err != nil {
			logPollUntilDoneExit(err)
			return nil, err
		}
	}
	// TODO: distinguish between success/failure
	return l.FinalResponse(ctx, respType)
}

var _ Poller = (*LROPoller)(nil)

// abstracts the differences between concrete poller types
type lroPoller interface {
	Done() bool
	Update(resp *Response) error
	FinalGetURL(resp *Response) (string, error)
	PollURL() string
	Status() string
}

// polls on the operation-location header
type opPoller struct {
	req     *http.Request
	pollURL string
	locURL  string
	status  string
}

func (p *opPoller) PollURL() string {
	return p.pollURL
}

func (p *opPoller) Done() bool {
	return strings.EqualFold(p.status, "succeeded") ||
		strings.EqualFold(p.status, "failed") ||
		strings.EqualFold(p.status, "cancelled")
}

func (p *opPoller) Update(resp *Response) error {
	// read the response body
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}
	// put the body back so it's available to our callers
	resp.Body = ioutil.NopCloser(bytes.NewReader(body))
	status, err := extractJSONValue(resp.Body, "status")
	if err != nil {
		return err
	}
	p.status = status
	// if the endpoint returned an operation-location header, update cached value
	if opLoc := resp.Header.Get(headerOperationLocation); opLoc != "" {
		p.pollURL = opLoc
	}
	return nil
}

func (p *opPoller) FinalGetURL(resp *Response) (string, error) {
	if !p.Done() {
		return "", errors.New("cannot return a final response from a poller in a non-terminal state")
	}
	resLoc, err := extractJSONValue(resp.Body, "resourceLocation")
	if err != nil {
		return "", err
	}
	if resLoc != "" {
		return resLoc, nil
	}
	if p.req.Method == http.MethodPatch || p.req.Method == http.MethodPut {
		return p.req.URL.String(), nil
	}
	if p.req.Method == http.MethodPost && p.locURL != "" {
		return p.locURL, nil
	}
	return "", nil
}

func (p *opPoller) Status() string {
	return p.status
}

// polls on the location header
type locPoller struct {
	pollURL string
	status  int
}

func (p *locPoller) PollURL() string {
	return p.pollURL
}

func (p *locPoller) Done() bool {
	// a 202 means the operation is still in progress
	return p.status != http.StatusAccepted
}

func (p *locPoller) Update(resp *Response) error {
	// if the endpoint returned a location header, update cached value
	if loc := resp.Header.Get(headerLocation); loc != "" {
		p.pollURL = loc
	}
	p.status = resp.StatusCode
	return nil
}

func (*locPoller) FinalGetURL(*Response) (string, error) {
	return "", nil
}

func (p *locPoller) Status() string {
	return strconv.Itoa(p.status)
}

// used if the endpoint didn't return any polling headers (synchronous completion)
type nopPoller struct{}

func (*nopPoller) PollURL() string {
	return ""
}

func (*nopPoller) Done() bool {
	return true
}

func (*nopPoller) Update(*Response) error {
	return nil
}

func (*nopPoller) FinalGetURL(*Response) (string, error) {
	return "", nil
}

func (*nopPoller) Status() string {
	return "succeeded"
}

func extractJSONValue(r io.Reader, val string) (string, error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	// unmarshall the body to get the value
	var jsonBody map[string]interface{}
	if err = json.Unmarshal(body, &jsonBody); err != nil {
		return "", err
	}
	if len(jsonBody) == 0 {
		return "", errors.New("the response does not contain a body")
	}
	v, ok := jsonBody[val]
	if !ok {
		return "", fmt.Errorf("the response body does not contain field %s", v)
	}
	vv, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("the status value %v was not in string format", vv)
	}
	return vv, nil
}

func delay(ctx context.Context, delay time.Duration) error {
	select {
	case <-time.After(delay):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
