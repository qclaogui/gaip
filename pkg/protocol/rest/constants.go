// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package rest

import "fmt"

const (
	// CharsField contains the characters allowed in a field name (URL path or body)
	CharsField = `-_.0-9a-zA-Z`

	// CharsLiteral contains the the characters allowed in a URL path literal segment.
	CharsLiteral = `-_.~0-9a-zA-Z%`

	// RegexURLPathSingleSegmentValue contains the regex expression for matching a single URL
	// path segment (i.e. `/` is not allowed). Since gorilla/mux hands uses the decoded paths to
	// match, we can just accept any characters.
	RegexURLPathSingleSegmentValue = "[^:]+"

	// RegexURLPathMultipleSegmentValue contains the regex expression for matching multiple URL
	// path segments (i.e. `/` is allowed). Since gorilla/mux hands uses the decoded paths to
	// match, we can just accept any characters.
	RegexURLPathMultipleSegmentValue = "[^:]+"
)

var (
	// CharsFieldPath contains the characters allowed in a dotted field path.
	CharsFieldPath string

	// RegexField contains the regex expression for matching a single (not nested) field name.
	RegexField string

	// RegexFieldPath contains the regex expression for matching a dotted field path.
	RegexFieldPath string

	// RegexLiteral contains the regex expression for matching a URL path literal segment.
	RegexLiteral string
)

func init() {
	RegexField = fmt.Sprintf(`[%s]+`, CharsField)

	CharsFieldPath = CharsField + `.`
	RegexFieldPath = fmt.Sprintf(`^[%s]+`, CharsFieldPath)

	RegexLiteral = fmt.Sprintf(`^[%s]+`, CharsLiteral)
}

// BindingURIKeyType A key-type for storing binding URI value in the Context
type BindingURIKeyType string

const BindingURIKey = BindingURIKeyType("BindingURIKey")
