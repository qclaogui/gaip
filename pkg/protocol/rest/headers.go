// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package rest

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	headerNameContentType      = "Content-Type"
	headerValueContentTypeJSON = "application/json"

	headerNameAPIClient            = "X-Goog-Api-Client"
	headerValueTransportRESTPrefix = "rest/"
	headerValueClientGAPICPrefix   = "gapic/"
)

// PopulateRequestHeaders inspects request and adds the correct headers. This
// is useful for tests where we're not trying to send incorrect
// headers.
func PopulateRequestHeaders(request *http.Request) {
	header := http.Header{}
	header.Set(headerNameAPIClient, fmt.Sprintf("%s0.0.0 %s0.0.0", headerValueTransportRESTPrefix, headerValueClientGAPICPrefix))

	if request.Body != nil {
		header.Set(headerNameContentType, headerValueContentTypeJSON)
	}

	request.Header = header
}

// CheckContentType checks header to ensure the expected JSON content type is specified.
func CheckContentType(header http.Header) error {
	if content, ok := header[headerNameContentType]; !ok || len(content) != 1 || !strings.HasPrefix(strings.ToLower(strings.TrimSpace(content[0])), headerValueContentTypeJSON) {
		return fmt.Errorf("(HeaderContentTypeError) did not find expected HTTP header %q: %q", headerNameContentType, headerValueContentTypeJSON)
	}
	return nil
}

// CheckAPIClientHeader verifies that the "x-goog-api-client" header contains the expected tokens
// ("rest/..." and "gapic/...").
func CheckAPIClientHeader(header http.Header) error {
	content, ok := header[headerNameAPIClient]
	if !ok || len(content) != 1 {
		return fmt.Errorf("(HeaderAPIClientError) did not find expected HTTP header %q: %q\n                found: %q",
			headerNameAPIClient, headerValueTransportRESTPrefix, header)
	}

	var haveREST, haveGAPIC bool
	for _, token := range strings.Split(content[0], " ") {
		trimmed := strings.TrimSpace(token)
		if !haveREST && strings.HasPrefix(trimmed, headerValueTransportRESTPrefix) {
			haveREST = true
		} else if !haveGAPIC && strings.HasPrefix(trimmed, headerValueClientGAPICPrefix) {
			haveGAPIC = true
		} else {
			// nothing changed
			continue
		}
		if haveREST && haveGAPIC {
			return nil
		}
	}
	if !haveREST {
		return fmt.Errorf("(HeaderTransportRESTError) did not find expected HTTP header token %q: %q", headerNameAPIClient, headerValueTransportRESTPrefix)
	}
	if !haveGAPIC {
		return fmt.Errorf("(HeaderClientGAPICError) did not find expected HTTP header token %q: %q", headerNameAPIClient, headerValueClientGAPICPrefix)
	}
	return fmt.Errorf("internal inconsistency")
}
