package dto

type UpdateRequest struct{
	OldUsername string `json:"old_username"`
	NewUsername string `json:"new_username"`
	Password string `json:"password"`
}