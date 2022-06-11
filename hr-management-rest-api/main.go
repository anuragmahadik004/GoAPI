package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anuragmahadik004/hr_api/routers"
)

func main() {
	fmt.Println("Welcome to HR API")
	fmt.Println("Server is getting started...")

	r := routers.Router()

	log.Fatal(http.ListenAndServe(":4444", r))

	fmt.Println("Listening at port 4444...")
}
