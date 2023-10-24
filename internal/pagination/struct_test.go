// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package pagination

import (
	"testing"
)

func Test_PageTokenStruct(t *testing.T) {
	t.Parallel()
	type pageToken struct {
		Int    int
		String string
	}

	tests := []struct {
		desc string
		in   pageToken
	}{
		{
			desc: "all set",
			in: pageToken{
				Int:    42,
				String: "foo",
			},
		},
	}

	for _, test := range tests {
		tt := test
		t.Run(tt.desc, func(t *testing.T) {
			str := EncodePageToken(test.in)
			var out pageToken
			if err := DecodePageToken(str, &out); err != nil {
				t.Errorf("\nOops 🔥\x1b[91m Failed asserting that\x1b[39m\n"+
					"✘got: %v\n\x1b[92m"+
					"want: %v\x1b[39m", err, nil)
			}
			if tt.in != out {
				t.Errorf("\nOops 🔥\x1b[91m Failed asserting that\x1b[39m\n"+
					"✘got: %v\n\x1b[92m"+
					"want: %v\x1b[39m", false, "in==out")
			}
		})
	}
}
