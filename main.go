package main

import (
	"log"
	"merchant-bank-api/routes"
	"net/http"
)

func main() {
    routes.RegisterRoutes()
    log.Println("Server running on port 8080")
    http.ListenAndServe(":8080", nil)
}
