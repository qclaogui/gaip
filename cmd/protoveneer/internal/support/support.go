// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package support provides support functions for protoveneer. The protoveneer binary
// embeds it and extracts the needed functions when generating code.
//
// This package should not be imported. It is written as an ordinary Go package so
// it can be edited and tested with standard tools.
//
// The symbols begin with "pv" to reduce the chance of collision when the generated
// code is combined with user-written code in the same package.
package support

import (
	"fmt"
	"time"

	"cloud.google.com/go/civil"
	"github.com/googleapis/gax-go/v2/apierror"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/genproto/googleapis/type/date"
	gstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// pvTransformSlice applies f to each element of from and returns
// a new slice with the results.
func pvTransformSlice[From, To any](from []From, f func(From) To) []To {
	if from == nil {
		return nil
	}
	to := make([]To, len(from))
	for i, e := range from {
		to[i] = f(e)
	}
	return to
}

// pvTransformMapValues applies f to each value of from, returning a new map.
// It does not change the keys.
func pvTransformMapValues[K comparable, VFrom, VTo any](from map[K]VFrom, f func(VFrom) VTo) map[K]VTo {
	if from == nil {
		return nil
	}
	to := map[K]VTo{}
	for k, v := range from {
		to[k] = f(v)
	}
	return to
}

// pvAddrOrNil returns nil if x is the zero value for T,
// or &x otherwise.
func pvAddrOrNil[T comparable](x T) *T {
	var z T
	if x == z {
		return nil
	}
	return &x
}

// pvDerefOrZero returns the zero value for T if x is nil,
// or *x otherwise.
func pvDerefOrZero[T any](x *T) T {
	if x == nil {
		var z T
		return z
	}
	return *x
}

// pvCivilDateToProto converts a civil.Date to a date.Date.
func pvCivilDateToProto(d civil.Date) *date.Date {
	return &date.Date{
		Year:  int32(d.Year),
		Month: int32(d.Month),
		Day:   int32(d.Day),
	}
}

// pvCivilDateFromProto converts a date.Date to a civil.Date.
func pvCivilDateFromProto(p *date.Date) civil.Date {
	if p == nil {
		return civil.Date{}
	}
	return civil.Date{
		Year:  int(p.Year),
		Month: time.Month(p.Month),
		Day:   int(p.Day),
	}
}

// pvMapToStructPB converts a map into a structpb.Struct.
func pvMapToStructPB(m map[string]any) *structpb.Struct {
	if m == nil {
		return nil
	}
	s, err := structpb.NewStruct(m)
	if err != nil {
		panic(pvPanic(fmt.Errorf("pvMapToStructPB: %w", err)))
	}
	return s
}

// pvMapFromStructPB converts a structpb.Struct to a map.
func pvMapFromStructPB(p *structpb.Struct) map[string]any {
	if p == nil {
		return nil
	}
	return p.AsMap()
}

// pvTimeToProto converts a time.Time into a Timestamp.
func pvTimeToProto(t time.Time) *timestamppb.Timestamp {
	if t.IsZero() {
		return nil
	}
	return timestamppb.New(t)
}

// pvTimeFromProto converts a Timestamp into a time.Time.
func pvTimeFromProto(ts *timestamppb.Timestamp) time.Time {
	if ts == nil {
		return time.Time{}
	}
	return ts.AsTime()
}

// pvAPIErrorToProto converts an APIError to a proto Status.
func pvAPIErrorToProto(ae *apierror.APIError) *spb.Status {
	if ae == nil {
		return nil
	}
	return ae.GRPCStatus().Proto()
}

// pvAPIErrorFromProto converts a proto Status to an APIError.
func pvAPIErrorFromProto(s *spb.Status) *apierror.APIError {
	err := gstatus.ErrorProto(s)
	aerr, ok := apierror.ParseError(err, true)
	if !ok {
		// Should be impossible.
		return nil
	}
	return aerr
}

// pvDurationFromProto converts a Duration proto to a time.Duration.
func pvDurationFromProto(d *durationpb.Duration) time.Duration {
	if d == nil {
		return 0
	}
	return d.AsDuration()
}

// So callers can distinguish pv function panics from other panics.
// Keep in sync with the type generated by protoveneer.
type pvPanic error
