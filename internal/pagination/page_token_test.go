// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package pagination

import (
	"testing"

	"github.com/qclaogui/golang-api-server/genproto/library/apiv1/librarypb"
)

func TestParsePageToken(t *testing.T) {
	t.Parallel()

	t.Run("valid checksums", func(t *testing.T) {
		request1 := &librarypb.ListBooksRequest{
			Parent:   "shelves/1",
			PageSize: 10,
		}
		pageToken1, err := ParsePageToken(request1)
		if err != nil {
			t.Errorf("\nOops 🔥\x1b[91m Failed asserting that\x1b[39m\n"+
				"✘got: %v\n\x1b[92m"+
				"want: %v\x1b[39m", err, nil)
		}

		request2 := &librarypb.ListBooksRequest{
			Parent:    "shelves/1",
			PageSize:  20,
			PageToken: pageToken1.Next(request1).String(),
		}
		pageToken2, err := ParsePageToken(request2)
		if err != nil {
			t.Errorf("\nOops 🔥\x1b[91m Failed asserting that\x1b[39m\n"+
				"✘got: %v\n\x1b[92m"+
				"want: %v\x1b[39m", err, nil)
		}
		if pageToken2.Offset != int64(10) {
			t.Errorf("\nOops 🔥\x1b[91m Failed asserting that\x1b[39m\n"+
				"✘got: %v\n\x1b[92m"+
				"want: %v\x1b[39m", pageToken2.Offset, 10)
		}

		request3 := &librarypb.ListBooksRequest{
			Parent:    "shelves/1",
			PageSize:  30,
			PageToken: pageToken2.Next(request2).String(),
		}
		pageToken3, err := ParsePageToken(request3)
		if err != nil {
			t.Errorf("\nOops 🔥\x1b[91m Failed asserting that\x1b[39m\n"+
				"✘got: %v\n\x1b[92m"+
				"want: %v\x1b[39m", err, nil)
		}

		if pageToken3.Offset != int64(30) {
			t.Errorf("\nOops 🔥\x1b[91m Failed asserting that\x1b[39m\n"+
				"✘got: %v\n\x1b[92m"+
				"want: %v\x1b[39m", pageToken3.Offset, 30)
		}
	})
}
