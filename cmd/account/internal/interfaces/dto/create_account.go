package dto

import "time"

// CreateAccountRequestDto - account creation request dto.
type CreateAccountRequestDto struct {
	Email       string    `json:"email" validate:"required,email"`
	Password    string    `json:"password" validate:"min=10,required"`
	Active      bool      `json:"active" validate:"boolean"`
	FirstName   string    `json:"firstName" validate:"required"`
	LastName    string    `json:"lastName" validate:"required"`
	Phone       string    `json:"phone" validate:"required,numeric"`
	DateOfBirth time.Time `json:"dateOfBirth" validate:"required"`
}

// CreateAccountResponseDto - account creation response dto.
type CreateAccountResponseDto struct {
	ID string `json:"id"`
}
