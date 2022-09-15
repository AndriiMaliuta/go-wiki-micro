package main

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/rest/api/*", func(ctx *fiber.Ctx) error {
		page := Page{
			Title:     "",
			Body:      "asdasdasd",
			CreatedAt: time.Now(),
		}
		//jsonData, err := json.Marshal(page)
		//if err != nil {
		//	log.Fatalln(err)
		//}
		return ctx.JSON(page)
	})

	app.Listen(":4000")
}
