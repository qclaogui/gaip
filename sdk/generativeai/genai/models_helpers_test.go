// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package genai

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestContentHelpers(t *testing.T) {
	t.Run("Text", func(t *testing.T) {
		want := []*Content{
			NewContentFromText("test", RoleUser),
		}
		got := Text("test")

		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("Text mismatch (-want +got):\n%s", diff)
		}
	})
}
