package domain

import "errors"

var ErrInvalidUsername = errors.New("invalid Username")

type UsernameValueObject struct {
	value string
}

func NewUsername(value string) (*UsernameValueObject, error) {
	instance := &UsernameValueObject{value: value}
	return instance, nil
}

func LoadUsername(value string) (*UsernameValueObject, error) {
	instance := &UsernameValueObject{value: value}
	return instance, nil
}

func (u *UsernameValueObject) GetValue() string {
	return u.value
}