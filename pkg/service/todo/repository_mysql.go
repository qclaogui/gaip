// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package todo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	pb "github.com/qclaogui/golang-api-server/api/gen/proto/todo/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// MysqlRepository fulfills the Repository interface
type MysqlRepository struct {
	db *sql.DB
}

// NewMysqlRepository is a factory function to generate a new repository
func NewMysqlRepository(db *sql.DB) (*MysqlRepository, error) {
	repo := &MysqlRepository{db: db}
	return repo, nil
}

func (m *MysqlRepository) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database-> " + err.Error())
	}
	return c, nil
}

func (m *MysqlRepository) Read(ctx context.Context, req *pb.ReadRequest) (*pb.ReadResponse, error) {
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

	todo := &pb.ToDo{}
	var reminder time.Time
	if err = rows.Scan(&todo.Id, &todo.Title, &todo.Description, &reminder); err != nil {
		return nil, fmt.Errorf("failed to retrieve field values from ToDo row-> " + err.Error())
	}

	todo.Reminder = timestamppb.New(reminder)
	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ToDo rows with ID='%s'", id))
	}

	return &pb.ReadResponse{Api: apiVersion, ToDo: todo}, nil
}

func (m *MysqlRepository) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	// get SQL connection from pool
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = c.Close() }()

	todo := req.GetToDo()
	_, err = c.ExecContext(ctx, "INSERT INTO ToDo(`ID`, `Title`, `Description`, `Reminder`) VALUES(?, ?, ?, ?)",
		todo.GetId(), todo.GetTitle(), todo.GetDescription(), todo.GetReminder().AsTime())
	if err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("failed to insert into ToDo-> "+err.Error()))
	}
	return &pb.CreateResponse{Api: apiVersion, Id: todo.GetId()}, nil
}

func (m *MysqlRepository) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	// get SQL connection from pool
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = c.Close() }()

	todo := req.GetToDo()

	res, err := c.ExecContext(ctx, "UPDATE ToDo SET `Title`=?, `Description`=?, `Reminder`=? WHERE `ID`=?",
		todo.Title, todo.Description, todo.Reminder.AsTime(), todo.Id)
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

	return &pb.UpdateResponse{Api: apiVersion, Updated: rows}, nil
}

func (m *MysqlRepository) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
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
	return &pb.DeleteResponse{Api: apiVersion, Deleted: rows}, nil
}

func (m *MysqlRepository) ReadAll(ctx context.Context, _ *pb.ReadAllRequest) (*pb.ReadAllResponse, error) {
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

	var todos []*pb.ToDo
	for rows.Next() {
		var todo = &pb.ToDo{}
		var reminder time.Time
		if err = rows.Scan(&todo.Id, &todo.Title, &todo.Description, &reminder); err != nil {
			return nil, status.Error(codes.Unknown, fmt.Sprintf("failed to retrieve field values from ToDo row-> "+err.Error()))
		}
		todo.Reminder = timestamppb.New(reminder)
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("failed to retrieve data from ToDo-> "+err.Error()))
	}

	return &pb.ReadAllResponse{Api: apiVersion, ToDos: todos}, nil
}
