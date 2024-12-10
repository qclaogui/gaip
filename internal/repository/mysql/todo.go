// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"github.com/qclaogui/gaip/genproto/todo/apiv1/todopb"
	"github.com/qclaogui/gaip/internal/ent"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// NewTodo is a factory function to generate a new repository
func NewTodo(cfg Config) (todopb.ToDoServiceServer, error) {
	// add MySQL driver specific parameter to parse date/time
	// Drop it for another database
	param := "parseTime=true"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", cfg.User, cfg.Password, cfg.Host, cfg.Schema, param)

	client, err := ent.Open(dialect.MySQL, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to mysql: %v", err)
	}
	repo := &todoImpl{entClient: client}
	return repo, nil
}

// todoImpl fulfills the todoImpl interface
// All data are managed by MysqlCfg.
//
// todoImpl is used to implement ToDoServiceServer.
type todoImpl struct {
	todopb.UnimplementedToDoServiceServer

	sqlDB     *sql.DB
	entClient *ent.Client
}

// NewTodoWithSQLDB is a factory function to generate a new repository
func NewTodoWithSQLDB(db *sql.DB) (todopb.ToDoServiceServer, error) {
	repo := &todoImpl{sqlDB: db}
	return repo, nil
}

func (m *todoImpl) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.sqlDB.Conn(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database-> %w", err)
	}
	return c, nil
}

func (m *todoImpl) GetTodo(ctx context.Context, req *todopb.GetTodoRequest) (*todopb.GetTodoResponse, error) {
	// get SQL connection from pool
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = c.Close() }()

	id := req.GetId()
	rows, err := c.QueryContext(ctx, "SELECT `ID`, `Title`, `Description`, `Reminder` FROM ToDo WHERE `ID`=?", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("failed to select from ToDo-> %s", err.Error()))
	}
	defer func() { _ = rows.Close() }()

	if !rows.Next() {
		if err = rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, fmt.Sprintf("failed to select from ToDo-> %s", err.Error()))
		}
	}

	todo := &todopb.ToDo{}
	var reminder time.Time
	if err = rows.Scan(&todo.Id, &todo.Title, &todo.Description, &reminder); err != nil {
		return nil, fmt.Errorf("failed to retrieve field values from ToDo row-> %w", err)
	}

	todo.CreateTime = timestamppb.New(reminder)
	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ToDo rows with ID='%s'", id))
	}

	return &todopb.GetTodoResponse{Api: "v1", Item: todo}, nil
}

func (m *todoImpl) CreateTodo(ctx context.Context, req *todopb.CreateTodoRequest) (*todopb.CreateTodoResponse, error) {
	// get SQL connection from pool
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = c.Close() }()

	todo := req.GetItem()
	_, err = c.ExecContext(ctx, "INSERT INTO ToDo(`ID`, `Title`, `Description`, `Reminder`) VALUES(?, ?, ?, ?)",
		todo.GetId(), todo.GetTitle(), todo.GetDescription(), todo.GetCreateTime().AsTime())
	if err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("failed to insert into ToDo-> %s", err.Error()))
	}
	return &todopb.CreateTodoResponse{Api: "v1", Id: todo.GetId()}, nil
}

func (m *todoImpl) UpdateTodo(ctx context.Context, req *todopb.UpdateTodoRequest) (*todopb.UpdateTodoResponse, error) {
	// get SQL connection from pool
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = c.Close() }()

	todo := req.GetItem()

	res, err := c.ExecContext(ctx, "UPDATE ToDo SET `Title`=?, `Description`=?, `Reminder`=? WHERE `ID`=?",
		todo.Title, todo.Description, todo.GetCreateTime().AsTime(), todo.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("failed to update ToDo-> %s", err.Error()))
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("failed to retrieve rows affected value-> %s", err.Error()))
	}
	if rows == 0 {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("ToDo with ID='%s' is not found", todo.Id))
	}

	return &todopb.UpdateTodoResponse{Api: "v1", Updated: rows}, nil
}

func (m *todoImpl) DeleteTodo(ctx context.Context, req *todopb.DeleteTodoRequest) (*todopb.DeleteTodoResponse, error) {
	// get SQL connection from pool
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = c.Close() }()

	id := req.GetId()
	res, err := c.ExecContext(ctx, "DELETE FROM ToDo WHERE `ID`=?", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("failed to delete ToDo-> %s", err.Error()))
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("failed to retrieve rows affected value-> %s", err.Error()))
	}
	if rows == 0 {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("ToDo with ID='%s' is not found", id))
	}
	return &todopb.DeleteTodoResponse{Api: "v1", Deleted: rows}, nil
}

func (m *todoImpl) ListTodo(ctx context.Context, _ *todopb.ListTodoRequest) (*todopb.ListTodoResponse, error) {
	// get SQL connection from pool
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = c.Close() }()

	rows, err := c.QueryContext(ctx, "SELECT `ID`, `Title`, `Description`, `Reminder` FROM ToDo")
	if err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("failed to select from ToDo-> %s", err.Error()))
	}
	defer func() { _ = rows.Close() }()

	var todos []*todopb.ToDo
	for rows.Next() {
		todo := &todopb.ToDo{}
		var reminder time.Time
		if err = rows.Scan(&todo.Id, &todo.Title, &todo.Description, &reminder); err != nil {
			return nil, status.Error(codes.Unknown, fmt.Sprintf("failed to retrieve field values from ToDo row-> %s", err.Error()))
		}
		todo.CreateTime = timestamppb.New(reminder)
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("failed to retrieve data from ToDo-> %s", err.Error()))
	}

	return &todopb.ListTodoResponse{Api: "v1", Items: todos}, nil
}
