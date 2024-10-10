package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kudaibergenoff/mybankapi/routes"
	"log"
)

func main() {
	app := fiber.New()
	r := routes.NewRoute(app)

	r.Register()

	log.Fatal(app.Listen(":8080"))
}
