// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package rest

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/api/googleapi"
	"google.golang.org/grpc/codes"
)

func TestErrorResponse(t *testing.T) {
	for _, tc := range []struct {
		name    string
		message string
		status  int
		details []interface{}
	}{
		{
			name:    "internal_server",
			message: "Had an issue",
			status:  http.StatusInternalServerError,
			details: []interface{}{"foo"},
		},
		{
			name:    "bad_request",
			message: "The request was bad",
			status:  http.StatusBadRequest,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := httptest.NewRecorder()
			ErrorResponse(got, tc.status, tc.message, tc.details...)
			if got.Code != tc.status {
				t.Errorf("%s: Expected %d, but got %d", tc.name, tc.status, got.Code)
			}

			err := googleapi.CheckResponse(got.Result())
			var gerr *googleapi.Error
			if !errors.As(err, &gerr) {
				t.Fatalf("%s: Expected response to be a googleapi.Error, but got %v", tc.name, err)
			}

			if diff := cmp.Diff(gerr.Message, tc.message); diff != "" {
				t.Errorf("%s: got(-),want(+):%s\n", tc.name, diff)
			}

			if diff := cmp.Diff(gerr.Details, tc.details); diff != "" {
				t.Errorf("%s: got(-),want(+):%s\n", tc.name, diff)
			}
		})
	}
}

func TestGRPCToHTTP(t *testing.T) {
	testCase := []struct {
		code codes.Code
		want int
	}{
		{
			codes.Aborted,
			http.StatusConflict,
		},
		{
			100,
			http.StatusInternalServerError,
		},
	}

	for _, tc := range testCase {
		if got := GRPCToHTTP(tc.code); got != tc.want {
			t.Errorf("got %d, but expected %d", got, tc.want)
		}
	}
}
