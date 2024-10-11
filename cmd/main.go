package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kudaibergenoff/mybankapi/routes"
	"log"

	"github.com/gofiber/swagger"
	_ "github.com/kudaibergenoff/mybankapi/docs"
)

// @title MyBankAPI
// @version 1.0
// @description This is a sample API for bank transactions.
// @host localhost:8080
// @BasePath /api
func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	r := routes.NewRoute(app)
	r.Register()

	app.Get("/swagger/*", swagger.HandlerDefault)

	log.Fatal(app.Listen(":8080"))
}
