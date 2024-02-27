// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

// DO NOT EDIT. This is an auto-generated file containing the REST handlers
// for service #2: "MessagingService" (.qclaogui.showcase.v1beta1.MessagingService).

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

// HandleCreateRoom translates REST requests/responses on the wire to internal proto messages for CreateRoom
//
//	Generated for HTTP binding pattern: POST "/v1beta1/rooms"
func HandleCreateRoom(srv showcasepbpb.MessagingServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/rooms': %q", r.Method, r.URL))
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

		request := &showcasepbpb.CreateRoomRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/rooms")
		response, err := srv.CreateRoom(ctx, request)
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

// HandleGetRoom translates REST requests/responses on the wire to internal proto messages for GetRoom
//
//	Generated for HTTP binding pattern: GET "/v1beta1/{name=rooms/*}"
func HandleGetRoom(srv showcasepbpb.MessagingServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{name=rooms/*}': %q", r.Method, r.URL))
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

		request := &showcasepbpb.GetRoomRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/{name=rooms/*}")
		response, err := srv.GetRoom(ctx, request)
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

// HandleUpdateRoom translates REST requests/responses on the wire to internal proto messages for UpdateRoom
//
//	Generated for HTTP binding pattern: PATCH "/v1beta1/{room.name=rooms/*}"
func HandleUpdateRoom(srv showcasepbpb.MessagingServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{room.name=rooms/*}': %q", r.Method, r.URL))
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

		request := &showcasepbpb.UpdateRoomRequest{}
		// Intentional: Field values in the URL path override those set in the body.
		var bodyField showcasepbpb.Room
		var jsonReader bytes.Buffer
		bodyReader := io.TeeReader(r.Body, &jsonReader)
		rBytes := make([]byte, r.ContentLength)
		if _, err = bodyReader.Read(rBytes); err != nil && err != io.EOF {
			rest.Error(w, http.StatusBadRequest, "error reading body content: %s", err)
			return
		}

		if err = rest.FromJSON().Unmarshal(rBytes, &bodyField); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading body into request field 'room': %s", err)
			return
		}

		if err = rest.CheckRequestFormat(&jsonReader, r, request.ProtoReflect()); err != nil {
			rest.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
			return
		}
		request.Room = &bodyField

		if err = rest.PopulateSingularFields(request, urlPathParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
			return
		}

		// TODO: Decide whether query-param value or URL-path value takes precedence when a field appears in both
		excludedQueryParams := []string{"room", "room.name"}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/{room.name=rooms/*}")
		response, err := srv.UpdateRoom(ctx, request)
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

// HandleDeleteRoom translates REST requests/responses on the wire to internal proto messages for DeleteRoom
//
//	Generated for HTTP binding pattern: DELETE "/v1beta1/{name=rooms/*}"
func HandleDeleteRoom(srv showcasepbpb.MessagingServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{name=rooms/*}': %q", r.Method, r.URL))
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

		request := &showcasepbpb.DeleteRoomRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/{name=rooms/*}")
		response, err := srv.DeleteRoom(ctx, request)
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

// HandleListRooms translates REST requests/responses on the wire to internal proto messages for ListRooms
//
//	Generated for HTTP binding pattern: GET "/v1beta1/rooms"
func HandleListRooms(srv showcasepbpb.MessagingServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/rooms': %q", r.Method, r.URL))
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

		request := &showcasepbpb.ListRoomsRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/rooms")
		response, err := srv.ListRooms(ctx, request)
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

// HandleCreateBlurb translates REST requests/responses on the wire to internal proto messages for CreateBlurb
//
//	Generated for HTTP binding pattern: POST "/v1beta1/{parent=rooms/*}/blurbs"
func HandleCreateBlurb(srv showcasepbpb.MessagingServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{parent=rooms/*}/blurbs': %q", r.Method, r.URL))
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

		request := &showcasepbpb.CreateBlurbRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/{parent=rooms/*}/blurbs")
		response, err := srv.CreateBlurb(ctx, request)
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

// HandleCreateBlurb1 translates REST requests/responses on the wire to internal proto messages for CreateBlurb
//
//	Generated for HTTP binding pattern: POST "/v1beta1/{parent=users/*/profile}/blurbs"
func HandleCreateBlurb1(srv showcasepbpb.MessagingServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{parent=users/*/profile}/blurbs': %q", r.Method, r.URL))
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

		request := &showcasepbpb.CreateBlurbRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/{parent=users/*/profile}/blurbs")
		response, err := srv.CreateBlurb(ctx, request)
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

// HandleGetBlurb translates REST requests/responses on the wire to internal proto messages for GetBlurb
//
//	Generated for HTTP binding pattern: GET "/v1beta1/{name=rooms/*/blurbs/*}"
func HandleGetBlurb(srv showcasepbpb.MessagingServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{name=rooms/*/blurbs/*}': %q", r.Method, r.URL))
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

		request := &showcasepbpb.GetBlurbRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/{name=rooms/*/blurbs/*}")
		response, err := srv.GetBlurb(ctx, request)
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

// HandleGetBlurb1 translates REST requests/responses on the wire to internal proto messages for GetBlurb
//
//	Generated for HTTP binding pattern: GET "/v1beta1/{name=users/*/profile/blurbs/*}"
func HandleGetBlurb1(srv showcasepbpb.MessagingServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{name=users/*/profile/blurbs/*}': %q", r.Method, r.URL))
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

		request := &showcasepbpb.GetBlurbRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/{name=users/*/profile/blurbs/*}")
		response, err := srv.GetBlurb(ctx, request)
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

// HandleUpdateBlurb translates REST requests/responses on the wire to internal proto messages for UpdateBlurb
//
//	Generated for HTTP binding pattern: PATCH "/v1beta1/{blurb.name=rooms/*/blurbs/*}"
func HandleUpdateBlurb(srv showcasepbpb.MessagingServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{blurb.name=rooms/*/blurbs/*}': %q", r.Method, r.URL))
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

		request := &showcasepbpb.UpdateBlurbRequest{}
		// Intentional: Field values in the URL path override those set in the body.
		var bodyField showcasepbpb.Blurb
		var jsonReader bytes.Buffer
		bodyReader := io.TeeReader(r.Body, &jsonReader)
		rBytes := make([]byte, r.ContentLength)
		if _, err = bodyReader.Read(rBytes); err != nil && err != io.EOF {
			rest.Error(w, http.StatusBadRequest, "error reading body content: %s", err)
			return
		}

		if err = rest.FromJSON().Unmarshal(rBytes, &bodyField); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading body into request field 'blurb': %s", err)
			return
		}

		if err = rest.CheckRequestFormat(&jsonReader, r, request.ProtoReflect()); err != nil {
			rest.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
			return
		}
		request.Blurb = &bodyField

		if err = rest.PopulateSingularFields(request, urlPathParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
			return
		}

		// TODO: Decide whether query-param value or URL-path value takes precedence when a field appears in both
		excludedQueryParams := []string{"blurb", "blurb.name"}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/{blurb.name=rooms/*/blurbs/*}")
		response, err := srv.UpdateBlurb(ctx, request)
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

// HandleUpdateBlurb1 translates REST requests/responses on the wire to internal proto messages for UpdateBlurb
//
//	Generated for HTTP binding pattern: PATCH "/v1beta1/{blurb.name=users/*/profile/blurbs/*}"
func HandleUpdateBlurb1(srv showcasepbpb.MessagingServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{blurb.name=users/*/profile/blurbs/*}': %q", r.Method, r.URL))
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

		request := &showcasepbpb.UpdateBlurbRequest{}
		// Intentional: Field values in the URL path override those set in the body.
		var bodyField showcasepbpb.Blurb
		var jsonReader bytes.Buffer
		bodyReader := io.TeeReader(r.Body, &jsonReader)
		rBytes := make([]byte, r.ContentLength)
		if _, err = bodyReader.Read(rBytes); err != nil && err != io.EOF {
			rest.Error(w, http.StatusBadRequest, "error reading body content: %s", err)
			return
		}

		if err = rest.FromJSON().Unmarshal(rBytes, &bodyField); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading body into request field 'blurb': %s", err)
			return
		}

		if err = rest.CheckRequestFormat(&jsonReader, r, request.ProtoReflect()); err != nil {
			rest.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
			return
		}
		request.Blurb = &bodyField

		if err = rest.PopulateSingularFields(request, urlPathParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
			return
		}

		// TODO: Decide whether query-param value or URL-path value takes precedence when a field appears in both
		excludedQueryParams := []string{"blurb", "blurb.name"}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/{blurb.name=users/*/profile/blurbs/*}")
		response, err := srv.UpdateBlurb(ctx, request)
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

// HandleDeleteBlurb translates REST requests/responses on the wire to internal proto messages for DeleteBlurb
//
//	Generated for HTTP binding pattern: DELETE "/v1beta1/{name=rooms/*/blurbs/*}"
func HandleDeleteBlurb(srv showcasepbpb.MessagingServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{name=rooms/*/blurbs/*}': %q", r.Method, r.URL))
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

		request := &showcasepbpb.DeleteBlurbRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/{name=rooms/*/blurbs/*}")
		response, err := srv.DeleteBlurb(ctx, request)
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

// HandleDeleteBlurb1 translates REST requests/responses on the wire to internal proto messages for DeleteBlurb
//
//	Generated for HTTP binding pattern: DELETE "/v1beta1/{name=users/*/profile/blurbs/*}"
func HandleDeleteBlurb1(srv showcasepbpb.MessagingServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{name=users/*/profile/blurbs/*}': %q", r.Method, r.URL))
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

		request := &showcasepbpb.DeleteBlurbRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/{name=users/*/profile/blurbs/*}")
		response, err := srv.DeleteBlurb(ctx, request)
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

// HandleListBlurbs translates REST requests/responses on the wire to internal proto messages for ListBlurbs
//
//	Generated for HTTP binding pattern: GET "/v1beta1/{parent=rooms/*}/blurbs"
func HandleListBlurbs(srv showcasepbpb.MessagingServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{parent=rooms/*}/blurbs': %q", r.Method, r.URL))
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

		request := &showcasepbpb.ListBlurbsRequest{}
		if err = rest.CheckRequestFormat(nil, r, request.ProtoReflect()); err != nil {
			rest.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
			return
		}
		if err = rest.PopulateSingularFields(request, urlPathParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
			return
		}

		// TODO: Decide whether query-param value or URL-path value takes precedence when a field appears in both
		excludedQueryParams := []string{"parent"}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/{parent=rooms/*}/blurbs")
		response, err := srv.ListBlurbs(ctx, request)
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

// HandleListBlurbs1 translates REST requests/responses on the wire to internal proto messages for ListBlurbs
//
//	Generated for HTTP binding pattern: GET "/v1beta1/{parent=users/*/profile}/blurbs"
func HandleListBlurbs1(srv showcasepbpb.MessagingServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{parent=users/*/profile}/blurbs': %q", r.Method, r.URL))
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

		request := &showcasepbpb.ListBlurbsRequest{}
		if err = rest.CheckRequestFormat(nil, r, request.ProtoReflect()); err != nil {
			rest.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
			return
		}
		if err = rest.PopulateSingularFields(request, urlPathParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
			return
		}

		// TODO: Decide whether query-param value or URL-path value takes precedence when a field appears in both
		excludedQueryParams := []string{"parent"}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/{parent=users/*/profile}/blurbs")
		response, err := srv.ListBlurbs(ctx, request)
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

// HandleSearchBlurbs translates REST requests/responses on the wire to internal proto messages for SearchBlurbs
//
//	Generated for HTTP binding pattern: POST "/v1beta1/{parent=rooms/*}/blurbs:search"
func HandleSearchBlurbs(srv showcasepbpb.MessagingServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{parent=rooms/*}/blurbs:search': %q", r.Method, r.URL))
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

		request := &showcasepbpb.SearchBlurbsRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/{parent=rooms/*}/blurbs:search")
		response, err := srv.SearchBlurbs(ctx, request)
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

// HandleSearchBlurbs1 translates REST requests/responses on the wire to internal proto messages for SearchBlurbs
//
//	Generated for HTTP binding pattern: POST "/v1beta1/{parent=users/*/profile}/blurbs:search"
func HandleSearchBlurbs1(srv showcasepbpb.MessagingServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{parent=users/*/profile}/blurbs:search': %q", r.Method, r.URL))
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

		request := &showcasepbpb.SearchBlurbsRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/{parent=users/*/profile}/blurbs:search")
		response, err := srv.SearchBlurbs(ctx, request)
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

// HandleStreamBlurbs translates REST requests/responses on the wire to internal proto messages for StreamBlurbs
//
//	Generated for HTTP binding pattern: POST "/v1beta1/{name=rooms/*}/blurbs:stream"
func HandleStreamBlurbs(srv showcasepbpb.MessagingServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{name=rooms/*}/blurbs:stream': %q", r.Method, r.URL))
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

		request := &showcasepbpb.StreamBlurbsRequest{}
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

		serverStreamer, err := rest.NewServerStreamer(w, rest.ServerStreamingChunkSize)
		if err != nil {
			rest.Error(w, http.StatusInternalServerError, "server error: could not construct server streamer: %s", err.Error())
			return
		}
		defer func() { _ = serverStreamer.End() }()
		streamer := &MessagingServiceStreamBlurbsServer{serverStreamer}
		if err = srv.StreamBlurbs(request, streamer); err != nil {
			rest.ReportGRPCError(w, err)
		}
	}
}

// HandleStreamBlurbs1 translates REST requests/responses on the wire to internal proto messages for StreamBlurbs
//
//	Generated for HTTP binding pattern: POST "/v1beta1/{name=users/*/profile}/blurbs:stream"
func HandleStreamBlurbs1(srv showcasepbpb.MessagingServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{name=users/*/profile}/blurbs:stream': %q", r.Method, r.URL))
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

		request := &showcasepbpb.StreamBlurbsRequest{}
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

		serverStreamer, err := rest.NewServerStreamer(w, rest.ServerStreamingChunkSize)
		if err != nil {
			rest.Error(w, http.StatusInternalServerError, "server error: could not construct server streamer: %s", err.Error())
			return
		}
		defer func() { _ = serverStreamer.End() }()
		streamer := &MessagingServiceStreamBlurbsServer{serverStreamer}
		if err = srv.StreamBlurbs(request, streamer); err != nil {
			rest.ReportGRPCError(w, err)
		}
	}
}

// HandleSendBlurbs translates REST requests/responses on the wire to internal proto messages for SendBlurbs
//
//	Generated for HTTP binding pattern: POST "/v1beta1/{parent=rooms/*}/blurbs:send"
func HandleSendBlurbs(srv showcasepbpb.MessagingServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = srv
		_ = level.Warn(logger).Log("msg", fmt.Sprintf("client-streaming methods not implemented yet (request matched '/v1beta1/{parent=rooms/*}/blurbs:send': %q)", r.URL))
		rest.Error(w, http.StatusNotImplemented, "client-streaming methods not implemented yet (request matched '/v1beta1/{parent=rooms/*}/blurbs:send': %q)", r.URL)
	}
}

// HandleSendBlurbs1 translates REST requests/responses on the wire to internal proto messages for SendBlurbs
//
//	Generated for HTTP binding pattern: POST "/v1beta1/{parent=users/*/profile}/blurbs:send"
func HandleSendBlurbs1(srv showcasepbpb.MessagingServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = srv
		_ = level.Warn(logger).Log("msg", fmt.Sprintf("client-streaming methods not implemented yet (request matched '/v1beta1/{parent=users/*/profile}/blurbs:send': %q)", r.URL))
		rest.Error(w, http.StatusNotImplemented, "client-streaming methods not implemented yet (request matched '/v1beta1/{parent=users/*/profile}/blurbs:send': %q)", r.URL)
	}
}

// MessagingServiceStreamBlurbsServer implements showcasepbpb.MessagingServiceStreamBlurbsServer to provide server-side streaming over REST, returning all the
// individual responses as part of a long JSON list.
type MessagingServiceStreamBlurbsServer struct {
	*rest.ServerStreamer
}

// Send accumulates a response to be fetched later as part of response list returned over REST.
func (streamer *MessagingServiceStreamBlurbsServer) Send(response *showcasepbpb.StreamBlurbsResponse) error {
	return streamer.ServerStreamer.Send(response)
}

// RegisterHandlersMessagingService register REST requests/responses on the wire to internal proto messages for
func RegisterHandlersMessagingService(router *mux.Router, srv showcasepbpb.MessagingServiceServer, logger log.Logger) {
	router.Handle("/v1beta1/rooms", HandleCreateRoom(srv, logger)).Methods("POST")
	router.Handle("/v1beta1/{name:rooms/[^:]+}", HandleGetRoom(srv, logger)).Methods("GET")
	router.Handle("/v1beta1/{room.name:rooms/[^:]+}", HandleUpdateRoom(srv, logger)).Methods("PATCH")
	router.Handle("/v1beta1/{room.name:rooms/[^:]+}", HandleUpdateRoom(srv, logger)).HeadersRegexp("X-HTTP-Method-Override", "^PATCH$").Methods("POST")
	router.Handle("/v1beta1/{name:rooms/[^:]+}", HandleDeleteRoom(srv, logger)).Methods("DELETE")
	router.Handle("/v1beta1/rooms", HandleListRooms(srv, logger)).Methods("GET")
	router.Handle("/v1beta1/{parent:rooms/[^:]+}/blurbs", HandleCreateBlurb(srv, logger)).Methods("POST")
	router.Handle("/v1beta1/{parent:users/[^:]+/profile}/blurbs", HandleCreateBlurb1(srv, logger)).Methods("POST")
	router.Handle("/v1beta1/{name:rooms/[^:]+/blurbs/[^:]+}", HandleGetBlurb(srv, logger)).Methods("GET")
	router.Handle("/v1beta1/{name:users/[^:]+/profile/blurbs/[^:]+}", HandleGetBlurb1(srv, logger)).Methods("GET")
	router.Handle("/v1beta1/{blurb.name:rooms/[^:]+/blurbs/[^:]+}", HandleUpdateBlurb(srv, logger)).Methods("PATCH")
	router.Handle("/v1beta1/{blurb.name:rooms/[^:]+/blurbs/[^:]+}", HandleUpdateBlurb(srv, logger)).HeadersRegexp("X-HTTP-Method-Override", "^PATCH$").Methods("POST")
	router.Handle("/v1beta1/{blurb.name:users/[^:]+/profile/blurbs/[^:]+}", HandleUpdateBlurb1(srv, logger)).Methods("PATCH")
	router.Handle("/v1beta1/{blurb.name:users/[^:]+/profile/blurbs/[^:]+}", HandleUpdateBlurb1(srv, logger)).HeadersRegexp("X-HTTP-Method-Override", "^PATCH$").Methods("POST")
	router.Handle("/v1beta1/{name:rooms/[^:]+/blurbs/[^:]+}", HandleDeleteBlurb(srv, logger)).Methods("DELETE")
	router.Handle("/v1beta1/{name:users/[^:]+/profile/blurbs/[^:]+}", HandleDeleteBlurb1(srv, logger)).Methods("DELETE")
	router.Handle("/v1beta1/{parent:rooms/[^:]+}/blurbs", HandleListBlurbs(srv, logger)).Methods("GET")
	router.Handle("/v1beta1/{parent:users/[^:]+/profile}/blurbs", HandleListBlurbs1(srv, logger)).Methods("GET")
	router.Handle("/v1beta1/{parent:rooms/[^:]+}/blurbs:search", HandleSearchBlurbs(srv, logger)).Methods("POST")
	router.Handle("/v1beta1/{parent:users/[^:]+/profile}/blurbs:search", HandleSearchBlurbs1(srv, logger)).Methods("POST")
	router.Handle("/v1beta1/{name:rooms/[^:]+}/blurbs:stream", HandleStreamBlurbs(srv, logger)).Methods("POST")
	router.Handle("/v1beta1/{name:users/[^:]+/profile}/blurbs:stream", HandleStreamBlurbs1(srv, logger)).Methods("POST")
	router.Handle("/v1beta1/{parent:rooms/[^:]+}/blurbs:send", HandleSendBlurbs(srv, logger)).Methods("POST")
	router.Handle("/v1beta1/{parent:users/[^:]+/profile}/blurbs:send", HandleSendBlurbs1(srv, logger)).Methods("POST")
}
