package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/atul-007/user-login/router"
)

func main() {
	fmt.Println("Hello World!")
	r := router.Router()
	fmt.Println("server is getting started")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000...")
}
