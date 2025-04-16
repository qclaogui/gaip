// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package thirdparty

import (
	"embed"
)

//go:embed gen/openapiv2/*
var OpenAPI embed.FS
