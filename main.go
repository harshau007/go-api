package main

import (
	"fmt"
	"net/http"
	"github.com/harshau007/go-api/routes"
)

func main()  {
	r := routes.Router()
	r.Get("/", routes.Welcome)
	fmt.Println("Server started......")
	http.ListenAndServe(":3000", r)
}