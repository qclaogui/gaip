package v1

import (
	"context"
	"errors"
	"github.com/qclaogui/golang-api-server/pkg/service/todo"
	"reflect"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	pb "github.com/qclaogui/golang-api-server/pkg/api/todopb/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var ID = "e75b6f03-e5fc-488c-8f75-ad1747be3d3a"

func Test_toDoServiceServer_Create(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func() { _ = db.Close() }()
	//ssv, _ := NewToDoService(WithMemoryToDoRepository())
	repo, _ := todo.NewMysqlRepository(db)
	ssv, _ := NewToDoService(WithToDoRepository(repo))

	tm := time.Now().UTC().Add(time.Minute)
	reminder := timestamppb.New(tm)

	type args struct {
		ctx context.Context
		req *pb.CreateRequest
	}
	tests := []struct {
		name    string
		ssv     pb.ToDoServiceServer
		args    args
		mock    func()
		want    *pb.CreateResponse
		wantErr bool
	}{
		{
			name: "OK",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.CreateRequest{
					Api: "v1",
					ToDo: &pb.ToDo{
						Id:          ID,
						Title:       "title",
						Description: "description",
						Reminder:    reminder,
					},
				},
			},
			mock: func() {
				mock.ExpectExec("INSERT INTO ToDo").WithArgs(ID, "title", "description", tm).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: &pb.CreateResponse{
				Api: "v1",
				Id:  ID,
			},
		},
		{
			name: "Unsupported API",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.CreateRequest{
					Api: "v1000",
					ToDo: &pb.ToDo{
						Title:       "title",
						Description: "description",
						Reminder:    reminder,
					},
				},
			},
			mock:    func() {},
			wantErr: true,
		},
		{
			name: "Invalid Reminder field format",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.CreateRequest{
					Api: "v1",
					ToDo: &pb.ToDo{
						Title:       "title",
						Description: "description",
						Reminder: &timestamp.Timestamp{
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
			name: "INSERT failed",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.CreateRequest{
					Api: "v1",
					ToDo: &pb.ToDo{
						Title:       "title",
						Description: "description",
						Reminder:    reminder,
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
			name: "LastInsertId failed",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.CreateRequest{
					Api: "v1",
					ToDo: &pb.ToDo{
						Title:       "title",
						Description: "description",
						Reminder:    reminder,
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.ssv.Create(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToDoService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToDoService.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toDoServiceServer_Read(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%ssv' was not expected when opening a stub database connection", err)
	}
	defer func() { _ = db.Close() }()

	repo, _ := todo.NewMysqlRepository(db)
	ssv, _ := NewToDoService(WithToDoRepository(repo))

	tm := time.Now().UTC().Add(time.Minute)
	reminder := timestamppb.New(tm)

	type args struct {
		ctx context.Context
		req *pb.ReadRequest
	}
	tests := []struct {
		name    string
		ssv     pb.ToDoServiceServer
		args    args
		mock    func()
		want    *pb.ReadResponse
		wantErr bool
	}{
		{
			name: "OK",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.ReadRequest{
					Api: "v1",
					Id:  ID,
				},
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"ID", "Title", "Description", "Reminder"}).
					AddRow(ID, "title", "description", tm)
				mock.ExpectQuery("SELECT (.+) FROM ToDo").WithArgs(ID).WillReturnRows(rows)
			},
			want: &pb.ReadResponse{
				Api: "v1",
				ToDo: &pb.ToDo{
					Id:          ID,
					Title:       "title",
					Description: "description",
					Reminder:    reminder,
				},
			},
		},
		{
			name: "Unsupported API",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.ReadRequest{
					Api: "v1",
					Id:  "1",
				},
			},
			mock:    func() {},
			wantErr: true,
		},
		{
			name: "SELECT failed",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.ReadRequest{
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
			name: "Not found",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.ReadRequest{
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.ssv.Read(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToDoService.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToDoService.Read() = %v, want %v", got, tt.want)
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

	repo, _ := todo.NewMysqlRepository(db)
	ssv, _ := NewToDoService(WithToDoRepository(repo))

	tm := time.Now().UTC().Add(time.Minute)
	reminder := timestamppb.New(tm)

	type args struct {
		ctx context.Context
		req *pb.UpdateRequest
	}
	tests := []struct {
		name    string
		ssv     pb.ToDoServiceServer
		args    args
		mock    func()
		want    *pb.UpdateResponse
		wantErr bool
	}{
		{
			name: "OK",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.UpdateRequest{
					Api: "v1",
					ToDo: &pb.ToDo{
						Id:          ID,
						Title:       "new title",
						Description: "new description",
						Reminder:    reminder,
					},
				},
			},
			mock: func() {
				mock.ExpectExec("UPDATE ToDo").WithArgs("new title", "new description", tm, ID).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: &pb.UpdateResponse{
				Api:     "v1",
				Updated: 1,
			},
		},
		{
			name: "Unsupported API",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.UpdateRequest{
					Api: "v1",
					ToDo: &pb.ToDo{
						Id:          "1",
						Title:       "new title",
						Description: "new description",
						Reminder:    reminder,
					},
				},
			},
			mock:    func() {},
			wantErr: true,
		},
		{
			name: "Invalid Reminder field format",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.UpdateRequest{
					Api: "v1",
					ToDo: &pb.ToDo{
						Id:          "1",
						Title:       "new title",
						Description: "new description",
						Reminder: &timestamp.Timestamp{
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
			name: "UPDATE failed",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.UpdateRequest{
					Api: "v1",
					ToDo: &pb.ToDo{
						Id:          "1",
						Title:       "new title",
						Description: "new description",
						Reminder:    reminder,
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
			name: "RowsAffected failed",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.UpdateRequest{
					Api: "v1",
					ToDo: &pb.ToDo{
						Id:          "1",
						Title:       "new title",
						Description: "new description",
						Reminder:    reminder,
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
			name: "Not Found",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.UpdateRequest{
					Api: "v1",
					ToDo: &pb.ToDo{
						Id:          "1",
						Title:       "new title",
						Description: "new description",
						Reminder:    reminder,
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.ssv.Update(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToDoService.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToDoService.Update() = %v, want %v", got, tt.want)
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

	repo, _ := todo.NewMysqlRepository(db)
	ssv, _ := NewToDoService(WithToDoRepository(repo))

	type args struct {
		ctx context.Context
		req *pb.DeleteRequest
	}
	tests := []struct {
		name    string
		ssv     pb.ToDoServiceServer
		args    args
		mock    func()
		want    *pb.DeleteResponse
		wantErr bool
	}{
		{
			name: "OK",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.DeleteRequest{
					Api: "v1",
					Id:  ID,
				},
			},
			mock: func() {
				mock.ExpectExec("DELETE FROM ToDo").WithArgs(ID).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: &pb.DeleteResponse{
				Api:     "v1",
				Deleted: 1,
			},
		},
		{
			name: "Unsupported API",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.DeleteRequest{
					Api: "v1",
					Id:  "1",
				},
			},
			mock:    func() {},
			wantErr: true,
		},
		{
			name: "DELETE failed",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.DeleteRequest{
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
			name: "RowsAffected failed",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.DeleteRequest{
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
			name: "Not Found",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.DeleteRequest{
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.ssv.Delete(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToDoService.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToDoService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toDoServiceServer_ReadAll(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%ssv' was not expected when opening a stub database connection", err)
	}
	defer func() { _ = db.Close() }()

	repo, _ := todo.NewMysqlRepository(db)
	ssv, _ := NewToDoService(WithToDoRepository(repo))

	tm1 := time.Now().UTC().Add(time.Minute)
	reminder1 := timestamppb.New(tm1)

	tm2 := time.Now().UTC().Add(2 * time.Minute)
	reminder2 := timestamppb.New(tm2)

	var ID2 = "e75b6f83-e5fc-488c-8f75-ad1437be3d3a"

	type args struct {
		ctx context.Context
		req *pb.ReadAllRequest
	}
	tests := []struct {
		name    string
		ssv     pb.ToDoServiceServer
		args    args
		mock    func()
		want    *pb.ReadAllResponse
		wantErr bool
	}{
		{
			name: "OK",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.ReadAllRequest{
					Api: "v1",
				},
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"ID", "Title", "Description", "Reminder"}).
					AddRow(ID, "title 1", "description 1", tm1).
					AddRow(ID2, "title 2", "description 2", tm2)
				mock.ExpectQuery("SELECT (.+) FROM ToDo").WillReturnRows(rows)
			},
			want: &pb.ReadAllResponse{
				Api: "v1",
				ToDos: []*pb.ToDo{
					{
						Id:          ID,
						Title:       "title 1",
						Description: "description 1",
						Reminder:    reminder1,
					},
					{
						Id:          ID2,
						Title:       "title 2",
						Description: "description 2",
						Reminder:    reminder2,
					},
				},
			},
		},
		{
			name: "Empty",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.ReadAllRequest{
					Api: "v1",
				},
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"ID", "Title", "Description", "Reminder"})
				mock.ExpectQuery("SELECT (.+) FROM ToDo").WillReturnRows(rows)
			},
			want: &pb.ReadAllResponse{
				Api:   "v1",
				ToDos: []*pb.ToDo(nil),
			},
		},
		{
			name: "Unsupported API",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.ReadAllRequest{
					Api: "v1",
				},
			},
			mock:    func() {},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.ssv.ReadAll(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToDoService.ReadAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToDoService.ReadAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
