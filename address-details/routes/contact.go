package routes

import (
	"address-details/db"
	"address-details/model"
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var contactCollection *mongo.Collection = db.GetCollection(db.DB, "contacts")

func ContactRoute(app *fiber.App) {
	contactGroup := app.Group("contact", logger.New())
	contactGroup.Post("/getAllContacts", getAllContacts)
	contactGroup.Post("/saveContact", saveAllContacts)
}

func saveAllContacts(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var contacts []model.MContact
	err := c.BodyParser(&contacts)
	if err != nil {
		return err
	}

	contactInterface := make([]interface{}, len(contacts))

	for i, v := range contacts {
		contactInterface[i] = v
	}

	result, err := contactCollection.InsertMany(ctx, contactInterface)
	if err != nil {
		fmt.Println(result)
		return err
	}
	c.JSON(map[string]string{
		"message": "Saved data successfully!!!",
	})
	return nil
}

func getAllContacts(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var userFilter map[string]string
	err := c.BodyParser(&userFilter)
	if err != nil {
		return err
	}
	fmt.Println(userFilter)

	var contacts []model.MContact
	cursor, err := contactCollection.Find(ctx, bson.D{{"emailid", userFilter["emailid"]}})
	// userFilter["emailid"]
	if err != nil {
		return err
	}

	for cursor.Next(ctx) {
		var queryResult model.MContact
		if err := cursor.Decode(&queryResult); err != nil {
			return err
		}
		fmt.Println(queryResult)
		contacts = append(contacts, queryResult)
	}
	c.Status(200).JSON(contacts)
	return nil
}
