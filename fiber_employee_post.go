package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// Employee struct
type Employee struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func main() {
	app := fiber.New()

	app.Post("/employee", func(c *fiber.Ctx) error {
		employee := new(Employee)

		if err := c.BodyParser(employee); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid request",
			})
		}

		return c.JSON(fiber.Map{
			"message":  "Employee data received",
			"employee": employee,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
