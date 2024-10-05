package services

import (
	"errors"
	"merchant-bank-api/models"
	"merchant-bank-api/repository"
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
            return &customer, nil
        }
    }
    return nil, errors.New("invalid credentials")
}

func Logout() {
    loggedInCustomer = nil
}

func GetLoggedInCustomer() *models.Customer {
    return loggedInCustomer
}
