package services

import (
	"errors"
	"merchant-bank-api/models"
	"merchant-bank-api/repository"
	"time"
)

var loggedInCustomer *models.Customer

func Login(name, password string) (*models.Customer, error) {
    customers, err := repository.ReadCustomers()
    if err != nil {
        return nil, err
    }

    for _, customer := range customers {
        if customer.Name == name && customer.Password == password {
            loggedInCustomer = &customer

            // Log aktivitas login
            history := models.History{
                CustomerID: customer.ID,
                Action:     "Login",
                Timestamp:  time.Now().Format(time.RFC3339),
            }
            repository.AddHistory(history)

            return &customer, nil
        }
    }
    return nil, errors.New("invalid credentials")
}

func Logout() {
    if loggedInCustomer != nil {
        // Log aktivitas logout
        history := models.History{
            CustomerID: loggedInCustomer.ID,
            Action:     "Logout",
            Timestamp:  time.Now().Format(time.RFC3339),
        }
        repository.AddHistory(history)

        loggedInCustomer = nil
    }
}

func GetLoggedInCustomer() *models.Customer {
    return loggedInCustomer
}
