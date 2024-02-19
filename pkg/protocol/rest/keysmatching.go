// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package rest

import "strings"

// KeysMatchPath returns the keys in `examine` that match any of the elements in `lookFor`,
// interpreting the latter as dotted-field paths. This means a match occurs when (a) the two are
// identical, or (b) when any element of `lookFor` is a prefix of the `examine` key and is followed
// by a period. For example:
// KeysMatchPath(map[string][]string{"loc": nil, "loc.lat": nil, "location":nil},
//
//	         []string{"loc"})
//	== []string{"loc","loc.lat"}
func KeysMatchPath(examine map[string][]string, lookFor []string) []string {
	var matchingKeys []string
	for key := range examine {
		for _, target := range lookFor {
			if matchesSelfOrParent(key, target) {
				matchingKeys = append(matchingKeys, key)
				break
			}
		}
	}
	return matchingKeys
}

// matchesSelfOrParent returns whether any element of `lookFor` is or contains a full or partial
// path (in the dotted-field sense) to `examine`. In other words, this returns true when (a) the two
// are identical, or (b) when `examine` starts with `lookFor` and is followed by a period.
func matchesSelfOrParent(examine, lookFor string) bool {
	if !strings.HasPrefix(examine, lookFor) {
		return false
	}
	return len(examine) == len(lookFor) || examine[len(lookFor)] == '.'
}
