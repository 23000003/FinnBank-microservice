package models

type Account struct {
	Email       string `json:"email"`
	Full_Name   string `json:"full_name"`
	Phone       string `json:"phone_number"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	AccountType string `json:"account_type"`
}