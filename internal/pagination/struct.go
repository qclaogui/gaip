// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package pagination

import (
	"encoding/base64"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"strings"
)

// DecodePageToken decodes an encoded page token into an arbitrary struct.
func DecodePageToken(s string, v interface{}) error {
	dec := gob.NewDecoder(base64.NewDecoder(base64.URLEncoding, strings.NewReader(s)))
	if err := dec.Decode(v); err != nil && !errors.Is(err, io.EOF) {
		return fmt.Errorf("decode page token struct: %w", err)
	}
	return nil
}

// EncodePageToken encodes an arbitrary struct as a page token.
func EncodePageToken(v interface{}) string {
	var b strings.Builder
	base64Encoder := base64.NewEncoder(base64.URLEncoding, &b)
	gobEncoder := gob.NewEncoder(base64Encoder)
	_ = gobEncoder.Encode(v)
	_ = base64Encoder.Close()
	return b.String()
}
