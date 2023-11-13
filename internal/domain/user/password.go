package domain

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)


var ErrInvalidPassword = errors.New("invalid Password")
var ErrInvalidPasswordLenght = errors.New("password length is lower than 8")

type PasswordValueObject struct {
	value []byte
}

func NewPassword(value string) (*PasswordValueObject, error) {
	instance := &PasswordValueObject{}
	err := instance.hashPassword(value)
	if err != nil {
		return nil, err
	}

	return instance, nil
}

func LoadPassword(value string) (*PasswordValueObject, error) {
	instance := &PasswordValueObject{value: []byte(value)}
	return instance, nil
}

func (p *PasswordValueObject) String() string {
	return string(p.value)
}

func (p *PasswordValueObject) hashPassword(value string) error {
	if len(value) < 8 {
		return ErrInvalidPasswordLenght
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.value = hash
	return nil
}

func (p *PasswordValueObject) ComparePassword(pass string) error {
    err := bcrypt.CompareHashAndPassword(p.value, []byte(pass))
    if err != nil {
		return ErrInvalidPassword
	}
	return nil
}