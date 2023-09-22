// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package service

import "sync/atomic"

// UniqID provides a numerical id that is guaranteed to be unique.
type UniqID struct {
	i int64
}

// Next gets the next unique id.
func (u *UniqID) Next() int64 {
	return atomic.AddInt64(&u.i, 1) - 1
}
