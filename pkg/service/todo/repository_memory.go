package todo

import (
	"context"
	"sync"

	"github.com/google/uuid"
)

// MemoryRepository fulfills the Repository interface
type MemoryRepository struct {
	todos map[uuid.UUID]*Todo
	mu    sync.Mutex
}

// NewMemoryRepository is a factory function to generate a new repository
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		todos: make(map[uuid.UUID]*Todo),
	}
}

func (m *MemoryRepository) Create(_ context.Context, td *Todo) (*Todo, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if todo, ok := m.todos[td.ID]; ok {
		return todo, nil
	}

	if td.Title == "" && td.Description == "" {
		return nil, ErrTodoFailedToCreate
	}

	td.ID = uuid.NewSHA1(uuid.NameSpaceOID, []byte(td.Title+td.Description))

	// Create
	m.todos[td.ID] = td
	return td, nil

}

func (m *MemoryRepository) Read(_ context.Context, id uuid.UUID) (*Todo, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if td, ok := m.todos[id]; ok {
		return td, nil

	}
	return nil, ErrTodoNotFound
}

func (m *MemoryRepository) Update(_ context.Context, td *Todo) (int64, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.todos[td.ID]; !ok {
		return 0, ErrTodoNotFound
	}

	// Update
	m.todos[td.ID] = td
	return 1, nil
}

func (m *MemoryRepository) Delete(_ context.Context, id uuid.UUID) (int64, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.todos[id]; !ok {
		return 0, ErrTodoNotFound
	}

	// Delete
	delete(m.todos, id)
	return 1, nil
}

func (m *MemoryRepository) ReadAll(_ context.Context) ([]*Todo, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	var tds []*Todo
	for _, td := range m.todos {
		tds = append(tds, td)
	}

	if len(tds) < 1 {
		return nil, ErrTodoNotFound
	}

	return tds, nil
}
