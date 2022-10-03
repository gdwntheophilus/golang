package routes

import (
	"address-details/db"
	"address-details/model"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection = db.GetCollection(db.DB, "user")

func UserRoute(app *fiber.App) {
	user := app.Group("user", logger.New())
	user.Get("/", GetAllUsers)
	user.Post("/register", RegisterUser)
	user.Post("/getRegisteredUser", GetRegisteredUser)
}

func GetRegisteredUser(c *fiber.Ctx) error {
	var userDetails map[string]string
	err := c.BodyParser(&userDetails)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var queryResult map[string]interface{}
	err = collection.FindOne(ctx, bson.D{{"emailid", userDetails["emailid"]}, {"password", userDetails["password"]}}).Decode(&queryResult)
	if err != nil {
		c.Status(200).JSON(map[string]string{
			"status":  "error",
			"message": "User is not registered, Kindly register and login!!!",
		})
		return nil
	}
	if queryResult["emailid"] == userDetails["emailid"] && queryResult["password"] == userDetails["password"] {
		c.Status(200).JSON(map[string]string{
			"status":  "Ok",
			"message": "Login Successfull!!!",
		})

	} else {
		c.Status(200).JSON(map[string]string{
			"status":  "error",
			"message": "User is not registered, Kindly register and login!!!",
		})
	}
	return nil
}

func GetAllUsers(c *fiber.Ctx) error {
	var user []model.MUser
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	for cursor.Next(ctx) {
		var result model.MUser
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
		user = append(user, result)
	}
	c.Status(200).JSON(user)
	return nil
}

func RegisterUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()
	var userJson model.MUser
	err := c.BodyParser(&userJson)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var queryResult map[string]string
	err = collection.FindOne(ctx, bson.D{{"emailid", userJson.EmailId}}).Decode(&queryResult)
	if err != nil {

		result, err := collection.InsertOne(ctx, userJson)
		if err != nil {
			fmt.Println(err)
			fmt.Println(result)
			return err
		}

		c.Status(200).JSON(map[string]string{
			"status":  "OK",
			"message": "User Registered Successfully!!!!",
		})

	} else {
		c.Status(200).JSON(map[string]string{
			"status":  "Error",
			"message": "User already registered, Kindly login!!!!",
		})
	}

	return nil
}
