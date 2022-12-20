package dto

import "time"

type CreateAccountRequestDto struct {
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Active      bool      `json:"active"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Phone       string    `json:"phone"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

type CreateAccountResponseDto struct {
	ID string `json:"id"`
}
