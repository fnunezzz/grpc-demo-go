package domain

import (
	"errors"

	"github.com/google/uuid"
)

var ErrInvalidUserID = errors.New("invalid User id")
var ErrInvalidCredentials = errors.New("Informed credentials are invalid")

type User struct {
	id uuid.UUID
	username UsernameValueObject
	email EmailValueObject
	password PasswordValueObject
}

func NewUser(userName *UsernameValueObject, email *EmailValueObject, password *PasswordValueObject) (*User, error) {	
	id := uuid.New()
	if userName == nil {
		return nil, ErrInvalidUsername
	}

	if email == nil {
		return nil, ErrInvalidEmail
	}

	if password == nil {
		return nil, ErrInvalidPassword
	}

	user := &User{id: id, username: *userName, email: *email, password: *password}

	return user, nil
}

func LoadUser(id *uuid.UUID, userName *UsernameValueObject, email *EmailValueObject, password *PasswordValueObject) (*User, error) {	
	user := &User{id: *id, username: *userName, email: *email, password: *password}
	return user, nil
}

func (u *User) GetID() string {
	return u.id.String()
}

func (u *User) GetUsername() string {
	return u.username.value
}

func (u *User) GetEmail() *EmailValueObject {
	return &u.email
}

func (u *User) GetPassword() *PasswordValueObject {
	return &u.password
}

func (u *User) ValidateCredentials(password string) error {
	err := u.password.ComparePassword(password)
	if err != nil {
		return ErrInvalidCredentials
	}
	return nil
}

func (u *User) Login() {

}