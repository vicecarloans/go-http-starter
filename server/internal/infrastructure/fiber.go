package infrastructure

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Run() {
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


	app.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())
		
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status": "fail",
			"message": errorMessage,
		})
	})

	log.Fatal(app.Listen(":8080"))
}
