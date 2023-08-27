package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"peerswap/ad"
	"peerswap/reusable"
)

func main() {
	err := reusable.ConnectDefaultMongo(os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_NAME"))
	if err != nil {
		log.Fatalf(err.Error())
	}

	app := fiber.New()
	ad.NewModule(app).Register()

	app.Listen(":3000")
}
