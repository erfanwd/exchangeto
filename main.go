package main

import (
	"fmt"
	"log"

	"github.com/erfanwd/exchangeto/clients"
	"github.com/erfanwd/exchangeto/config"
	"github.com/erfanwd/exchangeto/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	bot := clients.Init()
	handlers.Trigger(bot)
	handlers.Init(bot)

	fmt.Println(handlers.Exchanges)
	log.Fatal(app.Listen(":" + config.Config("PORT")))
}
