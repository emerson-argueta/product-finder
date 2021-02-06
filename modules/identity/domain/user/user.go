package user

import (
	"emerson-argueta/m/v2/shared/domain"

	"github.com/google/uuid"
)

// Fields for user
type Fields struct {
	Email        *Email
	HashPassword *HashPassword
}

// User is part of the identity subdomain to support core domains which need it
type User interface {
	GetEmail() Email
	SetEmail(Email)
	GetHashPassword() HashPassword
	SetHashPassword(HashPassword)
	GetID() string
}

//User model
type user struct {
	ID           string
	Email        Email
	HashPassword HashPassword
	aggregate    *domain.AbstractAggregateRoot
}

// New user with role, email and password
func New(userFields *Fields, id *string) (res User, e error) {
	if userFields.Email == nil || userFields.HashPassword == nil {
		return nil, ErrUserIncompleteDetails
	}

	user := &user{
		Email:        *userFields.Email,
		HashPassword: *userFields.HashPassword,
	}

	isNewUser := (id == nil)
	if isNewUser {
		user.ID = uuid.New().String()
	}
	if !isNewUser {
		user.ID = *id
	}

	user.aggregate = &domain.AbstractAggregateRoot{}
	user.aggregate.DomainEvents = make([]domain.Event, 0)
	user.aggregate.Name = "User"
	user.aggregate.ID = user.ID

	if isNewUser {
		user.aggregate.AddDomainEvent(NewUserCreated(user))
	}

	return user, nil
}

func (u *user) GetEmail() Email {
	return u.Email
}
func (u *user) GetHashPassword() HashPassword {
	return u.HashPassword
}
func (u *user) GetID() string {
	return u.ID
}
func (u *user) SetEmail(email Email) {
	u.Email = email
}

func (u *user) SetHashPassword(hashPassword HashPassword) {
	u.HashPassword = hashPassword
}
