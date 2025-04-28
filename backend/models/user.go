package models

// User represents the user database model
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Age      int    `json:"age"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

// LoginRequest represents login payload
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
