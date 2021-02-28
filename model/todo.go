package model

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID
	IsComplete  bool
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (m *Ctx) FindAllTodos() (*[]Todo, error) {
	rows, err := m.db.Query("SELECT id, is_complete, title, description, created_at, updated_at FROM todos order by created_at")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []Todo{}

	for rows.Next() {
		todo := Todo{}
		if err := rows.Scan(&todo.ID, &todo.IsComplete, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			return nil, err // Exit if we get an error
		}
		todos = append(todos, todo)
	}

	return &todos, nil
}

func (m *Ctx) FindTodo(id uuid.UUID) (*Todo, error) {
	rows, err := m.db.Query("SELECT id, is_complete, title, description, created_at, updated_at FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todo := Todo{}

	for rows.Next() {
		if err := rows.Scan(&todo.ID, &todo.IsComplete, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			return nil, err // Exit if we get an error
		}
	}

	return &todo, nil
}

func (m *Ctx) CreateTodo(t Todo) error {
	_, err := m.db.Query("INSERT INTO todos (id, is_complete, title, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)", t.ID, t.IsComplete, t.Title, t.Description, t.CreatedAt, t.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
