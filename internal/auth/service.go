package auth

import (
	"github.com/astlaure/orchid-cms/internal/users"
	"golang.org/x/crypto/bcrypt"
)

func validateUser(body Login) error {
	user, err := users.FindByEmail(body.Email)

	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	return err
}
