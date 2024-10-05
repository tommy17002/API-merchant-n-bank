package controllers

import (
	"encoding/json"
	"merchant-bank-api/models"
	"merchant-bank-api/services"
	"net/http"
)

func Payment(w http.ResponseWriter, r *http.Request) {
    var paymentRequest models.PaymentRequest
    json.NewDecoder(r.Body).Decode(&paymentRequest)

    message, err := services.Payment(paymentRequest.MerchantID, paymentRequest.Amount)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.Write([]byte(message))
}
