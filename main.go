package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"go-wiki-micro/data"
	"log"
	"os"
	"time"
)

func main() {
	app := fiber.New()

	// Data
	//conn, err := data.GetConnection()

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println(err)
	}

	println(conn.Config().Config.ValidateConnect)

	//tx, err := conn.Begin(context.Background())
	if err != nil {
		log.Println(err)
	}
	pages := make([]data.Page, 0)
	rows, err := conn.Query(context.Background(), "select * from blogs")
	if err != nil {
		log.Println(err)
	}
	log.Println(rows.Values())
	//defer rows.Close()

	for rows.Next() {
		page := data.Page{}
		err2 := rows.Scan(&page.Title, &page.Body, &page.CreatedAt, &page.EditedAt)
		if err2 != nil {
			log.Fatalln(err2)
		}
		pages = append(pages, page)
	}
	log.Println(pages)
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

	app.Listen(":4000")
}
