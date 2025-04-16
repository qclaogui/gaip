// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package thirdparty

import "testing"

func TestOpenAPIFS(t *testing.T) {
	for _, f := range []string{
		"gen/openapiv2/index.css",
		"gen/openapiv2/index.html",
	} {
		if _, err := OpenAPI.Open(f); err != nil {
			t.Errorf("\nOops 🔥\x1b[91m Failed asserting that\x1b[39m\n"+
				"✘got: %v\n\x1b[92m"+
				"want: %v\x1b[39m", err, nil)
		}
	}
}
