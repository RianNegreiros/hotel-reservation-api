package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "Listen address")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/alive", handleAlive)
	app.Listen(*listenAddr)
}

func handleAlive(c *fiber.Ctx) error {
	return c.JSON((map[string]string{"message": "I'm alive!"}))
}
