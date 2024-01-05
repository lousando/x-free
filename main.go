package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{
		Views:       html.New("./views", ".tmpl"),
		ViewsLayout: "layouts/default",
	})

	app.Get("/", func(context *fiber.Ctx) error {
		return context.Render("index", fiber.Map{})
	})

	log.Fatal(app.Listen(":8080"))
}
