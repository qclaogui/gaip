// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

// DO NOT EDIT. This is an auto-generated file containing the REST handlers
// for service #3: "Operations" (.google.longrunning.Operations).

package genrest

import (
	"context"
	"fmt"
	"net/http"

	longrunningpbpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/gorilla/mux"
	"github.com/qclaogui/gaip/pkg/protocol/rest"
)

// HandleListOperations translates REST requests/responses on the wire to internal proto messages for ListOperations
//
//	Generated for HTTP binding pattern: GET "/v1beta1/operations"
func HandleListOperations(srv longrunningpbpb.OperationsServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/operations': %q", r.Method, r.URL))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlPathParams (expect 0, have %d): %q", numUrlPathParams, urlPathParams))

		if numUrlPathParams != 0 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 0, have %d: %#v", numUrlPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &longrunningpbpb.ListOperationsRequest{}
		if err = rest.CheckRequestFormat(nil, r, request.ProtoReflect()); err != nil {
			rest.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
			return
		}
		if err = rest.PopulateSingularFields(request, urlPathParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
			return
		}

		// TODO: Decide whether query-param value or URL-path value takes precedence when a field appears in both
		if err = rest.PopulateFields(request, queryParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading query params: %s", err)
			return
		}

		marshaler := rest.ToJSON()
		marshaler.UseEnumNumbers = systemParameters.EnumEncodingAsInt
		requestJSON, _ := marshaler.Marshal(request)
		_ = level.Info(logger).Log("msg", fmt.Sprintf("request: %s", requestJSON))

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/operations")
		response, err := srv.ListOperations(ctx, request)
		if err != nil {
			rest.ReportGRPCError(w, err)
			return
		}

		json, err := marshaler.Marshal(response)
		if err != nil {
			rest.Error(w, http.StatusInternalServerError, "error json-encoding response: %s", err.Error())
			return
		}

		_, _ = w.Write(json)
	}
}

// HandleGetOperation translates REST requests/responses on the wire to internal proto messages for GetOperation
//
//	Generated for HTTP binding pattern: GET "/v1beta1/{name=operations/**}"
func HandleGetOperation(srv longrunningpbpb.OperationsServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{name=operations/**}': %q", r.Method, r.URL))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlPathParams (expect 1, have %d): %q", numUrlPathParams, urlPathParams))

		if numUrlPathParams != 1 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 1, have %d: %#v", numUrlPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &longrunningpbpb.GetOperationRequest{}
		if err = rest.CheckRequestFormat(nil, r, request.ProtoReflect()); err != nil {
			rest.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
			return
		}
		if err = rest.PopulateSingularFields(request, urlPathParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
			return
		}

		// TODO: Decide whether query-param value or URL-path value takes precedence when a field appears in both
		excludedQueryParams := []string{"name"}
		if duplicates := rest.KeysMatchPath(queryParams, excludedQueryParams); len(duplicates) > 0 {
			rest.Error(w, http.StatusBadRequest, "(QueryParamsInvalidFieldError) found keys that should not appear in query params: %v", duplicates)
			return
		}
		if err = rest.PopulateFields(request, queryParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading query params: %s", err)
			return
		}

		marshaler := rest.ToJSON()
		marshaler.UseEnumNumbers = systemParameters.EnumEncodingAsInt
		requestJSON, _ := marshaler.Marshal(request)
		_ = level.Info(logger).Log("msg", fmt.Sprintf("request: %s", requestJSON))

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/{name=operations/**}")
		response, err := srv.GetOperation(ctx, request)
		if err != nil {
			rest.ReportGRPCError(w, err)
			return
		}

		json, err := marshaler.Marshal(response)
		if err != nil {
			rest.Error(w, http.StatusInternalServerError, "error json-encoding response: %s", err.Error())
			return
		}

		_, _ = w.Write(json)
	}
}

// HandleDeleteOperation translates REST requests/responses on the wire to internal proto messages for DeleteOperation
//
//	Generated for HTTP binding pattern: DELETE "/v1beta1/{name=operations/**}"
func HandleDeleteOperation(srv longrunningpbpb.OperationsServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{name=operations/**}': %q", r.Method, r.URL))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlPathParams (expect 1, have %d): %q", numUrlPathParams, urlPathParams))

		if numUrlPathParams != 1 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 1, have %d: %#v", numUrlPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &longrunningpbpb.DeleteOperationRequest{}
		if err = rest.CheckRequestFormat(nil, r, request.ProtoReflect()); err != nil {
			rest.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
			return
		}
		if err = rest.PopulateSingularFields(request, urlPathParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
			return
		}

		// TODO: Decide whether query-param value or URL-path value takes precedence when a field appears in both
		excludedQueryParams := []string{"name"}
		if duplicates := rest.KeysMatchPath(queryParams, excludedQueryParams); len(duplicates) > 0 {
			rest.Error(w, http.StatusBadRequest, "(QueryParamsInvalidFieldError) found keys that should not appear in query params: %v", duplicates)
			return
		}
		if err = rest.PopulateFields(request, queryParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading query params: %s", err)
			return
		}

		marshaler := rest.ToJSON()
		marshaler.UseEnumNumbers = systemParameters.EnumEncodingAsInt
		requestJSON, _ := marshaler.Marshal(request)
		_ = level.Info(logger).Log("msg", fmt.Sprintf("request: %s", requestJSON))

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/{name=operations/**}")
		response, err := srv.DeleteOperation(ctx, request)
		if err != nil {
			rest.ReportGRPCError(w, err)
			return
		}

		json, err := marshaler.Marshal(response)
		if err != nil {
			rest.Error(w, http.StatusInternalServerError, "error json-encoding response: %s", err.Error())
			return
		}

		_, _ = w.Write(json)
	}
}

// HandleCancelOperation translates REST requests/responses on the wire to internal proto messages for CancelOperation
//
//	Generated for HTTP binding pattern: POST "/v1beta1/{name=operations/**}:cancel"
func HandleCancelOperation(srv longrunningpbpb.OperationsServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{name=operations/**}:cancel': %q", r.Method, r.URL))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlPathParams (expect 1, have %d): %q", numUrlPathParams, urlPathParams))

		if numUrlPathParams != 1 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 1, have %d: %#v", numUrlPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &longrunningpbpb.CancelOperationRequest{}
		if err = rest.CheckRequestFormat(nil, r, request.ProtoReflect()); err != nil {
			rest.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
			return
		}
		if err = rest.PopulateSingularFields(request, urlPathParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
			return
		}

		// TODO: Decide whether query-param value or URL-path value takes precedence when a field appears in both
		excludedQueryParams := []string{"name"}
		if duplicates := rest.KeysMatchPath(queryParams, excludedQueryParams); len(duplicates) > 0 {
			rest.Error(w, http.StatusBadRequest, "(QueryParamsInvalidFieldError) found keys that should not appear in query params: %v", duplicates)
			return
		}
		if err = rest.PopulateFields(request, queryParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading query params: %s", err)
			return
		}

		marshaler := rest.ToJSON()
		marshaler.UseEnumNumbers = systemParameters.EnumEncodingAsInt
		requestJSON, _ := marshaler.Marshal(request)
		_ = level.Info(logger).Log("msg", fmt.Sprintf("request: %s", requestJSON))

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/{name=operations/**}:cancel")
		response, err := srv.CancelOperation(ctx, request)
		if err != nil {
			rest.ReportGRPCError(w, err)
			return
		}

		json, err := marshaler.Marshal(response)
		if err != nil {
			rest.Error(w, http.StatusInternalServerError, "error json-encoding response: %s", err.Error())
			return
		}

		_, _ = w.Write(json)
	}
}

// RegisterHandlersOperations register REST requests/responses on the wire to internal proto messages for
func RegisterHandlersOperations(router *mux.Router, srv longrunningpbpb.OperationsServer, logger log.Logger) {
	router.Handle("/v1beta1/operations", HandleListOperations(srv, logger)).Methods("GET")
	router.Handle("/v1beta1/{name:operations/[^:]+}", HandleGetOperation(srv, logger)).Methods("GET")
	router.Handle("/v1beta1/{name:operations/[^:]+}", HandleDeleteOperation(srv, logger)).Methods("DELETE")
	router.Handle("/v1beta1/{name:operations/[^:]+}:cancel", HandleCancelOperation(srv, logger)).Methods("POST")
}
