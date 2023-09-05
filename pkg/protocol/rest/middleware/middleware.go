// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package middleware

import "net/http"

func WrapperHandler(next http.Handler) http.Handler {

	next = RequestID(next)
	next = Throttle(1000)(next)

	return next
}
