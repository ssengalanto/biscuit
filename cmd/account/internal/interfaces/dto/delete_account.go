package dto

// DeleteAccountRequest - delete account request dto.
type DeleteAccountRequest struct {
	ID string `json:"id" validate:"required,uuid"`
}

// DeleteAccountResponse - delete account response dto.
type DeleteAccountResponse struct {
	ID string `json:"id"`
}
