package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	FirstName 	string `json:"firstName"`
	LastName 	string `json:"lastname"`

}

type Account struct {
	ID 			int			`json:"id"`
	FirstName 	string		`json:"first_name"`
	LastName	string		`json:"last_name"`
	Number 		int64		`json:"number"`
	Balance 	int64 		`json:"balance"`
	CreateAt  	time.Time 	`json:"createedAt"`
	
}

func NewAccount (firstName, lastName string) *Account {
	return &Account{
		FirstName:	firstName,
		LastName:	lastName,
		Number:		int64(rand.Intn(1000000)),
		Balance:	0, // Default balance
		CreateAt:	time.Now().UTC(),
	}
}