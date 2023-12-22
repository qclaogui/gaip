// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package mysql

import (
	"testing"

	"github.com/qclaogui/gaip/internal/ent"
	"github.com/qclaogui/gaip/internal/ent/enttest"
	"github.com/qclaogui/gaip/internal/ent/migrate"

	// go-sqlite3
	_ "github.com/mattn/go-sqlite3"
)

func Test_CreateShelf(t *testing.T) {

	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}

	// Create an ent.Client with in-memory SQLite database.
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer func() { _ = client.Close() }()

}
