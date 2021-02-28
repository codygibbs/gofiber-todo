package handler

import (
	"github.com/codygibbs/gofiber-todo/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Ctx) ListTodos(c *fiber.Ctx) error {
	todos, err := h.Model.FindAllTodos()
	if err != nil {
		return err
	}

	return c.Render("page/home", fiber.Map{
		"Todos": todos,
	}, "template/anonymous")
}

func (h *Ctx) CreateTodo(c *fiber.Ctx) error {
	todo := model.Todo{
		ID:          uuid.New(),
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
	}

	err := h.Model.CreateTodo(todo)
	if err != nil {
		return err
	}

	return c.Redirect("/")
}
