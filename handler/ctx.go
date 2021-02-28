package handler

import (
	"github.com/codygibbs/gofiber-todo/model"
	"github.com/gofiber/fiber/v2"
)

type Ctx struct {
	Model model.Model
}

type Handler interface {
	ListTodos(c *fiber.Ctx) error
	CreateTodo(c *fiber.Ctx) error
}

func New(m model.Model) Handler {
	return &Ctx{
		Model: m,
	}
}
