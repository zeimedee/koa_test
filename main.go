package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/koa_test/handlers"
)

func Handler(app *fiber.App) {
	app.Post("/api/convert", handlers.Convert)
}
func main() {
	app := fiber.New()
	Handler(app)
	log.Fatal(app.Listen(":3000"))
}
