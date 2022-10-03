package main

import (
	"address-details/config"
	"address-details/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", UpStatus)
	routes.UserRoute(app)
	routes.ContactRoute(app)
	routes.RelationshipRoute(app)
	app.Listen(":8080")
}

func UpStatus(c *fiber.Ctx) error {

	// response := make(map[string]string)
	// response["status"] = "UP"
	// response["contact"] = "GodwinTheophilus"

	response := map[string]string{
		"status":  "UP",
		"contact": "GodwinTheophilus",
	}
	// err := c.SendString(response)
	data := config.LoadConfig()
	fmt.Println(data)
	err := c.JSON(response)
	return err

}
