package dto

// DeleteAccountRequestDto - delete account request dto.
type DeleteAccountRequestDto struct {
	ID string `json:"id" validate:"required,uuid"`
}

// DeleteAccountResponseDto - delete account response dto.
type DeleteAccountResponseDto struct {
	ID string `json:"id"`
}
