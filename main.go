package main

import (
	"fmt"
	"net/http"

	"github.com/asibulhasanshanto/restapi/router"
)

func main() {

	r := router.Router()
	fmt.Println("Rest API")
	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", r)
}
