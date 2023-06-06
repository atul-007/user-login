package helper

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbname = "userlogin"
const colname = "user_data"

// most important
var collection *mongo.Collection

func Init() {
	err := godotenv.Load() //to load the env file
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	connectionstring := os.Getenv("connect")

	clientoption := options.Client().ApplyURI(connectionstring)

	//connect with mongodb
	client, err := mongo.Connect(context.TODO(), clientoption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongodb connection sucess")

	collection = client.Database(dbname).Collection(colname)

	//collection instance
	fmt.Println("collection instance is ready ")
}
