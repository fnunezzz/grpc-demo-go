package domain

import (
	"testing"
)

func TestCreateNewUser(t *testing.T) {
	passwordObject, err := NewPassword("123456789")

	if err != nil {
		t.Error("Error creating password")
	}

	usernameObject, err := NewUsername("username")

	if err != nil {
		t.Error("Error creating username")
	}


	emailObject, err := NewEmail("email@email.com")

	if err != nil {
		t.Error("Error creating email")
	}

	_, err = NewUser(usernameObject, emailObject, passwordObject)

	if err != nil {
		t.Error("Error creating user")
	}

	t.Log("User created with success")

}


func TestValidateCorrectCredentials(t *testing.T) {
	passwordObject, err := NewPassword("123456789")

	if err != nil {
		t.Error("Error creating password")
	}

	usernameObject, err := NewUsername("username")

	if err != nil {
		t.Error("Error creating username")
	}


	emailObject, err := NewEmail("email@email.com")

	if err != nil {
		t.Error("Error creating email")
	}

	user, err := NewUser(usernameObject, emailObject, passwordObject)

	if err != nil {
		t.Error("Error creating user")
	}

	err = user.ValidateCredentials("123456789")

	if err != nil {
		t.Error("Error validating credentials > ", err.Error())
	}

	t.Log("Credentials validated")

}

func TestValidateWrongCredentials(t *testing.T) {
	passwordObject, err := NewPassword("123456789")

	if err != nil {
		t.Error("Error creating password")
	}

	usernameObject, err := NewUsername("username")

	if err != nil {
		t.Error("Error creating username")
	}


	emailObject, err := NewEmail("email@email.com")

	if err != nil {
		t.Error("Error creating email")
	}

	user, err := NewUser(usernameObject, emailObject, passwordObject)

	if err != nil {
		t.Error("Error creating user")
	}

	err = user.ValidateCredentials("123")

	if err == nil {
		t.Error("Error validating credentials > ", err.Error())
	}
	
	t.Log("Credentials validated")


}