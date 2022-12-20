package dto

import "time"

type GetAccountRequestDto struct {
	ID string `json:"id"`
}

type PersonResponseDto struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

type GetAccountResponseDto struct {
	ID     string            `json:"id"`
	Email  string            `json:"email"`
	Active bool              `json:"active"`
	Person PersonResponseDto `json:"person"`
}
