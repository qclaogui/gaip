// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package rest

import (
	"strings"
	"sync"

	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/encoding/protojson"
)

// ToJSON returns a copy of the current global JSON marshaling options in
// JSONMarshaler. Modifications to this copy do not change the values of these options returned in
// subsequent calls to this function. This is the function Showcase REST endpoints should use to
// handle JSON marshaling.
func ToJSON() *protojson.MarshalOptions {
	cp := *JSONMarshaler.current
	return &cp
}

// FromJSON returns a copy of the current global JSON unmarshaling options. Modifications to this copy
// do not change the values of these options returned in subsequent calls to this function. This is
// the function Showcase REST endpoints should use to handle JSON unmarshaling.
func FromJSON() *protojson.UnmarshalOptions {
	return &protojson.UnmarshalOptions{}
}

// JSONMarshaler captures the JSON marshaling options. It is intended only for tests of Showcase
// functionality (not for normal Showcase use or tests of generators against Showcase).
var JSONMarshaler JSONMarshalOptions

// JSONMarshalOptions contains the current JSON marshaling options used by REST endpoints, and
// allows for temporarily replacing these global options and then restoring them. This functionality
// is useful for some tests.
type JSONMarshalOptions struct {
	current, saved *protojson.MarshalOptions
	mu             sync.Mutex
}

// Replace replaces the current JSON marshaling options with those provided by opt. Call Restore()
// to restore the production options. Only one replacement can be active at a time; subsequent calls
// hang waiting for the first call's mutex to be released.
//
// As a special case, if opt==nil, the replacement is with the production options themselves; this
// is useful for tests that need to lock the production options to protect from other tests which
// may need to change them.
func (jm *JSONMarshalOptions) Replace(opt *protojson.MarshalOptions) {
	jm.mu.Lock()
	if opt == nil {
		opt = jm.current
	}
	jm.saved = jm.current
	jm.current = opt
}

// Restore restores the production JSON marshaling options. There must be an active Replace() called
// previously.
func (jm *JSONMarshalOptions) Restore() {
	jm.current = jm.saved
	jm.saved = nil
	jm.mu.Unlock()
}

// ToDottedLowerCamel converts each segment of a dot-delimited fieldPath to be individually lower-camel-cased; the dots are preserved.
func ToDottedLowerCamel(fieldPath string) string {
	parts := strings.Split(fieldPath, ".")
	for idx, segment := range parts {
		parts[idx] = strcase.ToLowerCamel(segment)
	}
	return strings.Join(parts, ".")
}

func init() {
	JSONMarshaler.current = &protojson.MarshalOptions{
		Multiline:       true,
		AllowPartial:    false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false, // we want lower-camel-cased field names
	}

}
