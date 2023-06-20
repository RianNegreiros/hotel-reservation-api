package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbUri = "mongodb://localhost:27017"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(client)

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
