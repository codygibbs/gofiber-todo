package main

import (
	"fmt"
	"os"

	"log"

	"github.com/codygibbs/gofiber-todo/application"
	"github.com/codygibbs/gofiber-todo/handler"
	"github.com/codygibbs/gofiber-todo/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Could not find .env file. Did you mean to load one?")
	}

	err = application.Run(func(app application.Application) {
		model := model.New(app.GetDBConn())
		handler := handler.New(model)

		engine := html.New("./view", ".gohtml")

		fiber := fiber.New(fiber.Config{
			Views: engine,
		})

		fiber.Use(logger.New())

		fiber.Get("/", handler.ListTodos)

		fiber.Post("/todos", handler.CreateTodo)

		fiber.Static("/assets", "./public")

		fiber.Listen(":3000")
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
