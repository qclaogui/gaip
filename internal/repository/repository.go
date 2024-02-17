// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package repository

import (
	"github.com/pkg/errors"
	"github.com/qclaogui/gaip/genproto/bookstore/apiv1alpha1/bookstorepb"
	"github.com/qclaogui/gaip/genproto/generativelanguage/apiv1/generativelanguagepb"
	"github.com/qclaogui/gaip/genproto/library/apiv1/librarypb"
	"github.com/qclaogui/gaip/genproto/project/apiv1/projectpb"
	"github.com/qclaogui/gaip/genproto/routeguide/apiv1/routeguidepb"
	"github.com/qclaogui/gaip/genproto/showcase/apiv1beta1/showcasepb"
	"github.com/qclaogui/gaip/genproto/todo/apiv1/todopb"
	"github.com/qclaogui/gaip/internal/repository/external"
	"github.com/qclaogui/gaip/internal/repository/memory"
	"github.com/qclaogui/gaip/internal/repository/mysql"
	"github.com/qclaogui/gaip/internal/repository/postgres"
	"github.com/qclaogui/gaip/pkg/service/generativeai"
)

func NewBookstore(cfg Config) (bookstorepb.BookstoreServiceServer, error) {
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

func NewGenerativeai(cfg generativeai.Config) (generativelanguagepb.GenerativeServiceServer, error) {
	return external.NewGenerativeai(cfg.APIKey)
}

func NewLibrary(cfg Config) (librarypb.LibraryServiceServer, error) {
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

func NewProject(cfg Config) (projectpb.ProjectServiceServer, error) {
	switch cfg.Driver {
	case "":
		return nil, errors.Errorf("empty database driver %s", cfg.Driver)
	case DriverMemory:
		return memory.NewProject()
	default:
		return nil, errors.Errorf("unsupported driver for database %s", cfg.Driver)
	}
}

func NewIdentity(cfg Config) (showcasepb.IdentityServiceServer, error) {
	switch cfg.Driver {
	case "":
		return nil, errors.Errorf("empty database driver %s", cfg.Driver)
	case DriverMemory:
		return memory.NewIdentity()
	default:
		return nil, errors.Errorf("unsupported driver for database %s", cfg.Driver)
	}
}

func NewMessaging(cfg Config) (showcasepb.MessagingServiceServer, error) {
	switch cfg.Driver {
	case "":
		return nil, errors.Errorf("empty database driver %s", cfg.Driver)
	case DriverMemory:
		return memory.NewMessaging()
	default:
		return nil, errors.Errorf("unsupported driver for database %s", cfg.Driver)
	}
}

func NewRouteGuide(cfg Config) (routeguidepb.RouteGuideServiceServer, error) {
	switch cfg.Driver {
	case "":
		return nil, errors.Errorf("empty database driver %s", cfg.Driver)
	case DriverMemory:
		return memory.NewRouteNote(cfg.MemoryCfg)
	default:
		return nil, errors.Errorf("unsupported driver for database %s", cfg.Driver)
	}
}

func NewTodo(cfg Config) (todopb.ToDoServiceServer, error) {
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
