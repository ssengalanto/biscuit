package dto

// DeactivateAccountRequestDto - deactivate account request dto.
type DeactivateAccountRequestDto struct {
	ID string `json:"id" validate:"required,uuid"`
}

// DeactivateAccountResponseDto - deactivate account response dto.
type DeactivateAccountResponseDto struct {
	ID string `json:"id"`
}
