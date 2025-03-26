package main

import "github.com/gofiber/fiber/v2"

func renderHtml(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":       "Go Fiber Template Example",
		"Description": "An example template",
		"Greeting":    "Hello, world!",
	})
}
