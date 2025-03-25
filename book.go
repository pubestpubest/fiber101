package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	bookId, error := strconv.Atoi(c.Params("id"))
	if error != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	for _, book := range books {
		if book.ID == bookId {
			return c.JSON(book)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

func createBook(c *fiber.Ctx) error {
	newBook := new(Book)
	if err := c.BodyParser(newBook); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	newBook.ID = len(books) + 1
	books = append(books, *newBook)

	return c.JSON(newBook)
}

func deleteBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	for index, deletedBook := range books {
		if bookId == deletedBook.ID {
			books = append(books[:index], books[index+1:]...)
			return c.Status(fiber.StatusAccepted).JSON(deletedBook)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)

}

func updateBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	bodyBook := new(Book)
	if err := c.BodyParser(bodyBook); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	for i, updateBook := range books {
		if bookId == updateBook.ID {
			updateBook.Title = bodyBook.Title
			updateBook.Author = bodyBook.Author
			books[i] = updateBook
			return c.Status(fiber.StatusAccepted).JSON(updateBook)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}
