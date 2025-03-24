package main

import (
	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {
	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "Pubest", Author: "Pubest"})
	books = append(books, Book{2, "Time", "Meka"})

	app.Get("/", getBooks)

	app.Listen(":3000")
}

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}
