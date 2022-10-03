package routes

import (
	"address-details/db"
	"address-details/model"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var relationshipCollection *mongo.Collection = db.GetCollection(db.DB, "relationships")

func RelationshipRoute(app *fiber.App) {
	relation := app.Group("relationship", logger.New())
	relation.Get("/getAllRelationships", getAllRelationship)
}

func getAllRelationship(c *fiber.Ctx) error {
	var relationships []model.MRelationship
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	cursor, err := relationshipCollection.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	for cursor.Next(ctx) {
		var result model.MRelationship
		err = cursor.Decode(&result)
		if err != nil {
			return err
		}
		relationships = append(relationships, result)

	}
	c.JSON(relationships)
	return nil
}
