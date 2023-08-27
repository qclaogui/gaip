package cmd

import (
	"context"
	"flag"
	"fmt"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/qclaogui/golang-api-server/pkg/protocol/grpc"
	todov1 "github.com/qclaogui/golang-api-server/pkg/service/todo/v1"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string

	// DB Datastore parameters section
	// DatastoreDBHost is host of database
	DatastoreDBHost string
	// DatastoreDBUser is username to connect to database
	DatastoreDBUser string
	// DatastoreDBPassword password to connect to database
	DatastoreDBPassword string
	// DatastoreDBSchema is schema of database
	DatastoreDBSchema string

	// Log parameters section
	// LogLevel is global log level: Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5)
	LogLevel int
	// LogTimeFormat is print time format for logger e.g. 2006-01-02T15:04:05Z07:00
	LogTimeFormat string
}

// StartServer runs gRPC server and HTTP gateway
func StartServer() error {
	ctx := context.Background()
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "9095", "gRPC port to bind")
	flag.StringVar(&cfg.DatastoreDBHost, "db-host", "127.0.0.1", "Database host")
	flag.StringVar(&cfg.DatastoreDBUser, "db-user", "root", "Database user")
	flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "", "Database password")
	flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", "dev", "Database schema")
	flag.IntVar(&cfg.LogLevel, "log-level", 0, "Global log level")
	flag.StringVar(&cfg.LogTimeFormat, "log-time-format", "2006-01-02T15:04:05.999999999Z07:00",
		"Print time format for logger e.g. 2006-01-02T15:04:05Z07:00")
	flag.Parse()

	// initialize logger
	// if err := logger.Init(cfg.LogLevel, cfg.LogTimeFormat); err != nil {
	// 	return fmt.Errorf("failed to initialize logger: %v", err)
	// }

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	// add MySQL driver specific parameter to parse date/time
	// Drop it for another database

	//param := "parseTime=true"
	//dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
	//	cfg.DatastoreDBUser,
	//	cfg.DatastoreDBPassword,
	//	cfg.DatastoreDBHost,
	//	cfg.DatastoreDBSchema,
	//	param)
	//toDov1, err := td.NewToDoService(td.WithMysqlToDoRepository(dsn))

	toDov1, err := todov1.NewToDoService(todov1.WithMemoryRepository())
	if err != nil {
		return err
	}

	return grpc.RunServer(ctx, toDov1, cfg.GRPCPort)
}
