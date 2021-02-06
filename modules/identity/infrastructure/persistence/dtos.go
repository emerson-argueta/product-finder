package persistence

import (
	"emerson-argueta/m/v2/modules/identity/domain/user"
)

// UserDTO for scanning from database
type UserDTO struct {
	ID           *string `db:"id" json:"id"`
	Email        *string `db:"email" json:"email"`
	HashPassword *string `db:"hashpassword" json:"hashpassword"`
}

// UserPersistenceToDomain from persistence
func UserPersistenceToDomain(dto UserDTO) (res user.User, e error) {

	email, e := user.NewEmail(dto.Email)
	if e != nil {
		return nil, e
	}
	hashPassword, e := user.ToHashPassword(dto.HashPassword)
	if e != nil {
		return nil, e
	}
	userFields := &user.Fields{
		Email:        &email,
		HashPassword: &hashPassword,
	}
	return user.New(userFields, dto.ID)
}

// UserDomainToPersistence from domain
func UserDomainToPersistence(u user.User) *UserDTO {
	id := u.GetID()
	email := u.GetEmail().ToString()
	hashPassword := u.GetHashPassword().ToString()

	return &UserDTO{
		ID:           &id,
		Email:        &email,
		HashPassword: &hashPassword,
	}

}
