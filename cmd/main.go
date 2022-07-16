package main

import (
	"log"

	"github.com/deivit24/api-db/pkg/common/config"
	"github.com/deivit24/api-db/pkg/common/db"
	"github.com/deivit24/api-db/pkg/products"
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

	app.Listen(c.Port)
}
