package domain

import (
	"errors"
	"testing"
)

func TestPasswordMinimalLength(t *testing.T) {
	_, err := NewPassword("123")


	if err != nil && !errors.Is(err, ErrInvalidPasswordLenght) {
		t.Error("Error in password lenght validation")
	}

	if err == nil {
		t.Error("Error in password validation")
	}

}