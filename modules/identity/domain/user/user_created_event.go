package user

import (
	"emerson-argueta/m/v2/shared/domain"
	"time"
)

// UserCreatedEventName of domain event
const UserCreatedEventName = "UserCreated"

// Created domain event
type Created interface {
	GetUser() User
	domain.Event
}
type userCreated struct {
	User  User
	event *domain.AbstractEvent
}

// NewUserCreated domain event
func NewUserCreated(user User) Created {
	userCreated := &userCreated{
		User:  user,
		event: &domain.AbstractEvent{},
	}

	userCreated.event.TimeOccured = time.Now()
	userCreated.event.AggregateID = user.GetID()
	userCreated.event.Name = UserCreatedEventName

	return userCreated
}
func (e *userCreated) GetAggregateID() string {
	return e.event.AggregateID
}
func (e *userCreated) GetName() string {
	return e.event.Name
}
func (e *userCreated) GetUser() User {
	return e.User
}
