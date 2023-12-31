package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"peerswap/ad"
	"peerswap/order"
	"peerswap/reusable"
)

func main() {
	err := reusable.ConnectDefaultMongo(os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_NAME"))
	if err != nil {
		log.Fatalf(err.Error())
	}

	app := fiber.New()
	event := reusable.NewMapEvent()

	ad.NewModule(app, event).Register()
	order.NewModule(app, event).Register()

	app.Listen(":3000")
}
