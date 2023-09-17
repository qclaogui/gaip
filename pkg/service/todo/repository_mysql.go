// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package todo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/qclaogui/golang-api-server/genproto/todo/apiv1/todopb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// MysqlRepo fulfills the Repository interface
type MysqlRepo struct {
	db *sql.DB
}

// NewMysqlRepo is a factory function to generate a new repository
func NewMysqlRepo(db *sql.DB) (*MysqlRepo, error) {
	repo := &MysqlRepo{db: db}
	return repo, nil
}

func (m *MysqlRepo) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database-> " + err.Error())
	}
	return c, nil
}

func (m *MysqlRepo) Get(ctx context.Context, req *todopb.GetRequest) (*todopb.GetResponse, error) {
	// get SQL connection from pool
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = c.Close() }()

	id := req.GetId()
	rows, err := c.QueryContext(ctx, "SELECT `ID`, `Title`, `Description`, `Reminder` FROM ToDo WHERE `ID`=?", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ToDo-> "+err.Error())
	}
	defer func() { _ = rows.Close() }()

	if !rows.Next() {
		if err = rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to select from ToDo-> "+err.Error())
		}
	}

	todo := &todopb.ToDo{}
	var reminder time.Time
	if err = rows.Scan(&todo.Id, &todo.Title, &todo.Description, &reminder); err != nil {
		return nil, fmt.Errorf("failed to retrieve field values from ToDo row-> " + err.Error())
	}

	todo.CreatedAt = timestamppb.New(reminder)
	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ToDo rows with ID='%s'", id))
	}

	return &todopb.GetResponse{Api: apiVersion, Item: todo}, nil
}

func (m *MysqlRepo) Create(ctx context.Context, req *todopb.CreateRequest) (*todopb.CreateResponse, error) {
	// get SQL connection from pool
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = c.Close() }()

	todo := req.GetItem()
	_, err = c.ExecContext(ctx, "INSERT INTO ToDo(`ID`, `Title`, `Description`, `Reminder`) VALUES(?, ?, ?, ?)",
		todo.GetId(), todo.GetTitle(), todo.GetDescription(), todo.GetCreatedAt().AsTime())
	if err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("failed to insert into ToDo-> "+err.Error()))
	}
	return &todopb.CreateResponse{Api: apiVersion, Id: todo.GetId()}, nil
}

func (m *MysqlRepo) Update(ctx context.Context, req *todopb.UpdateRequest) (*todopb.UpdateResponse, error) {
	// get SQL connection from pool
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = c.Close() }()

	todo := req.GetItem()

	res, err := c.ExecContext(ctx, "UPDATE ToDo SET `Title`=?, `Description`=?, `Reminder`=? WHERE `ID`=?",
		todo.Title, todo.Description, todo.GetCreatedAt().AsTime(), todo.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("failed to update ToDo-> "+err.Error()))
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("failed to retrieve rows affected value-> "+err.Error()))
	}
	if rows == 0 {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("ToDo with ID='%s' is not found", todo.Id))
	}

	return &todopb.UpdateResponse{Api: apiVersion, Updated: rows}, nil
}

func (m *MysqlRepo) Delete(ctx context.Context, req *todopb.DeleteRequest) (*todopb.DeleteResponse, error) {
	// get SQL connection from pool
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = c.Close() }()

	id := req.GetId()
	res, err := c.ExecContext(ctx, "DELETE FROM ToDo WHERE `ID`=?", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("failed to delete ToDo-> "+err.Error()))
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("failed to retrieve rows affected value-> "+err.Error()))
	}
	if rows == 0 {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("ToDo with ID='%s' is not found", id))
	}
	return &todopb.DeleteResponse{Api: apiVersion, Deleted: rows}, nil
}

func (m *MysqlRepo) List(ctx context.Context, _ *todopb.ListRequest) (*todopb.ListResponse, error) {
	// get SQL connection from pool
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = c.Close() }()

	rows, err := c.QueryContext(ctx, "SELECT `ID`, `Title`, `Description`, `Reminder` FROM ToDo")
	if err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("failed to select from ToDo-> "+err.Error()))
	}
	defer func() { _ = rows.Close() }()

	var todos []*todopb.ToDo
	for rows.Next() {
		var todo = &todopb.ToDo{}
		var reminder time.Time
		if err = rows.Scan(&todo.Id, &todo.Title, &todo.Description, &reminder); err != nil {
			return nil, status.Error(codes.Unknown, fmt.Sprintf("failed to retrieve field values from ToDo row-> "+err.Error()))
		}
		todo.CreatedAt = timestamppb.New(reminder)
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("failed to retrieve data from ToDo-> "+err.Error()))
	}

	return &todopb.ListResponse{Api: apiVersion, Items: todos}, nil
}
