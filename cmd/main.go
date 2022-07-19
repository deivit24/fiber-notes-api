package main

import (
	"log"

	"github.com/deivit24/fiber-notes-api/pkg/common/config"
	"github.com/deivit24/fiber-notes-api/pkg/common/db"
	"github.com/deivit24/fiber-notes-api/pkg/products"
	"github.com/deivit24/fiber-notes-api/pkg/users"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(&c)
	app := fiber.New()

	app.Use(cors.New(cors.Config{
    AllowOrigins: "*",
    AllowHeaders:  "Origin, Content-Type, Accept",
	}))
	products.RegisterRoutes(app, h)
	users.RegisterRoutes(app, h)

	app.Listen(c.Port)
}
