// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package project

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
	"github.com/qclaogui/gaip/genproto/project/apiv1/projectpb"
	"github.com/qclaogui/gaip/pkg/protocol/rest"
)

// HandleEcho translates REST requests/responses on the wire to internal proto messages for Echo
//
//	HTTP binding pattern: POST "/v1/echo:echo"
func (srv *Server) HandleEcho() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numURLPathParams := len(urlPathParams)
		_ = level.Info(srv.logger).Log("msg", fmt.Sprintf("urlPathParams (expect 0, have %d): %q", numURLPathParams, urlPathParams))

		if numURLPathParams != 0 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 0, have %d: %#v", numURLPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := resttools.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &projectpb.EchoRequest{}
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

		ctx := context.WithValue(r.Context(), resttools.BindingURIKey, "/v1/echo:echo")
		response, err := srv.Echo(ctx, request)
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

// HandleEchoErrorDetails translates REST requests/responses on the wire to internal proto messages for EchoErrorDetails
//
//	Generated for HTTP binding pattern: POST "/v1/echo:error-details"
func (srv *Server) HandleEchoErrorDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numURLPathParams := len(urlPathParams)
		_ = level.Info(srv.logger).Log("msg", fmt.Sprintf("urlPathParams (expect 0, have %d): %q", numURLPathParams, urlPathParams))

		if numURLPathParams != 0 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 0, have %d: %#v", numURLPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := resttools.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &projectpb.EchoErrorDetailsRequest{}
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

		ctx := context.WithValue(r.Context(), resttools.BindingURIKey, "/v1/echo:error-details")
		response, err := srv.EchoErrorDetails(ctx, request)
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
