package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Aman123at/usermanage/route"
)

func main() {
	fmt.Println("\nWelcome to user management api")
	r := route.Router()
	fmt.Println("Server stating at port 4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}
