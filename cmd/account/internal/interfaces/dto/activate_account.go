package dto

// ActivateAccountRequestDto - activate account request dto.
type ActivateAccountRequestDto struct {
	ID string `json:"id" validate:"required,uuid"`
}

// ActivateAccountResponseDto - activate account response dto.
type ActivateAccountResponseDto struct {
	ID string `json:"id"`
}
