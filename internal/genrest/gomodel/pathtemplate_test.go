// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package gomodel

import (
	"testing"
)

func TestHasVariables(t *testing.T) {
	for idx, testCase := range []struct {
		stringTemplate   string
		expectVars       bool
		expectNestedVars bool
	}{
		{
			stringTemplate:   "/aa/cc/ee/*/gg/ii/jj/*/kk/**:ll",
			expectVars:       false,
			expectNestedVars: false,
		},
		{
			stringTemplate:   "/aa/{bb}/cc/{dd=ee/*/gg}/{hh=ii/jj/*/kk/**}:ll",
			expectVars:       true,
			expectNestedVars: false,
		},
		{
			stringTemplate:   "/aa/{bb}/cc/{dd=ee/*/gg/{hh=ii/jj/*/kk}/**}:ll",
			expectVars:       true,
			expectNestedVars: true,
		},
	} {
		parsed, err := ParseTemplate(testCase.stringTemplate)
		if err != nil {
			t.Errorf("testCase = %d: ParseTemplate failed: %s \n   Test case input: %v", idx, err, testCase)
		}

		hasVars, hasNestedVars := parsed.HasVariables()
		if got, want := hasVars, testCase.expectVars; got != want {
			t.Errorf("testCase = %d: HasVars() failed checking variables: got %v, want %v", idx, got, want)
		}
		if got, want := hasNestedVars, testCase.expectNestedVars; got != want {
			t.Errorf("testCase = %d: HasVars() failed checking nested variables: got %v, want %v", idx, got, want)
		}
	}
}

func TestListVariables(t *testing.T) {
	for idx, testCase := range []struct {
		stringTemplate string
		expectVars     []string
	}{
		{
			stringTemplate: "/aa/cc/ee/*/gg/ii/jj/*/kk/**:ll",
			expectVars:     nil,
		},
		{
			stringTemplate: "/aa/{bb}/cc/{dd=ee/*/gg}/{hh=ii/jj/*/kk/**}:ll",
			expectVars:     []string{"bb", "dd", "hh"},
		},
		{
			stringTemplate: "/aa/{bb}/cc/{dd=ee/*/gg/{hh=ii/jj/*/kk}/**}:ll",
			expectVars:     []string{"bb", "dd", "hh"},
		},
	} {
		parsed, err := ParseTemplate(testCase.stringTemplate)
		if err != nil {
			t.Errorf("testCase = %d: ParseTemplate failed: %s \n   Test case input: %v", idx, err, testCase)
		}

		varList := parsed.ListVariables()
		if got, want := len(varList), len(testCase.expectVars); got != want {
			t.Errorf("testCase = %d: ListVars() unexpected number of variables returned: got %v, want %v: returned elements: %v",
				idx, got, want, varList)
			continue
		}
		for varIdx, got := range varList {
			if want := testCase.expectVars[varIdx]; got != want {
				t.Errorf("testCase = %d: ListVars() variable %d unexpected: got %v, want %v", idx, varIdx, got, want)
			}
		}
	}
}
