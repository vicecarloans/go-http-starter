package infrastructure

import (
	"fmt"
	"go-http-server/internal/authors"
	"go-http-server/internal/books"
	"go-http-server/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Run() {
	config := utils.LoadConfig()

	db := ConnectToPostgres(config)

	app := fiber.New(fiber.Config{
		AppName: "Go HTTP Server",
		ServerHeader: "Fiber",
	})

	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "pong",
		})
	})


	// Create repositories
	bookRepository := books.NewBookRepository(db)
	authorRepository := authors.NewAuthorRepository(db)
	
	// Create services
	bookService := books.NewBookService(bookRepository)
	authorService := authors.NewAuthorService(authorRepository)

	books.NewBookHandler(app.Group("/api/v1/books"), bookService)
	authors.NewAuthorHandler(app.Group("/api/v1/authors"), authorService)

	app.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())
		
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status": "fail",
			"message": errorMessage,
		})
	})

	log.Fatal(app.Listen(":8080"))
}
