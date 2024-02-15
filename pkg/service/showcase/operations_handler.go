// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package showcase

import (
	"context"
	"fmt"
	"net/http"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/go-kit/log/level"
	"github.com/googleapis/gapic-showcase/util/genrest/resttools"
	"github.com/gorilla/mux"
	"github.com/qclaogui/gaip/pkg/protocol/rest"
)

// HandleListOperations translates REST requests/responses on the wire to internal proto messages for ListOperations
//
//	HTTP binding pattern: GET "/v1beta1/operations"
func (srv *Server) HandleListOperations() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numURLPathParams := len(urlPathParams)

		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/operations': %q", r.Method, r.URL))
		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("urlPathParams (expect 0, have %d): %q", numURLPathParams, urlPathParams))

		if numURLPathParams != 0 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 0, have %d: %#v", numURLPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := resttools.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &longrunningpb.ListOperationsRequest{}
		if err = resttools.CheckRequestFormat(nil, r, request.ProtoReflect()); err != nil {
			rest.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
			return
		}

		if err = resttools.PopulateSingularFields(request, urlPathParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
			return
		}

		if err = resttools.PopulateFields(request, queryParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading query params: %s", err)
			return
		}

		marshaler := resttools.ToJSON()
		marshaler.UseEnumNumbers = systemParameters.EnumEncodingAsInt
		requestJSON, _ := marshaler.Marshal(request)

		_ = level.Info(srv.logger).Log("msg", fmt.Sprintf("request: %s", requestJSON))

		ctx := context.WithValue(r.Context(), resttools.BindingURIKey, "/v1beta1/operations")
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
//	HTTP binding pattern: GET "/v1beta1/{name=operations/**}"
func (srv *Server) HandleGetOperation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numURLPathParams := len(urlPathParams)

		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{name=operations/**}': %q", r.Method, r.URL))
		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("urlPathParams (expect 1, have %d): %q", numURLPathParams, urlPathParams))

		if numURLPathParams != 1 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 0, have %d: %#v", numURLPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := resttools.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &longrunningpb.GetOperationRequest{}
		if err = resttools.CheckRequestFormat(nil, r, request.ProtoReflect()); err != nil {
			rest.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
			return
		}

		if err = resttools.PopulateSingularFields(request, urlPathParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
			return
		}

		excludedQueryParams := []string{"name"}
		if duplicates := resttools.KeysMatchPath(queryParams, excludedQueryParams); len(duplicates) > 0 {
			rest.Error(w, http.StatusBadRequest, "(QueryParamsInvalidFieldError) found keys that should not appear in query params: %v", duplicates)
			return
		}

		if err = resttools.PopulateFields(request, queryParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading query params: %s", err)
			return
		}

		marshaler := resttools.ToJSON()
		marshaler.UseEnumNumbers = systemParameters.EnumEncodingAsInt
		requestJSON, _ := marshaler.Marshal(request)

		_ = level.Info(srv.logger).Log("msg", fmt.Sprintf("request: %s", requestJSON))

		ctx := context.WithValue(r.Context(), resttools.BindingURIKey, "/v1beta1/{name=operations/**}")
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
//	HTTP binding pattern: DELETE "/v1beta1/{name=operations/**}"
func (srv *Server) HandleDeleteOperation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numURLPathParams := len(urlPathParams)

		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{name=operations/**}': %q", r.Method, r.URL))
		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("urlPathParams (expect 1, have %d): %q", numURLPathParams, urlPathParams))

		if numURLPathParams != 1 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 0, have %d: %#v", numURLPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := resttools.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &longrunningpb.DeleteOperationRequest{}
		if err = resttools.CheckRequestFormat(nil, r, request.ProtoReflect()); err != nil {
			rest.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
			return
		}

		if err = resttools.PopulateSingularFields(request, urlPathParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
			return
		}

		excludedQueryParams := []string{"name"}
		if duplicates := resttools.KeysMatchPath(queryParams, excludedQueryParams); len(duplicates) > 0 {
			rest.Error(w, http.StatusBadRequest, "(QueryParamsInvalidFieldError) found keys that should not appear in query params: %v", duplicates)
			return
		}

		if err = resttools.PopulateFields(request, queryParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading query params: %s", err)
			return
		}

		marshaler := resttools.ToJSON()
		marshaler.UseEnumNumbers = systemParameters.EnumEncodingAsInt
		requestJSON, _ := marshaler.Marshal(request)

		_ = level.Info(srv.logger).Log("msg", fmt.Sprintf("request: %s", requestJSON))

		ctx := context.WithValue(r.Context(), resttools.BindingURIKey, "/v1beta1/{name=operations/**}")
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
//	HTTP binding pattern: POST "/v1beta1/{name=operations/**}:cancel"
func (srv *Server) HandleCancelOperation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numURLPathParams := len(urlPathParams)

		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{name=operations/**}:cancel': %q", r.Method, r.URL))
		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("urlPathParams (expect 1, have %d): %q", numURLPathParams, urlPathParams))

		if numURLPathParams != 1 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 0, have %d: %#v", numURLPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := resttools.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &longrunningpb.CancelOperationRequest{}
		if err = resttools.CheckRequestFormat(nil, r, request.ProtoReflect()); err != nil {
			rest.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
			return
		}

		if err = resttools.PopulateSingularFields(request, urlPathParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
			return

		}

		excludedQueryParams := []string{"name"}
		if duplicates := resttools.KeysMatchPath(queryParams, excludedQueryParams); len(duplicates) > 0 {
			rest.Error(w, http.StatusBadRequest, "(QueryParamsInvalidFieldError) found keys that should not appear in query params: %v", duplicates)
			return
		}

		if err = resttools.PopulateFields(request, queryParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading query params: %s", err)
			return
		}

		marshaler := resttools.ToJSON()
		marshaler.UseEnumNumbers = systemParameters.EnumEncodingAsInt
		requestJSON, _ := marshaler.Marshal(request)

		_ = level.Info(srv.logger).Log("msg", fmt.Sprintf("request: %s", requestJSON))

		ctx := context.WithValue(r.Context(), resttools.BindingURIKey, "/v1beta1/{name=operations/**}:cancel")
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
