// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package showcase

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-kit/log/level"
	"github.com/googleapis/gapic-showcase/util/genrest/resttools"
	"github.com/gorilla/mux"
	pb "github.com/qclaogui/gaip/genproto/showcase/apiv1beta1/showcasepb"
	"github.com/qclaogui/gaip/pkg/protocol/rest"
)

// HandleCreateUser translates REST requests/responses on the wire to internal proto messages for CreateUser
//
//	HTTP binding pattern: POST "/v1beta1/users"
func (srv *Server) HandleCreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numURLPathParams := len(urlPathParams)

		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/users': %q", r.Method, r.URL))
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

		request := &pb.CreateUserRequest{}
		// Intentional: Field values in the URL path override those set in the body.
		var jsonReader bytes.Buffer
		bodyReader := io.TeeReader(r.Body, &jsonReader)
		rBytes := make([]byte, r.ContentLength)
		if _, err = bodyReader.Read(rBytes); err != nil && !errors.Is(err, io.EOF) {
			rest.Error(w, http.StatusBadRequest, "error reading body content: %s", err)
			return
		}

		if err = resttools.FromJSON().Unmarshal(rBytes, request); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading body params '*': %s", err)
			return
		}

		if err = resttools.CheckRequestFormat(&jsonReader, r, request.ProtoReflect()); err != nil {
			rest.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
			return
		}

		if len(queryParams) > 0 {
			rest.Error(w, http.StatusBadRequest, "encountered unexpected query params: %v", queryParams)
			return
		}

		if err = resttools.PopulateSingularFields(request, urlPathParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
			return
		}

		marshaler := resttools.ToJSON()
		marshaler.UseEnumNumbers = systemParameters.EnumEncodingAsInt
		requestJSON, _ := marshaler.Marshal(request)

		_ = level.Info(srv.logger).Log("msg", fmt.Sprintf("request: %s", requestJSON))

		ctx := context.WithValue(r.Context(), resttools.BindingURIKey, "/v1beta1/users")
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
func (srv *Server) HandleGetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numURLPathParams := len(urlPathParams)

		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{name=users/*}': %q", r.Method, r.URL))
		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("urlPathParams (expect 0, have %d): %q", numURLPathParams, urlPathParams))
		if numURLPathParams != 1 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 0, have %d: %#v", numURLPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := resttools.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &pb.GetUserRequest{}
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

		ctx := context.WithValue(r.Context(), resttools.BindingURIKey, "/v1beta1/{name=users/*}")
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

// HandleUpdateUser translates REST requests/responses on the wire to internal proto messages for UpdateUser
//
//	Generated for HTTP binding pattern: PATCH "/v1beta1/{user.name=users/*}"
func (srv *Server) HandleUpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numURLPathParams := len(urlPathParams)

		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{user.name=users/*}': %q", r.Method, r.URL))
		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("urlPathParams (expect 0, have %d): %q", numURLPathParams, urlPathParams))
		if numURLPathParams != 1 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 0, have %d: %#v", numURLPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := resttools.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &pb.UpdateUserRequest{}
		// Intentional: Field values in the URL path override those set in the body.
		var bodyField pb.User
		var jsonReader bytes.Buffer
		bodyReader := io.TeeReader(r.Body, &jsonReader)
		rBytes := make([]byte, r.ContentLength)
		if _, err = bodyReader.Read(rBytes); err != nil && !errors.Is(err, io.EOF) {
			rest.Error(w, http.StatusBadRequest, "error reading body content: %s", err)
			return
		}

		if err = resttools.FromJSON().Unmarshal(rBytes, &bodyField); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading body params '*': %s", err)
			return
		}

		if err = resttools.CheckRequestFormat(&jsonReader, r, request.ProtoReflect()); err != nil {
			rest.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
			return
		}

		request.User = &bodyField

		if err = resttools.PopulateSingularFields(request, urlPathParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
			return
		}

		excludedQueryParams := []string{"user", "user.name"}
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

		ctx := context.WithValue(r.Context(), resttools.BindingURIKey, "/v1beta1/{user.name=users/*}")
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
func (srv *Server) HandleDeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numURLPathParams := len(urlPathParams)

		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{name=users/*}': %q", r.Method, r.URL))
		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("urlPathParams (expect 0, have %d): %q", numURLPathParams, urlPathParams))
		if numURLPathParams != 1 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 0, have %d: %#v", numURLPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := resttools.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &pb.DeleteUserRequest{}
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

		ctx := context.WithValue(r.Context(), resttools.BindingURIKey, "/v1beta1/{name=users/*}")
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

// HandleListUsers translates REST requests/responses on the wire to internal proto messages for ListUsers
//
//	Generated for HTTP binding pattern: GET "/v1beta1/users"
func (srv *Server) HandleListUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numURLPathParams := len(urlPathParams)

		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/users': %q", r.Method, r.URL))
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

		request := &pb.ListUsersRequest{}
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

		ctx := context.WithValue(r.Context(), resttools.BindingURIKey, "/v1beta1/users")
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
