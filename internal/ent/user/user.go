// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldUserName holds the string denoting the user_name field in the database.
	FieldUserName = "user_name"
	// FieldLastActive holds the string denoting the last_active field in the database.
	FieldLastActive = "last_active"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldUserName,
	FieldLastActive,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID func() uuid.UUID
	// UserNameValidator is a validator for the "user_name" field. It is called by the builders before save.
	UserNameValidator func(string) error
	// DefaultLastActive holds the default value on creation for the "last_active" field.
	DefaultLastActive time.Time
	// UpdateDefaultLastActive holds the default value on update for the "last_active" field.
	UpdateDefaultLastActive func() time.Time
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByUserName orders the results by the user_name field.
func ByUserName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserName, opts...).ToFunc()
}

// ByLastActive orders the results by the last_active field.
func ByLastActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastActive, opts...).ToFunc()
}
