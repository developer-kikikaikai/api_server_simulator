package main

import (
	"net/http"

	"github.com/developer-kikikaikai/api_server_simulator/controllers"
)

func main() {
	router := controllers.CreateRouter()
	http.ListenAndServe(":8080", router)
}
