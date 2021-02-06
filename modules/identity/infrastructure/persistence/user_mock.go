package persistence

import (
	"emerson-argueta/m/v2/modules/identity/domain/user"
	"emerson-argueta/m/v2/modules/identity/repository"
)

var _ repository.UserRepo = &UserMock{}

// UserMock implements a mock user repo for testing
type UserMock struct {
}

// CreateUser returns ErrUserExists if user passed is not nil
func (s *UserMock) CreateUser(u user.User) (e error) {
	if u != nil {
		return user.ErrUserExists
	}

	return nil

}

// RetrieveUserByID If id string is empty,
// returns ErrUserNotFound.
func (s *UserMock) RetrieveUserByID(id string) (res user.User, e error) {
	if len(id) == 0 {
		return nil, user.ErrUserNotFound
	}
	emailStr := "test@test.com"
	email, _ := user.NewEmail(&emailStr)
	passwordStr := "password"
	password, _ := user.NewHashPassword(&passwordStr)

	userFields := user.Fields{
		Email:        &email,
		HashPassword: &password,
	}

	return user.New(&userFields, &id)

}

// RetrieveUserByEmail If the user does not exists,
// returns ErrUserNotFound.
func (s *UserMock) RetrieveUserByEmail(email user.Email) (res user.User, e error) {
	id := "f7599588-b899-4237-9f49-65f7cc14bdf4"

	if email == nil || email.ToString() != "test@test.com" {
		return nil, user.ErrUserNotFound
	}
	passwordStr := "password"
	password, _ := user.NewHashPassword(&passwordStr)

	userFields := user.Fields{
		Email:        &email,
		HashPassword: &password,
	}

	return user.New(&userFields, &id)

}

// UpdateUser If the user is nil, returns
// ErrUserNotFound.
func (s *UserMock) UpdateUser(u user.User) (e error) {

	if u == nil {
		return user.ErrUserNotFound
	}
	return nil
}

// DeleteUser if uuid is an empty string, returns
// ErrUserNotFound.
func (s *UserMock) DeleteUser(uuid string) (e error) {
	if len(uuid) == 0 {
		return user.ErrUserNotFound
	}

	return nil
}
