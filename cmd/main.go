package main

import (
	"log"

	"github.com/evgeniy-dammer/simpleapi/pkg/common/config"
	"github.com/evgeniy-dammer/simpleapi/pkg/common/db"
	"github.com/evgeniy-dammer/simpleapi/pkg/products"

	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfiguration()

	if err != nil {
		log.Fatalln("Configuration faild", err)
	}

	h := db.Init(&c)
	app := fiber.New()

	products.RegisterRoutes(app, h)

	app.Listen(c.Port)
}
