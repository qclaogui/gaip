// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

// DO NOT EDIT. This is an auto-generated file containing the REST handlers
// for service #0: "LibraryService" (.qclaogui.library.v1.LibraryService).

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
	librarypb "github.com/qclaogui/gaip/genproto/library/apiv1/librarypb"
	"github.com/qclaogui/gaip/pkg/protocol/rest"
)

// HandleCreateShelf translates REST requests/responses on the wire to internal proto messages for CreateShelf
//
//	Generated for HTTP binding pattern: POST "/v1/shelves"
func HandleCreateShelf(srv librarypb.LibraryServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1/shelves': %q", r.Method, r.URL))
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

		request := &librarypb.CreateShelfRequest{}
		// Intentional: Field values in the URL path override those set in the body.
		var bodyField librarypb.Shelf
		var jsonReader bytes.Buffer
		bodyReader := io.TeeReader(r.Body, &jsonReader)
		rBytes := make([]byte, r.ContentLength)
		if _, err = bodyReader.Read(rBytes); err != nil && err != io.EOF {
			rest.Error(w, http.StatusBadRequest, "error reading body content: %s", err)
			return
		}

		if err = rest.FromJSON().Unmarshal(rBytes, &bodyField); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading body into request field 'shelf': %s", err)
			return
		}

		if err = rest.CheckRequestFormat(&jsonReader, r, request.ProtoReflect()); err != nil {
			rest.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
			return
		}
		request.Shelf = &bodyField

		if err = rest.PopulateSingularFields(request, urlPathParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
			return
		}

		// TODO: Decide whether query-param value or URL-path value takes precedence when a field appears in both
		excludedQueryParams := []string{"shelf"}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1/shelves")
		response, err := srv.CreateShelf(ctx, request)
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

// HandleGetShelf translates REST requests/responses on the wire to internal proto messages for GetShelf
//
//	Generated for HTTP binding pattern: GET "/v1/{name=shelves/*}"
func HandleGetShelf(srv librarypb.LibraryServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1/{name=shelves/*}': %q", r.Method, r.URL))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlPathParams (expect 1, have %d): %q", numUrlPathParams, urlPathParams))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlRequestHeaders: %s", rest.PrettyPrintHeaders(r, "    ")))

		rest.IncludeRequestHeadersInResponse(w, r)

		if numUrlPathParams != 1 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 1, have %d: %#v", numUrlPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &librarypb.GetShelfRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1/{name=shelves/*}")
		response, err := srv.GetShelf(ctx, request)
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

// HandleListShelves translates REST requests/responses on the wire to internal proto messages for ListShelves
//
//	Generated for HTTP binding pattern: GET "/v1/shelves"
func HandleListShelves(srv librarypb.LibraryServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1/shelves': %q", r.Method, r.URL))
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

		request := &librarypb.ListShelvesRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1/shelves")
		response, err := srv.ListShelves(ctx, request)
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

// HandleDeleteShelf translates REST requests/responses on the wire to internal proto messages for DeleteShelf
//
//	Generated for HTTP binding pattern: DELETE "/v1/{name=shelves/*}"
func HandleDeleteShelf(srv librarypb.LibraryServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1/{name=shelves/*}': %q", r.Method, r.URL))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlPathParams (expect 1, have %d): %q", numUrlPathParams, urlPathParams))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlRequestHeaders: %s", rest.PrettyPrintHeaders(r, "    ")))

		rest.IncludeRequestHeadersInResponse(w, r)

		if numUrlPathParams != 1 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 1, have %d: %#v", numUrlPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &librarypb.DeleteShelfRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1/{name=shelves/*}")
		response, err := srv.DeleteShelf(ctx, request)
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

// HandleMergeShelves translates REST requests/responses on the wire to internal proto messages for MergeShelves
//
//	Generated for HTTP binding pattern: POST "/v1/{name=shelves/*}:merge"
func HandleMergeShelves(srv librarypb.LibraryServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1/{name=shelves/*}:merge': %q", r.Method, r.URL))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlPathParams (expect 1, have %d): %q", numUrlPathParams, urlPathParams))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlRequestHeaders: %s", rest.PrettyPrintHeaders(r, "    ")))

		rest.IncludeRequestHeadersInResponse(w, r)

		if numUrlPathParams != 1 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 1, have %d: %#v", numUrlPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &librarypb.MergeShelvesRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1/{name=shelves/*}:merge")
		response, err := srv.MergeShelves(ctx, request)
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

// HandleCreateBook translates REST requests/responses on the wire to internal proto messages for CreateBook
//
//	Generated for HTTP binding pattern: POST "/v1/{parent=shelves/*}/books"
func HandleCreateBook(srv librarypb.LibraryServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1/{parent=shelves/*}/books': %q", r.Method, r.URL))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlPathParams (expect 1, have %d): %q", numUrlPathParams, urlPathParams))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlRequestHeaders: %s", rest.PrettyPrintHeaders(r, "    ")))

		rest.IncludeRequestHeadersInResponse(w, r)

		if numUrlPathParams != 1 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 1, have %d: %#v", numUrlPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &librarypb.CreateBookRequest{}
		// Intentional: Field values in the URL path override those set in the body.
		var bodyField librarypb.Book
		var jsonReader bytes.Buffer
		bodyReader := io.TeeReader(r.Body, &jsonReader)
		rBytes := make([]byte, r.ContentLength)
		if _, err = bodyReader.Read(rBytes); err != nil && err != io.EOF {
			rest.Error(w, http.StatusBadRequest, "error reading body content: %s", err)
			return
		}

		if err = rest.FromJSON().Unmarshal(rBytes, &bodyField); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading body into request field 'book': %s", err)
			return
		}

		if err = rest.CheckRequestFormat(&jsonReader, r, request.ProtoReflect()); err != nil {
			rest.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
			return
		}
		request.Book = &bodyField

		if err = rest.PopulateSingularFields(request, urlPathParams); err != nil {
			rest.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
			return
		}

		// TODO: Decide whether query-param value or URL-path value takes precedence when a field appears in both
		excludedQueryParams := []string{"book", "parent"}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1/{parent=shelves/*}/books")
		response, err := srv.CreateBook(ctx, request)
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

// HandleGetBook translates REST requests/responses on the wire to internal proto messages for GetBook
//
//	Generated for HTTP binding pattern: GET "/v1/{name=shelves/*/books/*}"
func HandleGetBook(srv librarypb.LibraryServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1/{name=shelves/*/books/*}': %q", r.Method, r.URL))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlPathParams (expect 1, have %d): %q", numUrlPathParams, urlPathParams))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlRequestHeaders: %s", rest.PrettyPrintHeaders(r, "    ")))

		rest.IncludeRequestHeadersInResponse(w, r)

		if numUrlPathParams != 1 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 1, have %d: %#v", numUrlPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &librarypb.GetBookRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1/{name=shelves/*/books/*}")
		response, err := srv.GetBook(ctx, request)
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

// HandleListBooks translates REST requests/responses on the wire to internal proto messages for ListBooks
//
//	Generated for HTTP binding pattern: GET "/v1/{parent=shelves/*}/books"
func HandleListBooks(srv librarypb.LibraryServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1/{parent=shelves/*}/books': %q", r.Method, r.URL))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlPathParams (expect 1, have %d): %q", numUrlPathParams, urlPathParams))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlRequestHeaders: %s", rest.PrettyPrintHeaders(r, "    ")))

		rest.IncludeRequestHeadersInResponse(w, r)

		if numUrlPathParams != 1 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 1, have %d: %#v", numUrlPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &librarypb.ListBooksRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1/{parent=shelves/*}/books")
		response, err := srv.ListBooks(ctx, request)
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

// HandleDeleteBook translates REST requests/responses on the wire to internal proto messages for DeleteBook
//
//	Generated for HTTP binding pattern: DELETE "/v1/{name=shelves/*/books/*}"
func HandleDeleteBook(srv librarypb.LibraryServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1/{name=shelves/*/books/*}': %q", r.Method, r.URL))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlPathParams (expect 1, have %d): %q", numUrlPathParams, urlPathParams))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlRequestHeaders: %s", rest.PrettyPrintHeaders(r, "    ")))

		rest.IncludeRequestHeadersInResponse(w, r)

		if numUrlPathParams != 1 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 1, have %d: %#v", numUrlPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &librarypb.DeleteBookRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1/{name=shelves/*/books/*}")
		response, err := srv.DeleteBook(ctx, request)
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

// HandleUpdateBook translates REST requests/responses on the wire to internal proto messages for UpdateBook
//
//	Generated for HTTP binding pattern: PATCH "/v1/{book.name=shelves/*/books/*}"
func HandleUpdateBook(srv librarypb.LibraryServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1/{book.name=shelves/*/books/*}': %q", r.Method, r.URL))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlPathParams (expect 1, have %d): %q", numUrlPathParams, urlPathParams))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlRequestHeaders: %s", rest.PrettyPrintHeaders(r, "    ")))

		rest.IncludeRequestHeadersInResponse(w, r)

		if numUrlPathParams != 1 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 1, have %d: %#v", numUrlPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &librarypb.UpdateBookRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1/{book.name=shelves/*/books/*}")
		response, err := srv.UpdateBook(ctx, request)
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

// HandleMoveBook translates REST requests/responses on the wire to internal proto messages for MoveBook
//
//	Generated for HTTP binding pattern: POST "/v1/{name=shelves/*/books/*}:move"
func HandleMoveBook(srv librarypb.LibraryServiceServer, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathParams := mux.Vars(r)
		numUrlPathParams := len(urlPathParams)

		_ = level.Debug(logger).Log("msg", fmt.Sprintf("Received %s request matching '/v1/{name=shelves/*/books/*}:move': %q", r.Method, r.URL))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlPathParams (expect 1, have %d): %q", numUrlPathParams, urlPathParams))
		_ = level.Debug(logger).Log("msg", fmt.Sprintf("urlRequestHeaders: %s", rest.PrettyPrintHeaders(r, "    ")))

		rest.IncludeRequestHeadersInResponse(w, r)

		if numUrlPathParams != 1 {
			rest.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 1, have %d: %#v", numUrlPathParams, urlPathParams)
			return
		}

		systemParameters, queryParams, err := rest.GetSystemParameters(r)
		if err != nil {
			rest.Error(w, http.StatusBadRequest, "error in query string: %s", err)
			return
		}

		request := &librarypb.MoveBookRequest{}
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

		ctx := context.WithValue(r.Context(), rest.BindingURIKey, "/v1/{name=shelves/*/books/*}:move")
		response, err := srv.MoveBook(ctx, request)
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

// RegisterHandlersLibraryService register REST requests/responses on the wire to internal proto messages for
func RegisterHandlersLibraryService(router *mux.Router, srv librarypb.LibraryServiceServer, logger log.Logger) {
	router.Handle("/v1/shelves", HandleCreateShelf(srv, logger)).Methods("POST")
	router.Handle("/v1/{name:shelves/[^:]+}", HandleGetShelf(srv, logger)).Methods("GET")
	router.Handle("/v1/shelves", HandleListShelves(srv, logger)).Methods("GET")
	router.Handle("/v1/{name:shelves/[^:]+}", HandleDeleteShelf(srv, logger)).Methods("DELETE")
	router.Handle("/v1/{name:shelves/[^:]+}:merge", HandleMergeShelves(srv, logger)).Methods("POST")
	router.Handle("/v1/{parent:shelves/[^:]+}/books", HandleCreateBook(srv, logger)).Methods("POST")
	router.Handle("/v1/{name:shelves/[^:]+/books/[^:]+}", HandleGetBook(srv, logger)).Methods("GET")
	router.Handle("/v1/{parent:shelves/[^:]+}/books", HandleListBooks(srv, logger)).Methods("GET")
	router.Handle("/v1/{name:shelves/[^:]+/books/[^:]+}", HandleDeleteBook(srv, logger)).Methods("DELETE")
	router.Handle("/v1/{book.name:shelves/[^:]+/books/[^:]+}", HandleUpdateBook(srv, logger)).Methods("PATCH")
	router.Handle("/v1/{book.name:shelves/[^:]+/books/[^:]+}", HandleUpdateBook(srv, logger)).HeadersRegexp("X-HTTP-Method-Override", "^PATCH$").Methods("POST")
	router.Handle("/v1/{name:shelves/[^:]+/books/[^:]+}:move", HandleMoveBook(srv, logger)).Methods("POST")
}
