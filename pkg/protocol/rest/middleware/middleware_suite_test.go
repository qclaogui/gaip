// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package middleware_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMiddlewareSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Middleware REST")
}
