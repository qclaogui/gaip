// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package rest

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestProcessQueryString(t *testing.T) {
	for idx, tc := range []struct {
		queryString string
		wantInt     bool
		wantParams  map[string][]string
		wantError   bool
	}{
		{queryString: ""},
		{
			queryString: "foo=bar",
			wantParams: map[string][]string{
				"foo": {"bar"},
			},
		},
		{
			queryString: "$foo=bar",
			wantParams: map[string][]string{
				"$foo": {"bar"},
			},
		},
		{
			queryString: "%24foo=bar",
			wantParams: map[string][]string{
				"$foo": {"bar"},
			},
		},
		{
			queryString: "foo%3Dbar",
			wantParams: map[string][]string{
				"foo=bar": {""},
			},
		},
		{
			queryString: "foo%3Dbar=xyz",
			wantParams: map[string][]string{
				"foo=bar": {"xyz"},
			},
		},
		{
			queryString: "%24foo%3Dbar",
			wantParams: map[string][]string{
				"$foo=bar": {""},
			},
		},
		{
			queryString: "%24foo%3Dbar=xyz",
			wantParams: map[string][]string{
				"$foo=bar": {"xyz"},
			},
		},

		// system param incorrect
		{
			queryString: "%24alt%3Djson",
			wantParams: map[string][]string{
				"$alt=json": {""},
			},
		},
		{
			queryString: "$ALT=JSON",
			wantParams: map[string][]string{
				"$ALT": {"JSON"},
			},
		},
		{
			queryString: "%24ALT=JSON",
			wantParams: map[string][]string{
				"$ALT": {"JSON"},
			},
		},

		// system param by itself
		{queryString: "alt=json"},
		{queryString: "$alt=json"},
		{queryString: "%24alt=json"},
		{
			queryString: "alt=json%3Benum-encoding=int",
			wantInt:     true,
		},
		{
			queryString: "$alt=json%3Benum-encoding=int",
			wantInt:     true,
		},
		{
			queryString: "%24alt=json%3Benum-encoding=int",
			wantInt:     true,
		},
		{
			queryString: "alt=json%3Benum-encoding%3Dint",
			wantInt:     true,
		},
		{
			queryString: "$alt=json%3Benum-encoding%3Dint",
			wantInt:     true,
		},
		{
			queryString: "%24alt=json%3Benum-encoding%3Dint",
			wantInt:     true,
		},

		// system param+query params in front
		{
			queryString: "foo=bar&alt=json",
			wantParams: map[string][]string{
				"foo": {"bar"},
			},
		},
		{
			queryString: "foo=bar&$alt=json",
			wantParams: map[string][]string{
				"foo": {"bar"},
			},
		},
		{
			queryString: "foo=bar&%24alt=json",
			wantParams: map[string][]string{
				"foo": {"bar"},
			},
		},
		{
			queryString: "foo=bar&alt=json%3Benum-encoding=int",
			wantInt:     true,
			wantParams: map[string][]string{
				"foo": {"bar"},
			},
		},
		{
			queryString: "foo=bar&$alt=json%3Benum-encoding=int",
			wantInt:     true,
			wantParams: map[string][]string{
				"foo": {"bar"},
			},
		},
		{
			queryString: "foo=bar&%24alt=json%3Benum-encoding=int",
			wantInt:     true,
			wantParams: map[string][]string{
				"foo": {"bar"},
			},
		},

		// system param+query params in rear
		{
			queryString: "alt=json&foo=bar",
			wantParams: map[string][]string{
				"foo": {"bar"},
			},
		},
		{
			queryString: "$alt=json&foo=bar",
			wantParams: map[string][]string{
				"foo": {"bar"},
			},
		},
		{
			queryString: "%24alt=json&foo=bar",
			wantParams: map[string][]string{
				"foo": {"bar"},
			},
		},
		{
			queryString: "alt=json%3Benum-encoding=int&foo=bar",
			wantInt:     true,
			wantParams: map[string][]string{
				"foo": {"bar"},
			},
		},
		{
			queryString: "$alt=json%3Benum-encoding=int&foo=bar",
			wantInt:     true,
			wantParams: map[string][]string{
				"foo": {"bar"},
			},
		},
		{
			queryString: "%24alt=json%3Benum-encoding=int&foo=bar",
			wantInt:     true,
			wantParams: map[string][]string{
				"foo": {"bar"},
			},
		},

		// system param errors
		{
			queryString: "$alt=foo",
			wantError:   true,
		},
		{
			queryString: "$alt",
			wantError:   true,
		},
		{
			queryString: "$alt=JSON",
			wantError:   true,
		},
		{
			queryString: "$alt=json%3Benum-encoding=string",
			wantError:   true,
		},
		{
			queryString: "$alt=json;enum-encoding=int", // unencoded semicolon
			wantError:   true,
		},
		{
			queryString: "$alt=json%3Benum-encoding=INT",
			wantError:   true,
		},
		{
			queryString: "foo&$alt=json&bar&alt=json", // repeated
			wantError:   true,
		},
	} {
		label := fmt.Sprintf("[%2d %q]", idx, tc.queryString)

		systemParams, queryParams, err := processQueryString(tc.queryString)

		if got, want := err != nil, tc.wantError; got != want {
			t.Errorf("%s: error condition not met: want error: %v, got error:%v", label, tc.wantError, err)
			continue
		}
		if err != nil {
			continue
		}

		wantParams := tc.wantParams
		if wantParams == nil {
			wantParams = map[string][]string{}
		}
		if got, want := queryParams, wantParams; !cmp.Equal(got, want) {
			t.Errorf("%s: query params: want %#v, got %#v", label, want, got)
		}

		if got, want := systemParams.EnumEncodingAsInt, tc.wantInt; got != want {
			t.Errorf("%s: numeric enums: want %v, got %v", label, want, got)
		}
	}
}
