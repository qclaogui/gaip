// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package service

import "testing"

func TestUniqID_Next(t *testing.T) {
	var want int64 = 2

	u := &UniqID{want}
	if got := u.Next(); got != want {
		t.Errorf("\nOops 🔥\x1b[91m Failed asserting that\x1b[39m\n"+
			"✘got: %v\n\x1b[92m"+
			"want: %v\x1b[39m", got, want)
	}
}
