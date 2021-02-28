package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type Ctx struct {
	db *sql.DB
}

type Model interface {
	CreateTodo(t Todo) error
	FindAllTodos() (*[]Todo, error)
	FindTodo(id uuid.UUID) (*Todo, error)
}

func New(db *sql.DB) Model {
	return &Ctx{
		db: db,
	}
}
