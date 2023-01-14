package dto

// ActivateAccountRequest - activate account request dto.
type ActivateAccountRequest struct {
	ID string `json:"id" validate:"required,uuid"`
}

// ActivateAccountResponse - activate account response dto.
type ActivateAccountResponse struct {
	ID string `json:"id"`
}
