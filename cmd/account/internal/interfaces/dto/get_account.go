package dto

type GetAccountRequestDto struct {
	ID string `json:"id"`
}

type GetAccountResponseDto struct {
	ID     string            `json:"id"`
	Email  string            `json:"email"`
	Active bool              `json:"active"`
	Person PersonResponseDto `json:"person"`
}
