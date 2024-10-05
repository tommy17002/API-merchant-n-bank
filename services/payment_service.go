package services

import (
	"errors"
	"merchant-bank-api/models"
	"merchant-bank-api/repository"
	"time"
)

func Payment(merchantID string, amount float64) (string, error) {
    customer := GetLoggedInCustomer()
    if customer == nil {
        return "", errors.New("customer is not logged in")
    }

    if amount <= 0 {
        return "", errors.New("amount must be greater than zero")
    }

    if customer.Balance < amount {
        return "", errors.New("insufficient balance")
    }

    // Mengurangi saldo pelanggan
    customer.Balance -= amount

    // Simpan perubahan saldo pelanggan ke dalam file JSON
    err := repository.UpdateCustomerBalance(customer)
    if err != nil {
        return "", err
    }

    // Menambahkan log ke history pembayaran
    history := models.History{
        ID:         time.Now().String(),
        CustomerID: customer.ID,
        Action:     "Payment",
        Amount:     amount,
        Timestamp:  time.Now().Format(time.RFC3339),
    }

    err = repository.AddHistory(history)
    if err != nil {
        return "", err
    }

    return "Payment successful", nil
}
