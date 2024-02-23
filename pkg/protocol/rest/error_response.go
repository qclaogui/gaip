// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/api/googleapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// gRPCToHTTP status code mapping derived from internal source:
// go/http-canonical-mapping.
var gRPCToHTTP = map[codes.Code]int{
	codes.OK:                 http.StatusOK,
	codes.Canceled:           499, // There isn't a Go constant ClientClosedConnection
	codes.Unknown:            http.StatusInternalServerError,
	codes.InvalidArgument:    http.StatusBadRequest,
	codes.DeadlineExceeded:   http.StatusGatewayTimeout,
	codes.NotFound:           http.StatusNotFound,
	codes.AlreadyExists:      http.StatusConflict,
	codes.PermissionDenied:   http.StatusForbidden,
	codes.Unauthenticated:    http.StatusUnauthorized,
	codes.ResourceExhausted:  http.StatusTooManyRequests,
	codes.FailedPrecondition: http.StatusBadRequest,
	codes.Aborted:            http.StatusConflict,
	codes.OutOfRange:         http.StatusBadRequest,
	codes.Unimplemented:      http.StatusNotImplemented,
	codes.Internal:           http.StatusInternalServerError,
	codes.Unavailable:        http.StatusServiceUnavailable,
	codes.DataLoss:           http.StatusInternalServerError,
}

// Google API Errors, as defined by https://cloud.google.com/apis/design/errors
// will consist of a googleapi.Error nested as the key `error` in a JSON object.
// So we must create such a structure to wrap our googleapi.Error in.
type googleAPIError struct {
	Error *googleapi.Error `json:"error"`
}

// GRPCToHTTP maps the given gRPC Code to the canonical HTTP Status code as
// defined by the internal source go/http-canonical-mapping.
func GRPCToHTTP(c codes.Code) int {
	httpStatus, ok := gRPCToHTTP[c]
	if !ok {
		httpStatus = http.StatusInternalServerError
	}

	return httpStatus
}

// ErrorResponse is a helper that formats the given response information,
// including the HTTP Status code, a message, and any error detail types, into
// a googleAPIError and writes the response as JSON.
func ErrorResponse(w http.ResponseWriter, status int, message string, details ...interface{}) {
	apiError := &googleAPIError{
		Error: &googleapi.Error{
			Code:    status,
			Message: message,
			Details: details,
		},
	}
	w.WriteHeader(status)
	data, _ := json.Marshal(apiError)
	_, _ = w.Write(data)
}

func Error(w http.ResponseWriter, httpStatus int, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	//Print(message)
	ErrorResponse(w, httpStatus, message)
}

func ReportGRPCError(w http.ResponseWriter, err error) {
	st, ok := status.FromError(err)
	if !ok {
		Error(w, http.StatusInternalServerError, "server error: %s", err.Error())
		return
	}

	code := GRPCToHTTP(st.Code())
	ErrorResponse(w, code, st.Message(), st.Details()...)
}
