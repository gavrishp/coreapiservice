package dto

type AddUserRequestBody struct {
	Name        string `json:"name"`
	Role        string `json:"role"`
	Description string `json:"description"`
}

type UpdateUserRequestBody struct {
	Name        string `json:"name"`
	Role        string `json:"role"`
	Description string `json:"description"`
}
