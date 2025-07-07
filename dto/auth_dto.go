package dto

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginResponse struct {
    ID        string `json:"id"`
    Email     string `json:"email"`
    FullName  string `json:"full_name"`
    AvatarURL string `json:"avatar_url"`
}

type SignUpRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type SignUpResponse struct {
    ID        string `json:"id"`
    Email     string `json:"email"`
    FullName  string `json:"full_name"`
    AvatarURL string `json:"avatar_url"`
}