package v1

// DeactivateAccountRequest - deactivate account request dto.
type DeactivateAccountRequest struct {
	ID string `json:"id" validate:"required,uuid"`
}

// DeactivateAccountResponse - deactivate account response dto.
type DeactivateAccountResponse struct {
	ID string `json:"id"`
}
