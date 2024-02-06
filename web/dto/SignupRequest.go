package dto

type SignupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"usernameValid"`
	Password string `json:"password" binding:"passwordValid"`
}
