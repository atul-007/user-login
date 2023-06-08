package helper

import (
	"context"
	"fmt"
	"log"

	"github.com/atul-007/user-login/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	Init()
}

func Register(user models.User) interface{} {
	var result models.User

	pass, errr := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if errr != nil {
		panic(errr)
	}

	user.Password = string(pass)

	err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "username", Value: user.UserName}}).Decode(&result)

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
func Login(user models.User) interface{} {

	var result models.User

	collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "username", Value: user.UserName}}).Decode(&result)
	errf := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
	if errf == nil {

		err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "username", Value: user.UserName}, primitive.E{Key: "password", Value: result.Password}}).Decode(&result)

		if err == mongo.ErrNoDocuments {

			fmt.Println("invalid username or passeword")
			return "invalid username or password"
		} else {
			fmt.Println("logged-in")
			return "logged-in"
		}
	} else {
		return bcrypt.ErrMismatchedHashAndPassword
	}

}
func GetUserData(userName string) interface{} {
	var result models.User
	err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "username", Value: userName}}).Decode(&result) //to get all the values in the cursor

	if err != nil {
		log.Fatal(err)
	}

	//var data []primitive.M //similar to  bson.m but better in this case
	//data = result
	/*for cursor.Next(context.Background()) { //to get all the values in the cursor
		var udata bson.M
		err := cursor.Decode(&udata)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, udata)
	}*/
	//defer cursor.Close(context.Background())
	return result
}
