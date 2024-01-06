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
			return
		}

		systemParameters, queryParams, err := resttools.GetSystemParameters(r)
		if err != nil {
			return
		}

		request := &projectpb.EchoRequest{}
		// Intentional: Field values in the URL path override those set in the body.
		var jsonReader bytes.Buffer
		bodyReader := io.TeeReader(r.Body, &jsonReader)
		rBytes := make([]byte, r.ContentLength)
		_, err = bodyReader.Read(rBytes)
		if err != nil && !errors.Is(err, io.EOF) {
			_ = level.Error(srv.logger).Log("msg", "error reading body content", "error", err)
			return
		}

		if err = resttools.FromJSON().Unmarshal(rBytes, request); err != nil {
			_ = level.Error(srv.logger).Log("msg", "error reading body params", "error", err)
			return
		}

		if err = resttools.CheckRequestFormat(&jsonReader, r, request.ProtoReflect()); err != nil {
			_ = level.Error(srv.logger).Log("msg", "REST request failed format check", "error", err)
			return
		}

		if len(queryParams) > 0 {
			_ = level.Error(srv.logger).Log("msg", "encountered unexpected query params", "params", queryParams)
			return
		}
		if err = resttools.PopulateSingularFields(request, urlPathParams); err != nil {
			_ = level.Error(srv.logger).Log("msg", "error reading URL path params", "error", err)
			return
		}

		marshaler := resttools.ToJSON()
		marshaler.UseEnumNumbers = systemParameters.EnumEncodingAsInt
		requestJSON, _ := marshaler.Marshal(request)
		_ = level.Info(srv.logger).Log("msg", fmt.Sprintf("request: %s", requestJSON))

		ctx := context.WithValue(r.Context(), resttools.BindingURIKey, "/v1/echo:echo")

		response, err := srv.Echo(ctx, request)
		if err != nil {
			//backend.ReportGRPCError(w, err)
			return
		}

		json, err := marshaler.Marshal(response)
		if err != nil {
			_ = level.Info(srv.logger).Log("msg", "error json-encoding response", "error", err)
			return
		}

		_, _ = w.Write(json)

	}
}
