package routes

import (
	"merchant-bank-api/controllers"
	"net/http"
)

func RegisterRoutes() {
    http.HandleFunc("/login", controllers.Login)
    http.HandleFunc("/logout", controllers.Logout)
    http.HandleFunc("/payment", controllers.Payment) // Tambahkan ini untuk payment
}
