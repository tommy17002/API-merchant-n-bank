package models

type Customer struct {
    ID       string  `json:"id"`
    Name     string  `json:"name"`
    Password string  `json:"password"`
    Balance  float64 `json:"balance"`
}
