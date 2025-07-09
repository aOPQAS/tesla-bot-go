package models

type User struct {
	ID        string  `json:"id" db:"id"`
	Email     string  `json:"email" db:"email"`
	Password  string  `json:"password" db:"password"`
	IsActive  bool    `json:"is_active" db:"is_active"`
	CreatedAt float64 `json:"created_at" db:"created_at"`
	UpdatedAt float64 `json:"updated_at" db:"updated_at"`
}
