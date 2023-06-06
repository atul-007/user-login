package helper

import (
	"context"
	"fmt"
	"log"

	"github.com/atul-007/user-login/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	Init()
}

func Signin(user models.User) interface{} {
	var result models.User

	err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "username", Value: user.UserName}}).Decode(&result)

	fmt.Println(err)
	if err == mongo.ErrNoDocuments {
		inserted, err := collection.InsertOne(context.Background(), user)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("inserted one  id:", inserted.InsertedID)
		return inserted.InsertedID
	} else {
		fmt.Println("username already exists")
		return "username already exists"
	}

}
