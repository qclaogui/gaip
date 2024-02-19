// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package rest

import (
	"testing"
)

func TestKeysMatchPath(t *testing.T) {
	testCase := []struct {
		examine     map[string][]string
		lookFor     []string
		wantMatches map[string]bool
	}{
		{
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
	}

	for idx, tc := range testCase {
		matches := KeysMatchPath(tc.examine, tc.lookFor)
		if got, want := len(matches), len(tc.wantMatches); got != want {
			t.Errorf("tc = %d: unexpected number of variables returned: got %v, want %v: returned elements: %v",
				idx, got, want, matches)
			continue
		}
		for matchIdx, got := range matches {
			if !tc.wantMatches[got] {
				t.Errorf("testCase = %d: got unexpected match #%d %q; expected matches: %v", idx, matchIdx, got, matches)
			}
		}

	}
}
