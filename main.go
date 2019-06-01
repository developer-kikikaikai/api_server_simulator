package main

import (
	"net/http"

	"./controllers"
)

func main() {
	router := controllers.CreateRouter()
	http.ListenAndServe(":8080", router)
}
