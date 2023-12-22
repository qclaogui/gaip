// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package repository

import (
	"github.com/pkg/errors"
	"github.com/qclaogui/gaip/genproto/bookstore/apiv1alpha1/bookstorepb"
	"github.com/qclaogui/gaip/genproto/library/apiv1/librarypb"
	"github.com/qclaogui/gaip/genproto/project/apiv1/projectpb"
	"github.com/qclaogui/gaip/genproto/routeguide/apiv1/routeguidepb"
	"github.com/qclaogui/gaip/genproto/todo/apiv1/todopb"
	"github.com/qclaogui/gaip/internal/repository/memory"
	"github.com/qclaogui/gaip/internal/repository/mysql"
	"github.com/qclaogui/gaip/internal/repository/postgres"
)

type Bookstore interface {
	bookstorepb.BookstoreServiceServer
}

type Library interface {
	librarypb.LibraryServiceServer
}

type Project interface {
	projectpb.ProjectServiceServer
}

type RouteGuide interface {
	routeguidepb.RouteGuideServiceServer
}

type Todo interface {
	todopb.ToDoServiceServer
}

func NewBookstore(cfg Config) (Bookstore, error) {
	switch cfg.Driver {
	case "":
		return nil, errors.Errorf("empty database drivers %s", cfg.Driver)
	case DriverMemory:
		return memory.NewBookstore()
	case DriverMysql:
		return mysql.NewBookstore(cfg.MysqlCfg)
	default:
		return nil, errors.Errorf("unsupported drivers for database %s", cfg.Driver)
	}
}

func NewLibrary(cfg Config) (Library, error) {
	switch cfg.Driver {
	case "":
		return nil, errors.Errorf("empty database driver %s", cfg.Driver)
	case DriverMemory:
		return memory.NewLibrary()
	case DriverMysql:
		return mysql.NewLibrary(cfg.MysqlCfg)
	case DriverPostgres:
		return postgres.NewLibrary(cfg.PostgresCfg)
	default:
		return nil, errors.Errorf("unsupported driver for database %s", cfg.Driver)
	}
}

func NewProject(cfg Config) (Project, error) {
	switch cfg.Driver {
	case "":
		return nil, errors.Errorf("empty database driver %s", cfg.Driver)
	case DriverMemory:
		return memory.NewProject()
	default:
		return nil, errors.Errorf("unsupported driver for database %s", cfg.Driver)
	}
}

func NewRouteGuide(cfg Config) (RouteGuide, error) {
	switch cfg.Driver {
	case "":
		return nil, errors.Errorf("empty database driver %s", cfg.Driver)
	case DriverMemory:
		return memory.NewRouteNote(cfg.MemoryCfg)
	default:
		return nil, errors.Errorf("unsupported driver for database %s", cfg.Driver)
	}
}

func NewTodo(cfg Config) (Todo, error) {
	switch cfg.Driver {
	case "":
		return nil, errors.Errorf("empty database driver %s", cfg.Driver)
	case DriverMemory:
		return memory.NewTodo()
	case DriverMysql:
		return mysql.NewTodo(cfg.MysqlCfg)
	default:
		return nil, errors.Errorf("unsupported driver for database %s", cfg.Driver)
	}
}
