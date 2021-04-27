// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package testframework

import (
	"net/http"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
)

type RecordingSanitizer struct {
	recorder          *recorder.Recorder
	headersToSanitize map[string]*string
	urlSanitizer      StringSanitizer
	bodySanitizerReq  StringSanitizer
	bodySanitizerResp StringSanitizer
}

type StringSanitizer func(*string)

const SanitizedValue string = "sanitized"
const SanitizedBase64Value string = "Kg=="

var sanitizedValueSlice = []string{SanitizedValue}

func DefaultSanitizer(recorder *recorder.Recorder) *RecordingSanitizer {
	// The default sanitizer sanitizes the Authorization header
	s := &RecordingSanitizer{headersToSanitize: map[string]*string{"Authorization": nil}, recorder: recorder, urlSanitizer: DefaultStringSanitizer, bodySanitizerReq: DefaultStringSanitizer, bodySanitizerResp: DefaultStringSanitizer}
	recorder.AddSaveFilter(s.applySaveFilter)

	return s
}

// AddSanitizedHeaders adds the supplied header names to the list of headers to be sanitized on request and response recordings.
func (s *RecordingSanitizer) AddSanitizedHeaders(headers ...string) {
	for _, headerName := range headers {
		s.headersToSanitize[headerName] = nil
	}
}

// AddBodysanitizerReq configures the supplied StringSanitizer to sanitize recording request body
func (s *RecordingSanitizer) AddBodysanitizerReq(sanitizer StringSanitizer) {
	s.bodySanitizerReq = sanitizer
}

// AddBodysanitizerResp configures the supplied StringSanitizer to sanitize recording response bodies
func (s *RecordingSanitizer) AddBodysanitizerResp(sanitizer StringSanitizer) {
	s.bodySanitizerResp = sanitizer
}

// AddUriSanitizer configures the supplied StringSanitizer to sanitize recording request and response URLs
func (s *RecordingSanitizer) AddUrlSanitizer(sanitizer StringSanitizer) {
	s.urlSanitizer = sanitizer
}

func (s *RecordingSanitizer) sanitizeHeaders(header http.Header) {
	for headerName := range s.headersToSanitize {
		if _, ok := header[headerName]; ok {
			header[headerName] = sanitizedValueSlice
		}
	}
}

func (s *RecordingSanitizer) sanitizeBodiesReq(body *string) {
	s.bodySanitizerReq(body)
}

func (s *RecordingSanitizer) sanitizeBodiesResp(body *string) {
	s.bodySanitizerResp(body)
}

func (s *RecordingSanitizer) sanitizeURL(url *string) {
	s.urlSanitizer(url)
}

func (s *RecordingSanitizer) applySaveFilter(i *cassette.Interaction) error {
	s.sanitizeHeaders(i.Request.Headers)
	s.sanitizeHeaders(i.Response.Headers)
	s.sanitizeURL(&i.Request.URL)
	if len(i.Request.Body) > 0 {
		s.sanitizeBodiesReq(&i.Request.Body)
	}
	if len(i.Response.Body) > 0 {
		s.sanitizeBodiesResp(&i.Response.Body)
	}
	return nil
}

func DefaultStringSanitizer(s *string) {}
