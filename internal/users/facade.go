package users

func FindByEmail(email string) (*User, error) {
	return retrieveUserByEmail(email)
}
