// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package gomodel

import (
	"fmt"
	"strings"
)

////////////////////////////////////////
// PathTemplate

// PathTemplate contains a sequence of parsed Segment to represent an HTTP binding.
type PathTemplate []*Segment

// NewPathTemplate parses `pattern` to return the corresponding PathTemplate.
func NewPathTemplate(pattern string) (PathTemplate, error) {
	return ParseTemplate(pattern)
}

// Flatten returns a flattened PathTemplate, which contains no recursively nested
// PathTemplate. Effectively, this removes any Segment with `Kind==Variable`.
func (pt PathTemplate) Flatten() PathTemplate {
	flat := PathTemplate{}
	for _, seg := range pt {
		flat = append(flat, seg.Flatten()...)
	}
	return flat
}

// HasVariables returns two booleans depending on whether `pt` has top-level and nested
// (lower-level) variables.
func (pt PathTemplate) HasVariables() (topVar, nestedVar bool) {
	for _, segment := range pt {
		if segment.Kind == Variable {
			segTopVar, segNestedVar := segment.Subsegments.HasVariables()
			nestedVar = nestedVar || segTopVar || segNestedVar
			topVar = true
		}
	}
	return topVar, nestedVar
}

// ListVariables returns a list of all the variables (proto field names) found in `pt`,
func (pt PathTemplate) ListVariables() []string {
	varNames := []string{}
	for _, segment := range pt {
		if segment.Kind == Variable {
			varNames = append(varNames, segment.Value)
			varNames = append(varNames, segment.Subsegments.ListVariables()...)
		}
	}
	return varNames
}

// asGoLiteral returns a Go-syntax representation of this PathTemplate. This is useful for
// constructing and debugging tests.
func (pt PathTemplate) asGoLiteral() string {
	parts := make([]string, len(pt))
	for idx, segment := range pt {
		parts[idx] = "&" + segment.asGoLiteral()
	}
	return fmt.Sprintf("PathTemplate{ %s }", strings.Join(parts, ", "))
}

////////////////////////////////////////
// Segment

// Segment is a single structural element in an HTTP binding
type Segment struct {
	Kind SegmentKind

	// the semantics of Value depend on Kind:
	// Kind==Variable: field path
	// Kind==Literal: literal value
	// Kind==SingleValue: "*"
	// Kind== MultipleValue: "**"
	Value string

	Subsegments PathTemplate
}

// String returns a string representation of this Segment.
func (seg *Segment) String() string {
	switch seg.Kind {
	case Literal:
		return fmt.Sprintf("%q", seg.Value)
	case SingleValue, MultipleValue:
		return seg.Value
	case Variable:
		subsegments := "!!ERROR: no subsegments"
		if len(seg.Subsegments) > 0 {
			subsegments = fmt.Sprintf("%s", seg.Subsegments)
		}
		return fmt.Sprintf("{%s = %s}", seg.Value, subsegments)
	default:
	}

	// Out of range: print as much info as possible
	return fmt.Sprintf("{%s(%d) %q %s}", seg.Kind, seg.Kind, seg.Value, seg.Subsegments)
}

// Flatten returns a flattened PathTemplate containing either this Segment or its flattened
// sub-segments.  Effectively, this removes any Segment with `Kind==Variable`.
func (seg *Segment) Flatten() PathTemplate {
	switch seg.Kind {
	case Variable:
		return seg.Subsegments.Flatten()
	default:
		return PathTemplate{seg}
	}
}

// asGoLiteral returns a Go-syntax representation of this Segment. This is useful for
// constructing and debugging tests.
func (seg *Segment) asGoLiteral() string {
	subsegments := "nil"
	if seg.Subsegments != nil {
		subsegments = seg.Subsegments.asGoLiteral()
	}

	return fmt.Sprintf("Segment{ %s, %q, %s }", seg.Kind.asGoLiteral(), seg.Value, subsegments)
}

////////////////////////////////////////
// SegmentKind

// SegmentKind describes a type of Segment.
type SegmentKind int

const (
	KindUndefined SegmentKind = iota
	Literal
	Variable
	SingleValue
	MultipleValue
	KindEnd
)

// Valid returns true iff this SegmentKind value is valid.
func (sk SegmentKind) Valid() bool {
	return sk > KindUndefined && sk < KindEnd
}

// String returns a string representation of this SegmentKind.
func (sk SegmentKind) String() string {
	names := []string{"(UNDEFINED)", "LITERAL", "VARIABLE", "SINGLEVAL", "MULTIVAL", "(END)"}
	if !sk.Valid() {
		return "INVALID"
	}
	return names[sk]
}

// asGoLiteral returns a Go-syntax representation of this SegmentKind. This is useful for
// constructing and debugging tests.
func (sk SegmentKind) asGoLiteral() string {
	names := []string{"KindUndefined", "Literal", "Variable", "SingleValue", "MultipleValue", "KindEnd"}
	return names[sk]
}
