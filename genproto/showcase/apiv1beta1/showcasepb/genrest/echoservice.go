// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

// DO NOT EDIT. This is an auto-generated file containing the REST handlers
// for service #0: "EchoService" (.qclaogui.showcase.v1beta1.EchoService).

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
	showcasepb "github.com/qclaogui/gaip/genproto/showcase/apiv1beta1/showcasepb"
	"github.com/qclaogui/gaip/pkg/protocol/rest"
)

// HandleEcho translates REST requests/responses on the wire to internal proto messages for Echo
//
//	Generated for HTTP binding pattern: POST "/v1beta1/echo:echo"
func HandleEcho(srv showcasepb.EchoServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/echo:echo': %q", r.Method, r.URL))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlPathParams (expect 0, have %d): %q", numUrlPathParams, urlPathParams))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlRequestHeaders: %s", rest.PrettyPrintHeaders(r, "    ")))

		rest.IncludeRequestHeadersInResponse(w, r)

		if numUrlPathParams != 0 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 0, have %d: %#v", numUrlPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &showcasepb.EchoRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/echo:echo")
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
//	Generated for HTTP binding pattern: POST "/v1beta1/echo:error-details"
func HandleEchoErrorDetails(srv showcasepb.EchoServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/echo:error-details': %q", r.Method, r.URL))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlPathParams (expect 0, have %d): %q", numUrlPathParams, urlPathParams))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlRequestHeaders: %s", rest.PrettyPrintHeaders(r, "    ")))

		rest.IncludeRequestHeadersInResponse(w, r)

		if numUrlPathParams != 0 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 0, have %d: %#v", numUrlPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &showcasepb.EchoErrorDetailsRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/echo:error-details")
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

// HandleFailEchoWithDetails translates REST requests/responses on the wire to internal proto messages for FailEchoWithDetails
//
//	Generated for HTTP binding pattern: POST "/v1beta1/echo:failWithDetails"
func HandleFailEchoWithDetails(srv showcasepb.EchoServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/echo:failWithDetails': %q", r.Method, r.URL))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlPathParams (expect 0, have %d): %q", numUrlPathParams, urlPathParams))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlRequestHeaders: %s", rest.PrettyPrintHeaders(r, "    ")))

		rest.IncludeRequestHeadersInResponse(w, r)

		if numUrlPathParams != 0 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 0, have %d: %#v", numUrlPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &showcasepb.FailEchoWithDetailsRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/echo:failWithDetails")
		response, err := srv.FailEchoWithDetails(ctx, request)
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

// HandleExpand translates REST requests/responses on the wire to internal proto messages for Expand
//
//	Generated for HTTP binding pattern: POST "/v1beta1/echo:expand"
func HandleExpand(srv showcasepb.EchoServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/echo:expand': %q", r.Method, r.URL))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlPathParams (expect 0, have %d): %q", numUrlPathParams, urlPathParams))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlRequestHeaders: %s", rest.PrettyPrintHeaders(r, "    ")))

		rest.IncludeRequestHeadersInResponse(w, r)

		if numUrlPathParams != 0 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 0, have %d: %#v", numUrlPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &showcasepb.ExpandRequest{}
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
		streamer := &EchoServiceExpandServer{serverStreamer}
		if err = srv.Expand(request, streamer); err != nil {
			rest.ReportGRPCError(w, err)
		}
	}
}

// HandleCollect translates REST requests/responses on the wire to internal proto messages for Collect
//
//	Generated for HTTP binding pattern: POST "/v1beta1/echo:collect"
func HandleCollect(srv showcasepb.EchoServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = srv
		_ = level.Warn(logger).Log("msg", fmt.Sprintf("client-streaming methods not implemented yet (request matched '/v1beta1/echo:collect': %q)", r.URL))
		rest.Error(w, http.StatusNotImplemented, "client-streaming methods not implemented yet (request matched '/v1beta1/echo:collect': %q)", r.URL)
	}
}

// HandlePagedExpand translates REST requests/responses on the wire to internal proto messages for PagedExpand
//
//	Generated for HTTP binding pattern: POST "/v1beta1/echo:pagedExpand"
func HandlePagedExpand(srv showcasepb.EchoServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/echo:pagedExpand': %q", r.Method, r.URL))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlPathParams (expect 0, have %d): %q", numUrlPathParams, urlPathParams))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlRequestHeaders: %s", rest.PrettyPrintHeaders(r, "    ")))

		rest.IncludeRequestHeadersInResponse(w, r)

		if numUrlPathParams != 0 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 0, have %d: %#v", numUrlPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &showcasepb.PagedExpandRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/echo:pagedExpand")
		response, err := srv.PagedExpand(ctx, request)
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

// HandleWait translates REST requests/responses on the wire to internal proto messages for Wait
//
//	Generated for HTTP binding pattern: POST "/v1beta1/echo:wait"
func HandleWait(srv showcasepb.EchoServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/echo:wait': %q", r.Method, r.URL))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlPathParams (expect 0, have %d): %q", numUrlPathParams, urlPathParams))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlRequestHeaders: %s", rest.PrettyPrintHeaders(r, "    ")))

		rest.IncludeRequestHeadersInResponse(w, r)

		if numUrlPathParams != 0 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 0, have %d: %#v", numUrlPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &showcasepb.WaitRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/echo:wait")
		response, err := srv.Wait(ctx, request)
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

// HandleBlock translates REST requests/responses on the wire to internal proto messages for Block
//
//	Generated for HTTP binding pattern: POST "/v1beta1/echo:block"
func HandleBlock(srv showcasepb.EchoServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1beta1/echo:block': %q", r.Method, r.URL))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlPathParams (expect 0, have %d): %q", numUrlPathParams, urlPathParams))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlRequestHeaders: %s", rest.PrettyPrintHeaders(r, "    ")))

		rest.IncludeRequestHeadersInResponse(w, r)

		if numUrlPathParams != 0 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 0, have %d: %#v", numUrlPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &showcasepb.BlockRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1beta1/echo:block")
		response, err := srv.Block(ctx, request)
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

// EchoServiceExpandServer implements showcasepb.EchoServiceExpandServer to provide server-side streaming over REST, returning all the
// individual responses as part of a long JSON list.
type EchoServiceExpandServer struct {
	*rest.ServerStreamer
}

// Send accumulates a response to be fetched later as part of response list returned over REST.
func (streamer *EchoServiceExpandServer) Send(response *showcasepb.EchoResponse) error {
	return streamer.ServerStreamer.Send(response)
}

// RegisterHandlersEchoService register REST requests/responses on the wire to internal proto messages for
func RegisterHandlersEchoService(router *mux.Router, srv showcasepb.EchoServiceServer, logger log.Logger) {
	router.Handle("/v1beta1/echo:echo", HandleEcho(srv, logger)).Methods("POST")
	router.Handle("/v1beta1/echo:error-details", HandleEchoErrorDetails(srv, logger)).Methods("POST")
	router.Handle("/v1beta1/echo:failWithDetails", HandleFailEchoWithDetails(srv, logger)).Methods("POST")
	router.Handle("/v1beta1/echo:expand", HandleExpand(srv, logger)).Methods("POST")
	router.Handle("/v1beta1/echo:collect", HandleCollect(srv, logger)).Methods("POST")
	router.Handle("/v1beta1/echo:pagedExpand", HandlePagedExpand(srv, logger)).Methods("POST")
	router.Handle("/v1beta1/echo:wait", HandleWait(srv, logger)).Methods("POST")
	router.Handle("/v1beta1/echo:block", HandleBlock(srv, logger)).Methods("POST")
}
