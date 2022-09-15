package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-wiki-micro/data"
	"log"
	"time"
)

func main() {
	app := fiber.New()

	// Data
	conn, err := data.GetConnection()
	if err != nil {
		log.Fatalln(err)
	}

	pages := make([]data.Page, 0)
	conn.QueryRow(context.Background(), "select * from pages limit 100").Scan()
	defer conn.Close(context.Background())

	// APIs
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/rest/api/*", func(ctx *fiber.Ctx) error {
		page := data.Page{
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
