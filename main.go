package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/wpcodevo/golang-fiber-mysql/controllers"
	"github.com/wpcodevo/golang-fiber-mysql/initializers"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	app := fiber.New()
	micro := fiber.New()

	app.Mount("/api", micro)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	micro.Route("/notes", func(router fiber.Router) {
		router.Post("/", controllers.CreateNoteHandler)
		router.Get("", controllers.FindNotes)
	})
	micro.Route("/notes/:noteId", func(router fiber.Router) {
		router.Delete("", controllers.DeleteNote)
		router.Get("", controllers.FindNoteById)
		router.Patch("", controllers.UpdateNote)
	})
	micro.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang, Fiber, SQLite, and GORM",
		})
	})

	log.Fatal(app.Listen(":8000"))
}
