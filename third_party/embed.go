// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package third_party

import (
	"embed"
)

//go:embed gen/openapiv2/*
var OpenAPI embed.FS
