package main

import (
	"log"
	"portfolio-backend/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "http://127.0.0.1:5502")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		c.Set("Access-Control-Allow-Methods", "GET, POST")
		return c.Next()
	})

	app.Post("/api/blogs", handlers.CreateBlog)

	app.Get("/api/blogs", handlers.GetAllBlogs)

	log.Fatal(app.Listen(":3000"))
}
