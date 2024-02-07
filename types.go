package main

import (
	"math/rand"
	"time"
)

type Account struct {
	ID            int    `json: "id`
	FirstName     string `json: "fname"`
	LastName      string `json: "lname"`
	Number int64  `json:"number"`
	Password      string `json: "password"`
	Balance       int64  `json: "balance"`
	CreatedAt      time.Time	`json: "createdAt"`
}

type LoginResponse struct{
	Number	int64	`json: "number"`
}

type LoginRequest struct{
	Number	int64	`json: "number"`
	Password	string	`json: "password"`
}

type TranferRequest struct{
	ToAccount	int `json: "toAccount"`
	Amount	int `json: "amount"`
}

type CreateAccountRequest struct{
	FirstName	string	`json: "firstName"`
	LastName	string	`json: "lastName"`
	Password	string	`json: "password"`	
}


func NewAccount (firstName, lastName, password string) (*Account, error){
	return &Account{
		FirstName: firstName,
		LastName: lastName,
		Password: password,
		Number: int64(rand.Intn(1000000)),
		CreatedAt: time.Now(),
	}, nil
}