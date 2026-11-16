// REST API with the Fiber framework that adds two numbers.
// POST http://localhost:3000/add  body: {"num1": 10, "num2": 20}
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Request struct {
	Num1 float64 `json:"num1"`
	Num2 float64 `json:"num2"`
}

func addNumbers(c *fiber.Ctx) error {
	var req Request

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	result := req.Num1 + req.Num2

	return c.JSON(fiber.Map{
		"num1":   req.Num1,
		"num2":   req.Num2,
		"result": result,
	})
}

func main() {
	app := fiber.New()

	app.Post("/add", addNumbers)

	log.Fatal(app.Listen(":3000"))
}
