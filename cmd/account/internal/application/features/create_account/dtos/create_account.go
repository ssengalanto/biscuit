package dtos

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

type PersonResponseDto struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

type CreateAccountResponseDto struct {
	ID     string            `json:"id"`
	Email  string            `json:"email"`
	Active bool              `json:"active"`
	Person PersonResponseDto `json:"person"`
}
