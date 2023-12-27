// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package project

import (
	"testing"

	"entgo.io/ent/dialect"

	// go-sqlite3
	_ "github.com/mattn/go-sqlite3"
	"github.com/qclaogui/gaip/internal/ent/enttest"
)

func TestService(t *testing.T) {
	// Create an ent.Client with in-memory SQLite database.
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	defer func() { _ = client.Close() }()

	//svc, _ := NewIdentityService()
	//ctx := context.Background()
	//_, err := svc.CreateUser(ctx, &pb.CreateUserRequest{
	//	User: &pb.User{
	//		Email:       "test@example.com",
	//		DisplayName: "qc",
	//	},
	//})
	//if err != nil {
	//	t.Errorf("\nOops 🔥\x1b[91m Failed asserting that\x1b[39m\n"+
	//		"✘got: %v\n\x1b[92m"+
	//		"want: %v\x1b[39m", err, nil)
	//}

	//want := 1
	//if got := client.User.Query().CountX(ctx); got != want {
	//	t.Errorf("\nOops 🔥\x1b[91m Failed asserting that\x1b[39m\n"+
	//		"✘got: %v\n\x1b[92m"+
	//		"want: %v\x1b[39m", got, want)
	//}
}
