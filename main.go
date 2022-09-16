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
	rows, err := conn.Query(context.Background(), "select * from pages limit 100", nil)
	for rows.Next() {
		page := data.Page{}
		rows.Scan(&page.Title, &page.Body, &page.CreatedAt, &page.EditedAt)
		pages = append(pages, page)
	}

	defer conn.Close(context.Background())

	// APIs
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/rest/api/*", func(ctx *fiber.Ctx) error {
		page := data.Page{
			Title:     "",
			Body:      "asdasdas asd asdasd abc",
			CreatedAt: time.Now(),
		}
		//jsonData, err := json.Marshal(page)
		//if err != nil {
		//	log.Fatalln(err)
		//}
		return ctx.JSON(page)
	})

	log.Println(pages)

	app.Listen(":4000")
}
