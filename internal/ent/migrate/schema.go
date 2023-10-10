// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TodosColumns holds the columns for the "todos" table.
	TodosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// TodosTable holds the schema information for the "todos" table.
	TodosTable = &schema.Table{
		Name:       "todos",
		Columns:    TodosColumns,
		PrimaryKey: []*schema.Column{TodosColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeUUID, Comment: "用户唯一id"},
		{Name: "user_name", Type: field.TypeString, Comment: "用户名"},
		{Name: "last_active", Type: field.TypeTime, Comment: "最后活跃时间戳", Default: "CURRENT_TIMESTAMP"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Comment:    "User is the model entity for the User schema.",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "idx_user_id",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[1]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TodosTable,
		UsersTable,
	}
)

func init() {
	UsersTable.Annotation = &entsql.Annotation{}
}