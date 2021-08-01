package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/zeimedee/koa_test/handlers"
)

func Handler(app *fiber.App) {
	app.Post("/api/convert", handlers.Convert)
}

func main() {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "${ip}${url} - ${status} - ${method}\n",
		Output: os.Stdout,
	}))
	Handler(app)

	log.Fatal(app.Listen(":3000"))
}
