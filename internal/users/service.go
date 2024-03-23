package users

import (
	"github.com/astlaure/orchid-cms/internal/core"
	"golang.org/x/crypto/bcrypt"
)

func retrieveUserPage(page uint) (*[]User, error) {
	var users []User
	result := core.DB.Limit(12).Offset(int(page * 12)).Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return &users, nil
}

func retrieveUserById(id uint) (*User, error) {
	var user User
	result := core.DB.First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func retrieveUserByEmail(email string) (*User, error) {
	var user User
	result := core.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func insertUser(createUser CreateUser) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(createUser.Password), 10)

	if err != nil {
		return nil, err
	}

	var user = User{
		Name:     createUser.Name,
		Email:    createUser.Email,
		Password: string(hash),
		Role:     createUser.Role,
	}

	result := core.DB.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func modifyUser(id uint, body UpdateUser) error {
	var user User
	result := core.DB.First(&user, id)

	if result.Error != nil {
		return result.Error
	}

	user.Name = body.Name
	user.Email = body.Email
	user.Role = body.Role

	result = core.DB.Save(&user)
	return result.Error
}

func modifyUserPassword(id uint, body UpdateUserPassword) error {
	var user User
	result := core.DB.First(&user, id)

	if result.Error != nil {
		return result.Error
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		return err
	}

	user.Name = string(hash)

	result = core.DB.Save(&user)
	return result.Error
}

func destroyUser(id uint) error {
	result := core.DB.Delete(&User{}, id)
	return result.Error
}
