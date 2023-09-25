// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}

func run() error {
	if err := entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureUpsert,
			gen.FeatureVersionedMigration,
			gen.FeatureIntercept,
			gen.FeatureNamedEdges,
		},
	}); err != nil {
		return errors.Wrap(err, "ent codegen")
	}

	return nil
}
