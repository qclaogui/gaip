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
	"github.com/gorilla/mux"
	pb "github.com/qclaogui/gaip/genproto/showcase/apiv1beta1/showcasepb"
	"github.com/qclaogui/gaip/pkg/protocol/rest"
)

// HandleCreateRoom translates REST requests/responses on the wire to internal proto messages for CreateRoom
//
//	Generated for HTTP binding pattern: POST "/v1beta1/rooms"
func (srv *Server) HandleCreateRoom() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numURLPathParams := len(urlPathParams)

		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/rooms': %q", r.Method, r.URL))
		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("urlPathParams (expect 0, have %d): %q", numURLPathParams, urlPathParams))

		if numURLPathParams != 0 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 0, have %d: %#v", numURLPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &pb.CreateRoomRequest{}
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
		_ = level.Info(srv.logger).Log("msg", fmt.Sprintf("request: %s", requestJSON))

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
func (srv *Server) HandleGetRoom() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numURLPathParams := len(urlPathParams)

		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{name=rooms/*}': %q", r.Method, r.URL))
		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("urlPathParams (expect 1, have %d): %q", numURLPathParams, urlPathParams))

		if numURLPathParams != 1 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 1, have %d: %#v", numURLPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &pb.GetRoomRequest{}
		if err = rest.CheckRequestFormat(nil, r, request.ProtoReflect()); err != nil {
			rest.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
			return
		}

		if err = rest.PopulateSingularFields(request, urlPathParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
			return
		}

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

		_ = level.Info(srv.logger).Log("msg", fmt.Sprintf("request: %s", requestJSON))

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
func (srv *Server) HandleUpdateRoom() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numURLPathParams := len(urlPathParams)

		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{room.name=rooms/*}': %q", r.Method, r.URL))
		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("urlPathParams (expect 1, have %d): %q", numURLPathParams, urlPathParams))

		if numURLPathParams != 1 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 1, have %d: %#v", numURLPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &pb.UpdateRoomRequest{}
		// Intentional: Field values in the URL path override those set in the body.
		var bodyField pb.Room
		var jsonReader bytes.Buffer
		bodyReader := io.TeeReader(r.Body, &jsonReader)
		rBytes := make([]byte, r.ContentLength)
		if _, err = bodyReader.Read(rBytes); err != nil && !errors.Is(err, io.EOF) {
			rest.Error(w, http.StatusBadRequest, "error reading body content: %s", err)
			return
		}

		if err = rest.FromJSON().Unmarshal(rBytes, &bodyField); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading body params '*': %s", err)
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

		_ = level.Info(srv.logger).Log("msg", fmt.Sprintf("request: %s", requestJSON))

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
func (srv *Server) HandleDeleteRoom() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numURLPathParams := len(urlPathParams)

		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{name=rooms/*}': %q", r.Method, r.URL))
		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("urlPathParams (expect 1, have %d): %q", numURLPathParams, urlPathParams))

		if numURLPathParams != 1 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 1, have %d: %#v", numURLPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &pb.DeleteRoomRequest{}
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

		_ = level.Info(srv.logger).Log("msg", fmt.Sprintf("request: %s", requestJSON))

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
func (srv *Server) HandleListRooms() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numURLPathParams := len(urlPathParams)

		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/rooms': %q", r.Method, r.URL))
		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("urlPathParams (expect 0, have %d): %q", numURLPathParams, urlPathParams))

		if numURLPathParams != 0 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 0, have %d: %#v", numURLPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &pb.ListRoomsRequest{}
		if err = rest.CheckRequestFormat(nil, r, request.ProtoReflect()); err != nil {
			rest.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
			return
		}

		if err = rest.PopulateSingularFields(request, urlPathParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
			return
		}

		if err = rest.PopulateFields(request, queryParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading query params: %s", err)
			return
		}

		marshaler := rest.ToJSON()
		marshaler.UseEnumNumbers = systemParameters.EnumEncodingAsInt
		requestJSON, _ := marshaler.Marshal(request)

		_ = level.Info(srv.logger).Log("msg", fmt.Sprintf("request: %s", requestJSON))

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
func (srv *Server) HandleCreateBlurb() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numURLPathParams := len(urlPathParams)

		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{parent=rooms/*}/blurbs': %q", r.Method, r.URL))
		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("urlPathParams (expect 1, have %d): %q", numURLPathParams, urlPathParams))

		if numURLPathParams != 1 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 0, have %d: %#v", numURLPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &pb.CreateBlurbRequest{}
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
		_ = level.Info(srv.logger).Log("msg", fmt.Sprintf("request: %s", requestJSON))

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
func (srv *Server) HandleCreateBlurb1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numURLPathParams := len(urlPathParams)

		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/{parent=users/*/profile}/blurbs': %q", r.Method, r.URL))
		_ = level.Debug(srv.logger).Log("msg", fmt.Sprintf("urlPathParams (expect 1, have %d): %q", numURLPathParams, urlPathParams))

		if numURLPathParams != 1 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 0, have %d: %#v", numURLPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &pb.CreateBlurbRequest{}
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
		_ = level.Info(srv.logger).Log("msg", fmt.Sprintf("request: %s", requestJSON))

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
func (srv *Server) HandleGetBlurb() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rest.Error(w, http.StatusNotImplemented, "methods not implemented yet (request matched '/v1beta1/{name=rooms/*/blurbs/*}': %q)", r.URL)
	}
}

// HandleGetBlurb1 translates REST requests/responses on the wire to internal proto messages for GetBlurb
//
//	Generated for HTTP binding pattern: GET "/v1beta1/{name=users/*/profile/blurbs/*}"
func (srv *Server) HandleGetBlurb1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rest.Error(w, http.StatusNotImplemented, "methods not implemented yet (request matched '/v1beta1/{name=users/*/profile/blurbs/*}': %q)", r.URL)
	}
}

// HandleUpdateBlurb translates REST requests/responses on the wire to internal proto messages for UpdateBlurb
//
//	Generated for HTTP binding pattern: PATCH "/v1beta1/{blurb.name=rooms/*/blurbs/*}"
func (srv *Server) HandleUpdateBlurb() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rest.Error(w, http.StatusNotImplemented, "methods not implemented yet (request matched '/v1beta1/{blurb.name=rooms/*/blurbs/*}': %q)", r.URL)
	}
}

// HandleUpdateBlurb1 translates REST requests/responses on the wire to internal proto messages for UpdateBlurb
//
//	Generated for HTTP binding pattern: PATCH "/v1beta1/{blurb.name=users/*/profile/blurbs/*}"
func (srv *Server) HandleUpdateBlurb1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rest.Error(w, http.StatusNotImplemented, "methods not implemented yet (request matched '/v1beta1/{blurb.name=users/*/profile/blurbs/*}': %q)", r.URL)
	}
}

// HandleDeleteBlurb translates REST requests/responses on the wire to internal proto messages for DeleteBlurb
//
//	Generated for HTTP binding pattern: DELETE "/v1beta1/{name=rooms/*/blurbs/*}"
func (srv *Server) HandleDeleteBlurb() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rest.Error(w, http.StatusNotImplemented, "methods not implemented yet (request matched '/v1beta1/{name=rooms/*/blurbs/*}': %q)", r.URL)
	}
}

// HandleDeleteBlurb1 translates REST requests/responses on the wire to internal proto messages for DeleteBlurb
//
//	Generated for HTTP binding pattern: DELETE "/v1beta1/{name=users/*/profile/blurbs/*}"
func (srv *Server) HandleDeleteBlurb1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rest.Error(w, http.StatusNotImplemented, "methods not implemented yet (request matched '/v1beta1/{name=users/*/profile/blurbs/*}': %q)", r.URL)
	}
}

// HandleListBlurbs translates REST requests/responses on the wire to internal proto messages for ListBlurbs
//
//	Generated for HTTP binding pattern: GET "/v1beta1/{parent=rooms/*}/blurbs"
func (srv *Server) HandleListBlurbs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rest.Error(w, http.StatusNotImplemented, "methods not implemented yet (request matched '/v1beta1/{parent=rooms/*}/blurbs': %q)", r.URL)
	}
}

// HandleListBlurbs1 translates REST requests/responses on the wire to internal proto messages for ListBlurbs
//
//	Generated for HTTP binding pattern: GET "/v1beta1/{parent=users/*/profile}/blurbs"
func (srv *Server) HandleListBlurbs1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rest.Error(w, http.StatusNotImplemented, "methods not implemented yet (request matched '/v1beta1/{parent=users/*/profile}/blurbs': %q)", r.URL)
	}
}

// HandleSearchBlurbs translates REST requests/responses on the wire to internal proto messages for SearchBlurbs
//
//	Generated for HTTP binding pattern: POST "/v1beta1/{parent=rooms/*}/blurbs:search"
func (srv *Server) HandleSearchBlurbs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rest.Error(w, http.StatusNotImplemented, "methods not implemented yet (request matched '/v1beta1/{parent=rooms/*}/blurbs:search': %q)", r.URL)
	}
}

// HandleSearchBlurbs1 translates REST requests/responses on the wire to internal proto messages for SearchBlurbs
//
//	Generated for HTTP binding pattern: POST "/v1beta1/{parent=users/*/profile}/blurbs:search"
func (srv *Server) HandleSearchBlurbs1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rest.Error(w, http.StatusNotImplemented, "methods not implemented yet (request matched '/v1beta1/{parent=users/*/profile}/blurbs:search': %q)", r.URL)
	}
}

// HandleStreamBlurbs translates REST requests/responses on the wire to internal proto messages for StreamBlurbs
//
//	Generated for HTTP binding pattern: POST "/v1beta1/{name=rooms/*}/blurbs:stream"
func (srv *Server) HandleStreamBlurbs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rest.Error(w, http.StatusNotImplemented, "methods not implemented yet (request matched '/v1beta1/{name=users/*/profile}/blurbs:stream': %q)", r.URL)
	}
}

// HandleStreamBlurbs1 translates REST requests/responses on the wire to internal proto messages for StreamBlurbs
//
//	Generated for HTTP binding pattern: POST "/v1beta1/{name=users/*/profile}/blurbs:stream"
func (srv *Server) HandleStreamBlurbs1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rest.Error(w, http.StatusNotImplemented, "methods not implemented yet (request matched '/v1beta1/{name=users/*/profile}/blurbs:stream': %q)", r.URL)
	}
}

// HandleSendBlurbs translates REST requests/responses on the wire to internal proto messages for SendBlurbs
//
//	Generated for HTTP binding pattern: POST "/v1beta1/{parent=rooms/*}/blurbs:send"
func (srv *Server) HandleSendBlurbs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rest.Error(w, http.StatusNotImplemented, "client-streaming methods not implemented yet (request matched '/v1beta1/{parent=rooms/*}/blurbs:send': %q)", r.URL)
	}
}

// HandleSendBlurbs1 translates REST requests/responses on the wire to internal proto messages for SendBlurbs
//
//	Generated for HTTP binding pattern: POST "/v1beta1/{parent=users/*/profile}/blurbs:send"
func (srv *Server) HandleSendBlurbs1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rest.Error(w, http.StatusNotImplemented, "client-streaming methods not implemented yet (request matched '/v1beta1/{parent=users/*/profile}/blurbs:send': %q)", r.URL)
	}
}
