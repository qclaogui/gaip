// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package pagination

import (
	"testing"

	"github.com/qclaogui/gaip/genproto/library/apiv1/librarypb"
)

func TestCalculateRequestChecksum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc     string
		request1 Request
		request2 Request
		equal    bool
	}{
		{
			desc: "same request",
			request1: &librarypb.ListBooksRequest{
				Parent:    "shelves/1",
				PageSize:  100,
				PageToken: "token",
			},
			request2: &librarypb.ListBooksRequest{
				Parent:    "shelves/1",
				PageSize:  100,
				PageToken: "token",
			},
			equal: true,
		},
		{
			desc: "different parents",
			request1: &librarypb.ListBooksRequest{
				Parent:    "shelves/1",
				PageSize:  100,
				PageToken: "token",
			},
			request2: &librarypb.ListBooksRequest{
				Parent:    "shelves/2",
				PageSize:  100,
				PageToken: "token",
			},
			equal: false,
		},
		{
			desc: "different page sizes",
			request1: &librarypb.ListBooksRequest{
				Parent:    "shelves/1",
				PageSize:  100,
				PageToken: "token",
			},
			request2: &librarypb.ListBooksRequest{
				Parent:    "shelves/1",
				PageSize:  200,
				PageToken: "token",
			},
			equal: true,
		},
		{
			desc: "different page sizes",
			request1: &librarypb.ListBooksRequest{
				Parent:    "shelves/1",
				PageSize:  100,
				PageToken: "token",
			},
			request2: &librarypb.ListBooksRequest{
				Parent:    "shelves/1",
				PageSize:  100,
				PageToken: "token2",
			},
			equal: true,
		},
	}

	for _, test := range tests {
		tt := test
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()
			checksum1, err := calculateRequestChecksum(tt.request1)
			if err != nil {
				t.Errorf("\nOops 🔥\x1b[91m Failed asserting that\x1b[39m\n"+
					"✘got: %v\n\x1b[92m"+
					"want: %v\x1b[39m", err, nil)
			}
			checksum2, err := calculateRequestChecksum(tt.request2)
			if err != nil {
				t.Errorf("\nOops 🔥\x1b[91m Failed asserting that\x1b[39m\n"+
					"✘got: %v\n\x1b[92m"+
					"want: %v\x1b[39m", err, nil)
			}

			if tt.equal {
				// checksum1 should equal checksum2
				if checksum1 != checksum2 {
					t.Errorf("\nOops 🔥\x1b[91m Failed asserting that\x1b[39m\n"+
						"✘got: %v\n\x1b[92m"+
						"want: %v\x1b[39m", false, true)
				}
			} else {
				if checksum1 == checksum2 {
					t.Errorf("\nOops 🔥\x1b[91m Failed asserting that\x1b[39m\n"+
						"✘got: %v\n\x1b[92m"+
						"want: %v\x1b[39m", true, false)
				}
			}
		})
	}
}
