package domain

import "errors"

var ErrInvalidEmail = errors.New("invalid email")

type EmailValueObject struct {
	value string
}

func NewEmail(value string) (*EmailValueObject, error) {
	instance := &EmailValueObject{value: value}
	return instance, nil
}

func LoadEmail(value string) (*EmailValueObject, error) {
	instance := &EmailValueObject{value: value}
	return instance, nil
}

func (e *EmailValueObject) GetValue() string {
	return e.value
}