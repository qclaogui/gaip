// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package genai

import (
	"testing"
)

func TestInferFullModelName(t *testing.T) {
	for _, test := range []struct {
		name string
		want string
	}{
		{"xyz", "projects/proj/locations/loc/publishers/google/models/xyz"},
		{"models/abc", "projects/proj/locations/loc/publishers/google/models/abc"},
		{"publishers/foo/xyz", "projects/proj/locations/loc/publishers/foo/xyz"},
		{"x/y/z", "x/y/z"},
	} {
		t.Run(test.name, func(t *testing.T) {
			got := inferFullModelName("proj", "loc", test.name)
			if got != test.want {
				t.Errorf("got %q, want %q", got, test.want)
			}
		})
	}
}

func TestInferLocation(t *testing.T) {
	for _, test := range []struct {
		name                             string
		arg                              string
		cloudRegionEnv, cloudMlRegionEnv string
		want                             string
	}{
		{"arg passed", "us-west4", "abc", "def", "us-west4"},
		{"first env", "", "abc", "", "abc"},
		{"second env", "", "", "klm", "klm"},
		{"default", "", "", "", defaultLocation},
		{"first env precedence", "", "101", "klm", "101"},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("GOOGLE_CLOUD_REGION", test.cloudRegionEnv)
			t.Setenv("CLOUD_ML_REGION", test.cloudMlRegionEnv)
			got := inferLocation(test.arg)
			if got != test.want {
				t.Errorf("got %q, want %q", got, test.want)
			}
		})
	}
}
