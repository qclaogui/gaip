// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package rest

import (
	"testing"
)

func TestKeysMatchPath(t *testing.T) {
	for _, tc := range []struct {
		name        string
		examine     map[string][]string
		lookFor     []string
		wantMatches map[string]bool
	}{
		{
			name: "loc",
			examine: map[string][]string{
				"loc":           nil,
				"location":      nil,
				"loc.lat":       nil,
				"extra.loc.lat": nil,
				"location.lat":  nil,
			},
			lookFor:     []string{"loc"},
			wantMatches: map[string]bool{"loc": true, "loc.lat": true},
		},
		{
			name: "location",
			examine: map[string][]string{
				"loc":           nil,
				"location":      nil,
				"loc.lat":       nil,
				"extra.loc.lat": nil,
				"location.lat":  nil,
			},
			lookFor:     []string{"location", "loc"},
			wantMatches: map[string]bool{"loc": true, "location": true, "loc.lat": true, "location.lat": true},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			matches := KeysMatchPath(tc.examine, tc.lookFor)
			if got, want := len(matches), len(tc.wantMatches); got != want {
				t.Errorf("unexpected number of variables returned: got %v, want %v: returned elements: %v",
					got, want, matches)
			}
			for matchIdx, got := range matches {
				if !tc.wantMatches[got] {
					t.Errorf("got unexpected match #%d %q; expected matches: %v", matchIdx, got, matches)
				}
			}
		})
	}
}
