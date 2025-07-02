package dto

type UserCreateRequest struct {
	GoogleID  string `json:"google_id" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	FullName  string `json:"full_name"`
	AvatarURL string `json:"avatar_url"`
}

type UserUpdateRequest struct {
	FullName  string `json:"full_name"`
	AvatarURL string `json:"avatar_url"`
}