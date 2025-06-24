package dto

type AuthRequest struct {
	Name     string `json:"name,omitempty"`
	Role     string `json:"role,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
