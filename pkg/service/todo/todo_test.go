// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package todo

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/qclaogui/gaip/genproto/todo/apiv1/todopb"
	"github.com/qclaogui/gaip/internal/repository/mysql"
	lg "github.com/qclaogui/gaip/tools/log"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var ID = "e75b6f03-e5fc-488c-8f75-ad1747be3d3a"

func serverSetupWithSQLDB(db *sql.DB) *ServiceServer {
	var cfg = Config{}
	repo, _ := mysql.NewTodoWithSQLDB(db)
	return &ServiceServer{
		Cfg:        cfg,
		logger:     lg.Logger,
		Registerer: prometheus.DefaultRegisterer,
		repo:       repo,
	}
}

func Test_toDoServiceServer_Create(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func() { _ = db.Close() }()

	ssv := serverSetupWithSQLDB(db)

	tm := time.Now().UTC().Add(time.Minute)
	reminder := timestamppb.New(tm)

	type args struct {
		ctx context.Context
		req *todopb.CreateRequest
	}
	cases := []struct {
		desc    string
		ssv     todopb.ToDoServiceServer
		args    args
		mock    func()
		want    *todopb.CreateResponse
		wantErr bool
	}{
		{
			desc: "OK",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &todopb.CreateRequest{
					Api: "v1",
					Item: &todopb.ToDo{
						Id:          ID,
						Title:       "title",
						Description: "description",
						CreatedAt:   reminder,
					},
				},
			},
			mock: func() {
				mock.ExpectExec("INSERT INTO ToDo").WithArgs(ID, "title", "description", tm).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: &todopb.CreateResponse{
				Api: "v1",
				Id:  ID,
			},
		},
		{
			desc: "Unsupported API",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &todopb.CreateRequest{
					Api: "v1000",
					Item: &todopb.ToDo{
						Title:       "title",
						Description: "description",
						CreatedAt:   reminder,
					},
				},
			},
			mock:    func() {},
			wantErr: true,
		},
		{
			desc: "Invalid Reminder field format",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &todopb.CreateRequest{
					Api: "v1",
					Item: &todopb.ToDo{
						Title:       "title",
						Description: "description",
						CreatedAt: &timestamp.Timestamp{
							Seconds: 1,
							Nanos:   -1,
						},
					},
				},
			},
			mock:    func() {},
			wantErr: true,
		},
		{
			desc: "INSERT failed",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &todopb.CreateRequest{
					Api: "v1",
					Item: &todopb.ToDo{
						Title:       "title",
						Description: "description",
						CreatedAt:   reminder,
					},
				},
			},
			mock: func() {
				mock.ExpectExec("INSERT INTO ToDo").WithArgs("ID", "title", "description", tm).
					WillReturnError(errors.New("INSERT failed"))
			},
			wantErr: true,
		},
		{
			desc: "LastInsertId failed",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &todopb.CreateRequest{
					Api: "v1",
					Item: &todopb.ToDo{
						Title:       "title",
						Description: "description",
						CreatedAt:   reminder,
					},
				},
			},
			mock: func() {
				mock.ExpectExec("INSERT INTO ToDo").WithArgs("title", "description", tm).
					WillReturnResult(sqlmock.NewErrorResult(errors.New("LastInsertId failed")))
			},
			wantErr: true,
		},
	}
	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			c.mock()
			got, err := c.ssv.Create(c.args.ctx, c.args.req)
			if (err != nil) != c.wantErr {
				t.Errorf("ToDoService.Create() error = %v, wantErr %v", err, c.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, c.want) {
				t.Errorf("ToDoService.Create() = %v, want %v", got, c.want)
			}
		})
	}
}

func Test_toDoServiceServer_Get(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%ssv' was not expected when opening a stub database connection", err)
	}
	defer func() { _ = db.Close() }()

	ssv := serverSetupWithSQLDB(db)

	tm := time.Now().UTC().Add(time.Minute)
	reminder := timestamppb.New(tm)

	type args struct {
		ctx context.Context
		req *todopb.GetRequest
	}

	cases := []struct {
		desc    string
		ssv     todopb.ToDoServiceServer
		args    args
		mock    func()
		want    *todopb.GetResponse
		wantErr bool
	}{
		{
			desc: "OK",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &todopb.GetRequest{
					Api: "v1",
					Id:  ID,
				},
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"ID", "Title", "Description", "Reminder"}).
					AddRow(ID, "title", "description", tm)
				mock.ExpectQuery("SELECT (.+) FROM ToDo").WithArgs(ID).WillReturnRows(rows)
			},
			want: &todopb.GetResponse{
				Api: "v1",
				Item: &todopb.ToDo{
					Id:          ID,
					Title:       "title",
					Description: "description",
					CreatedAt:   reminder,
				},
			},
		},
		{
			desc: "Unsupported API",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &todopb.GetRequest{
					Api: "v1",
					Id:  "1",
				},
			},
			mock:    func() {},
			wantErr: true,
		},
		{
			desc: "SELECT failed",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &todopb.GetRequest{
					Api: "v1",
					Id:  "1",
				},
			},
			mock: func() {
				mock.ExpectQuery("SELECT (.+) FROM ToDo").WithArgs(ID).
					WillReturnError(errors.New("SELECT failed"))
			},
			wantErr: true,
		},
		{
			desc: "Not found",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &todopb.GetRequest{
					Api: "v1",
					Id:  ID,
				},
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"ID", "Title", "Description", "Reminder"})
				mock.ExpectQuery("SELECT (.+) FROM ToDo").WithArgs(ID).WillReturnRows(rows)
			},
			wantErr: true,
		},
	}
	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			c.mock()
			got, err := c.ssv.Get(c.args.ctx, c.args.req)
			if (err != nil) != c.wantErr {
				t.Errorf("ToDoService.Get() error = %v, wantErr %v", err, c.wantErr)
				return
			}

			if err == nil && !reflect.DeepEqual(got, c.want) {
				t.Errorf("ToDoService.Get() = %v, want %v", got, c.want)
			}
		})
	}
}

func Test_toDoServiceServer_Update(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%ssv' was not expected when opening a stub database connection", err)
	}
	defer func() { _ = db.Close() }()

	ssv := serverSetupWithSQLDB(db)

	tm := time.Now().UTC().Add(time.Minute)
	reminder := timestamppb.New(tm)

	type args struct {
		ctx context.Context
		req *todopb.UpdateRequest
	}
	cases := []struct {
		desc    string
		ssv     todopb.ToDoServiceServer
		args    args
		mock    func()
		want    *todopb.UpdateResponse
		wantErr bool
	}{
		{
			desc: "OK",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &todopb.UpdateRequest{
					Api: "v1",
					Item: &todopb.ToDo{
						Id:          ID,
						Title:       "new title",
						Description: "new description",
						CreatedAt:   reminder,
					},
				},
			},
			mock: func() {
				mock.ExpectExec("UPDATE ToDo").WithArgs("new title", "new description", tm, ID).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: &todopb.UpdateResponse{
				Api:     "v1",
				Updated: 1,
			},
		},
		{
			desc: "Unsupported API",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &todopb.UpdateRequest{
					Api: "v1",
					Item: &todopb.ToDo{
						Id:          "1",
						Title:       "new title",
						Description: "new description",
						CreatedAt:   reminder,
					},
				},
			},
			mock:    func() {},
			wantErr: true,
		},
		{
			desc: "Invalid Reminder field format",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &todopb.UpdateRequest{
					Api: "v1",
					Item: &todopb.ToDo{
						Id:          "1",
						Title:       "new title",
						Description: "new description",
						CreatedAt: &timestamp.Timestamp{
							Seconds: 1,
							Nanos:   -1,
						},
					},
				},
			},
			mock:    func() {},
			wantErr: true,
		},
		{
			desc: "UPDATE failed",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &todopb.UpdateRequest{
					Api: "v1",
					Item: &todopb.ToDo{
						Id:          "1",
						Title:       "new title",
						Description: "new description",
						CreatedAt:   reminder,
					},
				},
			},
			mock: func() {
				mock.ExpectExec("UPDATE ToDo").WithArgs("new title", "new description", tm, 1).
					WillReturnError(errors.New("UPDATE failed"))
			},
			wantErr: true,
		},
		{
			desc: "RowsAffected failed",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &todopb.UpdateRequest{
					Api: "v1",
					Item: &todopb.ToDo{
						Id:          "1",
						Title:       "new title",
						Description: "new description",
						CreatedAt:   reminder,
					},
				},
			},
			mock: func() {
				mock.ExpectExec("UPDATE ToDo").WithArgs("new title", "new description", tm, 1).
					WillReturnResult(sqlmock.NewErrorResult(errors.New("RowsAffected failed")))
			},
			wantErr: true,
		},
		{
			desc: "Not Found",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &todopb.UpdateRequest{
					Api: "v1",
					Item: &todopb.ToDo{
						Id:          "1",
						Title:       "new title",
						Description: "new description",
						CreatedAt:   reminder,
					},
				},
			},
			mock: func() {
				mock.ExpectExec("UPDATE ToDo").WithArgs("new title", "new description", tm, 1).
					WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
	}
	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			c.mock()
			got, err := c.ssv.Update(c.args.ctx, c.args.req)
			if (err != nil) != c.wantErr {
				t.Errorf("ToDoService.Update() error = %v, wantErr %v", err, c.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, c.want) {
				t.Errorf("ToDoService.Update() = %v, want %v", got, c.want)
			}
		})
	}
}

func Test_toDoServiceServer_Delete(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%ssv' was not expected when opening a stub database connection", err)
	}
	defer func() { _ = db.Close() }()

	ssv := serverSetupWithSQLDB(db)

	type args struct {
		ctx context.Context
		req *todopb.DeleteRequest
	}
	cases := []struct {
		desc    string
		ssv     todopb.ToDoServiceServer
		args    args
		mock    func()
		want    *todopb.DeleteResponse
		wantErr bool
	}{
		{
			desc: "OK",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &todopb.DeleteRequest{
					Api: "v1",
					Id:  ID,
				},
			},
			mock: func() {
				mock.ExpectExec("DELETE FROM ToDo").WithArgs(ID).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: &todopb.DeleteResponse{
				Api:     "v1",
				Deleted: 1,
			},
		},
		{
			desc: "Unsupported API",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &todopb.DeleteRequest{
					Api: "v1",
					Id:  "1",
				},
			},
			mock:    func() {},
			wantErr: true,
		},
		{
			desc: "DELETE failed",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &todopb.DeleteRequest{
					Api: "v1",
					Id:  "1",
				},
			},
			mock: func() {
				mock.ExpectExec("DELETE FROM ToDo").WithArgs(1).
					WillReturnError(errors.New("DELETE failed"))
			},
			wantErr: true,
		},
		{
			desc: "RowsAffected failed",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &todopb.DeleteRequest{
					Api: "v1",
					Id:  "1",
				},
			},
			mock: func() {
				mock.ExpectExec("DELETE FROM ToDo").WithArgs(1).
					WillReturnResult(sqlmock.NewErrorResult(errors.New("RowsAffected failed")))
			},
			wantErr: true,
		},
		{
			desc: "Not Found",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &todopb.DeleteRequest{
					Api: "v1",
					Id:  "1",
				},
			},
			mock: func() {
				mock.ExpectExec("DELETE FROM ToDo").WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
	}
	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			c.mock()
			got, err := c.ssv.Delete(c.args.ctx, c.args.req)
			if (err != nil) != c.wantErr {
				t.Errorf("ToDoService.Delete() error = %v, wantErr %v", err, c.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, c.want) {
				t.Errorf("ToDoService.Delete() = %v, want %v", got, c.want)
			}
		})
	}
}

func Test_toDoServiceServer_List(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%ssv' was not expected when opening a stub database connection", err)
	}
	defer func() { _ = db.Close() }()

	ssv := serverSetupWithSQLDB(db)

	tm1 := time.Now().UTC().Add(time.Minute)
	reminder1 := timestamppb.New(tm1)

	tm2 := time.Now().UTC().Add(2 * time.Minute)
	reminder2 := timestamppb.New(tm2)

	var ID2 = "e75b6f83-e5fc-488c-8f75-ad1437be3d3a"

	type args struct {
		ctx context.Context
		req *todopb.ListRequest
	}
	cases := []struct {
		desc    string
		ssv     todopb.ToDoServiceServer
		args    args
		mock    func()
		want    *todopb.ListResponse
		wantErr bool
	}{
		{
			desc: "OK",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &todopb.ListRequest{
					Api: "v1",
				},
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"ID", "Title", "Description", "Reminder"}).
					AddRow(ID, "title 1", "description 1", tm1).
					AddRow(ID2, "title 2", "description 2", tm2)
				mock.ExpectQuery("SELECT (.+) FROM ToDo").WillReturnRows(rows)
			},
			want: &todopb.ListResponse{
				Api: "v1",
				Items: []*todopb.ToDo{
					{
						Id:          ID,
						Title:       "title 1",
						Description: "description 1",
						CreatedAt:   reminder1,
					},
					{
						Id:          ID2,
						Title:       "title 2",
						Description: "description 2",
						CreatedAt:   reminder2,
					},
				},
			},
		},
		{
			desc: "Empty",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &todopb.ListRequest{
					Api: "v1",
				},
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"ID", "Title", "Description", "Reminder"})
				mock.ExpectQuery("SELECT (.+) FROM ToDo").WillReturnRows(rows)
			},
			want: &todopb.ListResponse{
				Api:   "v1",
				Items: []*todopb.ToDo(nil),
			},
		},
		{
			desc: "Unsupported API",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &todopb.ListRequest{
					Api: "v1",
				},
			},
			mock:    func() {},
			wantErr: true,
		},
	}
	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			c.mock()
			got, err := c.ssv.List(c.args.ctx, c.args.req)
			if (err != nil) != c.wantErr {
				t.Errorf("ToDoService.List() error = %v, wantErr %v", err, c.wantErr)
				return
			}
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("ToDoService.List() = %v, want %v", got, c.want)
			}
		})
	}
}
