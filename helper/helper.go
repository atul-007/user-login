package helper

import (
	"context"
	"fmt"
	"log"

	"github.com/atul-007/user-login/models"
)

func init() {
	Init()
}

func Signin(user models.User) interface{} {
	inserted, err := collection.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("inserted one  id:", inserted.InsertedID)
	return inserted.InsertedID
}
