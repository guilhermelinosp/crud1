package requests

type UserRequest struct {
	Name     string `json:"username" binding:"required,min=3"` // Username must be at least 3 characters long
	Email    string `json:"email" binding:"required,email"`    // Email must be a valid email address
	Password string `json:"password" binding:"required,min=8"` // Password must be at least 8 characters long and match the specified pattern
}
