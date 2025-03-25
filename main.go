package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	books = append(books, Book{ID: 1, Title: "Pubest", Author: "Pubest"})
	books = append(books, Book{2, "Time", "Meka"})

	app.Get("/", getBooks)
	app.Get("/:id", getBook)
	app.Post("/", createBook)
	app.Put("/:id", updateBook)
	app.Delete("/:id", deleteBook)
	app.Post("/upload", uploadFile)
	app.Get("/html", renderHtml)

	log.Fatal(app.Listen(":3000"))
}

func uploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	err = c.SaveFile(file, "./upload/"+file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendString("File upload successfully: " + file.Filename)
}

func renderHtml(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":       "Go Fiber Template Example",
		"Description": "An example template",
		"Greeting":    "Hello, world!",
	})
}
