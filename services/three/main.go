package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jobrunner/sampletogo/libs/hello"
)

type Sample struct {
	hello.Greeting
}

func main() {
	app := fiber.New()
	app.Get("/three/greeter/:uuid<guid>", GetSample)
	log.Fatal(app.Listen(":8030"))
}

func GetSample(c *fiber.Ctx) error {
	sampleId := c.Params("uuid")
	audience, err := hello.NewAudience(sampleId)

	if err != nil {
		return c.Status(404).JSON(Sample{})
	}

	return c.Status(200).JSON(audience.Greet())
}
