// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package memory

import (
	"slices"

	"google.golang.org/protobuf/reflect/protoreflect"
)

// applyFieldMask applies the values from the src message to the values of the
// dst message according to the contents of the given field mask paths.
// If paths is empty/nil, or contains *, it is considered a full update.
//
// TODO: Does not support nested message paths. Currently only used with flat
// resource messages.
func applyFieldMask(src, dst protoreflect.Message, paths []string) {
	fullUpdate := len(paths) == 0 || slices.Contains(paths, "*")

	fields := dst.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		isOneof := field.ContainingOneof() != nil && !field.ContainingOneof().IsSynthetic()

		// Set field in dst with value from src, skipping true oneofs, while
		// handling proto3_optional fields represented as synthetic oneofs.
		if (fullUpdate || slices.Contains(paths, string(field.Name()))) && !isOneof {
			dst.Set(field, src.Get(field))
		}
	}

	oneofs := dst.Descriptor().Oneofs()
	for i := 0; i < oneofs.Len(); i++ {
		oneof := oneofs.Get(i)
		// Skip proto3_optional synthetic oneofs.
		if oneof.IsSynthetic() {
			continue
		}

		setOneof := src.WhichOneof(oneof)
		if setOneof == nil && fullUpdate {
			// Full update with no field set in this oneof of
			// src means clear all fields for this oneof in dst.
			fields = oneof.Fields()
			for j := 0; j < fields.Len(); j++ {
				dst.Clear(fields.Get(j))
			}
		} else if setOneof != nil && (fullUpdate || slices.Contains(paths, string(setOneof.Name()))) {
			// Full update or targeted updated with a field set in this oneof of
			// src means set that field for the same oneof in dst, which implicitly
			// clears any previously set field for this oneof.
			dst.Set(setOneof, src.Get(setOneof))
		}
	}
}
