package controllers

import (
	"encoding/json"
	"merchant-bank-api/models"
	"merchant-bank-api/services"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
    var loginRequest models.LoginRequest
    json.NewDecoder(r.Body).Decode(&loginRequest)

    customer, err := services.Login(loginRequest.Name, loginRequest.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    json.NewEncoder(w).Encode(customer)
}

func Logout(w http.ResponseWriter, r *http.Request) {
    services.Logout()
    w.Write([]byte("Logged out successfully"))
}
