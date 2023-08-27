package todo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
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

func (m *MysqlRepository) Read(ctx context.Context, id uuid.UUID) (*Todo, error) {
	// get SQL connection from pool
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = c.Close() }()

	rows, err := c.QueryContext(ctx, "SELECT `ID`, `Title`, `Description`, `Reminder` FROM ToDo WHERE `ID`=?", id.String())
	if err != nil {
		return nil, fmt.Errorf("failed to select from ToDo-> " + err.Error())
	}
	defer func() { _ = rows.Close() }()

	if !rows.Next() {
		if err = rows.Err(); err != nil {
			return nil, fmt.Errorf("failed to retrieve data from ToDo-> " + err.Error())
		}
		return nil, fmt.Errorf("ToDo with ID='%d' is not found", id)
	}

	todo := &Todo{}
	if err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Reminder); err != nil {
		return nil, fmt.Errorf("failed to retrieve field values from ToDo row-> " + err.Error())
	}

	if rows.Next() {
		return nil, fmt.Errorf("found multiple ToDo rows with ID='%d'", id)
	}

	return todo, nil
}

func (m *MysqlRepository) Create(ctx context.Context, todo *Todo) (*Todo, error) {
	// get SQL connection from pool
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = c.Close() }()

	_, err = c.ExecContext(ctx, "INSERT INTO ToDo(`ID`, `Title`, `Description`, `Reminder`) VALUES(?, ?, ?, ?)",
		todo.ID, todo.Title, todo.Description, todo.Reminder)
	if err != nil {
		return nil, fmt.Errorf("failed to insert into ToDo-> " + err.Error())
	}

	return todo, nil
}

func (m *MysqlRepository) Update(ctx context.Context, todo *Todo) (int64, error) {
	// get SQL connection from pool
	c, err := m.connect(ctx)
	if err != nil {
		return 0, err
	}
	defer func() { _ = c.Close() }()

	res, err := c.ExecContext(ctx, "UPDATE ToDo SET `Title`=?, `Description`=?, `Reminder`=? WHERE `ID`=?",
		todo.Title, todo.Description, todo.Reminder, todo.ID)
	if err != nil {
		return 0, fmt.Errorf("failed to update ToDo-> " + err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve rows affected value-> " + err.Error())
	}

	if rows == 0 {
		return 0, fmt.Errorf("ToDo with ID='%d' is not found", todo.ID)
	}

	return rows, nil
}

func (m *MysqlRepository) Delete(ctx context.Context, id uuid.UUID) (int64, error) {
	// get SQL connection from pool
	c, err := m.connect(ctx)
	if err != nil {
		return 0, err
	}
	defer func() { _ = c.Close() }()

	// delete ToDo
	res, err := c.ExecContext(ctx, "DELETE FROM ToDo WHERE `ID`=?", id)
	if err != nil {
		return 0, fmt.Errorf("failed to delete ToDo-> " + err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve rows affected value-> " + err.Error())
	}

	if rows == 0 {
		return 0, fmt.Errorf("ToDo with ID='%d' is not found", id)
	}

	return rows, nil
}

func (m *MysqlRepository) ReadAll(ctx context.Context) ([]*Todo, error) {
	// get SQL connection from pool
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = c.Close() }()

	rows, err := c.QueryContext(ctx, "SELECT `ID`, `Title`, `Description`, `Reminder` FROM ToDo")
	if err != nil {
		return nil, fmt.Errorf("failed to select from ToDo-> " + err.Error())
	}
	defer func() { _ = rows.Close() }()

	var todos []*Todo
	for rows.Next() {
		var todo = &Todo{}
		if err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Reminder); err != nil {
			return nil, fmt.Errorf("failed to retrieve field values from ToDo row-> " + err.Error())
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to retrieve data from ToDo-> " + err.Error())
	}

	return todos, nil
}
