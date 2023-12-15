package repository

import (
	"flag"
	"fmt"
	"slices"
	"strings"

	"github.com/pkg/errors"
	"github.com/qclaogui/gaip/genproto/todo/apiv1/todopb"
)

const (
	BackendMemory = "memory"
	BackendMysql  = "mysql"
)

var (
	supportedDatabaseBackends = []string{BackendMemory, BackendMysql}

	// ErrNotFound is returned when a item is not found.
	ErrNotFound = errors.New("the item was not found in the repository")

	// ErrFailedToCreate is returned when a item is create Failed
	ErrFailedToCreate = errors.New("failed to add the todo to the repository")
)

type Repository interface {
	todopb.ToDoServiceServer
}

type Config struct {
	Backend string `yaml:"backend"`

	Memory MemoryConfig `yaml:"memory"`
	Mysql  MysqlConfig  `yaml:"mysql"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	prefix := "todo.database."
	fs.StringVar(&cfg.Backend, prefix+"backend", BackendMemory, fmt.Sprintf("Backend storage to use. Supported backends are: %s.", strings.Join(supportedDatabaseBackends, ", ")))

	cfg.Memory.RegisterFlagsWithPrefix(prefix, fs)
	cfg.Mysql.RegisterFlagsWithPrefix(prefix, fs)
}

func (cfg *Config) Validate() error {
	if cfg.Backend != "" && !slices.Contains(supportedDatabaseBackends, cfg.Backend) {
		return fmt.Errorf("unsupported RepoCfg backend: %s", cfg.Backend)
	}

	return nil
}

func NewRepository(cfg Config) (Repository, error) {
	switch cfg.Backend {
	case "":
		return nil, errors.Errorf("empty database backend %s", cfg.Backend)
	case BackendMemory:
		return NewMemoryRepo(), nil
	case BackendMysql:
		return NewMysqlRepo(cfg.Mysql)
	default:
		return nil, errors.Errorf("unsupported backend for database %s", cfg.Backend)
	}
}
