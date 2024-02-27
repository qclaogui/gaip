//Copyright © Weifeng Wang <qclaogui@gmail.com>
//
//Licensed under the Apache License 2.0.

// DO NOT EDIT. This is an auto-generated file containing the REST handlers
// for service #1: "IdentityService" (.qclaogui.showcase.v1beta1.IdentityService).

package genrest

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/gorilla/mux"
	showcasepbpb "github.com/qclaogui/gaip/genproto/showcase/apiv1beta1/showcasepb"
	"github.com/qclaogui/gaip/pkg/protocol/rest"
)

// HandleCreateUser translates REST requests/responses on the wire to internal proto messages for CreateUser
//
//	Generated for HTTP binding pattern: POST "/v1beta1/users"
func HandleCreateUser(srv showcasepbpb.IdentityServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/users': %q", r.Method, r.URL))
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

		request := &showcasepbpb.CreateUserRequest{}
		// Intentional: Field values in the URL path override those set in the body.
		var jsonReader bytes.Buffer
		bodyReader := io.TeeReader(r.Body, &jsonReader)
		rBytes := make([]byte, r.ContentLength)
		if _, err = bodyReader.Read(rBytes); err != nil && !errors.Is(err, io.EOF) {
			rest.Error(w, http.StatusBadRequest, "error reading body content: %s", err)
			return
		}

		if err = rest.FromJSON().Unmarshal(rBytes, request); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading body params '*': %s", err)
			return
		}

		if err = rest.CheckRequestFormat(&jsonReader, r, request.ProtoReflect()); err != nil {
			rest.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
			return
		}

		if len(queryParams) > 0 {
			rest.Error(w, http.StatusBadRequest, "encountered unexpected query params: %v", queryParams)
			return
		}
		if err = rest.PopulateSingularFields(request, urlPathParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
			return
		}

		marshaler := rest.ToJSON()
		marshaler.UseEnumNumbers = systemParameters.EnumEncodingAsInt
		requestJSON, _ := marshaler.Marshal(request)
		_ = level.Info(logger).Log("msg", fmt.Sprintf("request: %s", requestJSON))

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/users")
		response, err := srv.CreateUser(ctx, request)
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

// HandleGetUser translates REST requests/responses on the wire to internal proto messages for GetUser
//
//	Generated for HTTP binding pattern: GET "/v1beta1/{name=users/*}"
func HandleGetUser(srv showcasepbpb.IdentityServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{name=users/*}': %q", r.Method, r.URL))
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

		request := &showcasepbpb.GetUserRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/{name=users/*}")
		response, err := srv.GetUser(ctx, request)
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

// HandleListUsers translates REST requests/responses on the wire to internal proto messages for ListUsers
//
//	Generated for HTTP binding pattern: GET "/v1beta1/users"
func HandleListUsers(srv showcasepbpb.IdentityServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/users': %q", r.Method, r.URL))
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

		request := &showcasepbpb.ListUsersRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/users")
		response, err := srv.ListUsers(ctx, request)
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

// HandleUpdateUser translates REST requests/responses on the wire to internal proto messages for UpdateUser
//
//	Generated for HTTP binding pattern: PATCH "/v1beta1/{user.name=users/*}"
func HandleUpdateUser(srv showcasepbpb.IdentityServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{user.name=users/*}': %q", r.Method, r.URL))
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

		request := &showcasepbpb.UpdateUserRequest{}
		// Intentional: Field values in the URL path override those set in the body.
		var bodyField showcasepbpb.User
		var jsonReader bytes.Buffer
		bodyReader := io.TeeReader(r.Body, &jsonReader)
		rBytes := make([]byte, r.ContentLength)
		if _, err = bodyReader.Read(rBytes); err != nil && err != io.EOF {
			rest.Error(w, http.StatusBadRequest, "error reading body content: %s", err)
			return
		}

		if err = rest.FromJSON().Unmarshal(rBytes, &bodyField); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading body into request field 'user': %s", err)
			return
		}

		if err = rest.CheckRequestFormat(&jsonReader, r, request.ProtoReflect()); err != nil {
			rest.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
			return
		}
		request.User = &bodyField

		if err = rest.PopulateSingularFields(request, urlPathParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
			return
		}

		// TODO: Decide whether query-param value or URL-path value takes precedence when a field appears in both
		excludedQueryParams := []string{"user", "user.name"}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/{user.name=users/*}")
		response, err := srv.UpdateUser(ctx, request)
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

// HandleDeleteUser translates REST requests/responses on the wire to internal proto messages for DeleteUser
//
//	Generated for HTTP binding pattern: DELETE "/v1beta1/{name=users/*}"
func HandleDeleteUser(srv showcasepbpb.IdentityServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{name=users/*}': %q", r.Method, r.URL))
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

		request := &showcasepbpb.DeleteUserRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/{name=users/*}")
		response, err := srv.DeleteUser(ctx, request)
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

// RegisterHandlersIdentityService register REST requests/responses on the wire to internal proto messages for
func RegisterHandlersIdentityService(router *mux.Router, srv showcasepbpb.IdentityServiceServer, logger log.Logger) {
	router.Handle("/v1beta1/users", HandleCreateUser(srv, logger)).Methods("POST")
	router.Handle("/v1beta1/{name:users/[^:]+}", HandleGetUser(srv, logger)).Methods("GET")
	router.Handle("/v1beta1/users", HandleListUsers(srv, logger)).Methods("GET")
	router.Handle("/v1beta1/{user.name:users/[^:]+}", HandleUpdateUser(srv, logger)).Methods("PATCH")
	router.Handle("/v1beta1/{user.name:users/[^:]+}", HandleUpdateUser(srv, logger)).HeadersRegexp("X-HTTP-Method-Override", "^PATCH$").Methods("POST")
	router.Handle("/v1beta1/{name:users/[^:]+}", HandleDeleteUser(srv, logger)).Methods("DELETE")
}
