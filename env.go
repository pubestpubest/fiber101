package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func getConfig(c *fiber.Ctx) error {
	secret := getEnv("SECRET", "default")
	return c.JSON(fiber.Map{
		"SERECT": secret,
	})
}

func getDotenv(c *fiber.Ctx) error {
	dotenvsecret := os.Getenv("DOTENVSECRET")
	if dotenvsecret == "" {
		dotenvsecret = "default"
	}
	return c.JSON(fiber.Map{
		"DOTENVSECRET": dotenvsecret,
	})
}

func getEnv(key, fallback string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	} else {
		return fallback
	}
}
