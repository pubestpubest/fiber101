package main

import (
	"fmt"
	"log"
	"os"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var defaultUser = User{
	Email:    "email@email.com",
	Password: "email",
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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

	app.Post("/login", login)
	app.Post("/login-jwt", jwtLogin)
	app.Get("/backdoor", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": os.Getenv("JWT_SECRET"),
		})
	})
	app.Use(logMiddleware)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(os.Getenv("JWT_SECRET")),
		},
	}))
	app.Get("/", getBooks)
	app.Get("/:id", getBook)
	app.Post("/", createBook)
	app.Put("/:id", updateBook)
	app.Delete("/:id", deleteBook)
	app.Post("/upload", uploadFile)
	app.Get("/html", renderHtml)
	app.Get("/api/config", getConfig)
	app.Get("/api/dotenv", getDotenv)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}

func jwtLogin(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if user.Email != defaultUser.Email || user.Password != defaultUser.Password {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	claims := jwt.MapClaims{
		"email":    user.Email,
		"password": user.Password,
		"role":     "admin",
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{
		"token": t,
	})
}

func login(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}
	if user.Email != defaultUser.Email || user.Password != defaultUser.Password {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.JSON(fiber.Map{
		"message": "Login successful",
	})
}

func logMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	duration := time.Since(start)
	fmt.Printf("Request URL: %s\nMethod: %s\nDuration: %s\n", c.OriginalURL(), c.Method(), duration)
	return err
}
