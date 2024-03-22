package users

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"`
	Role      UserRole  `json:"role"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type CreateUser struct {
	Name                 string `validate:"required"`
	Email                string `validate:"required,email"`
	Password             string
	PasswordConfirmation string
	Role                 UserRole
}

type UpdateUser struct {
	Name  string
	Email string
	Role  UserRole
}

type UpdateUserPassword struct {
	Password             string
	PasswordConfirmation string
}
