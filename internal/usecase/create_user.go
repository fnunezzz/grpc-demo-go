package usecase

import (
	"errors"

	userDomain "github.com/fnunezzz/grpc-demo-go/internal/domain/user"
	"github.com/fnunezzz/grpc-demo-go/internal/repository"
)

type createUserUseCase struct {
	dao repository.DAO
}

type CreateUserUseCase interface {
	Handle(username string,
		password string,
		email string) error
}

func NewCreateUserUseCase(dao repository.DAO) CreateUserUseCase {
	return &createUserUseCase{dao: dao}
}

func (u *createUserUseCase) Handle(
	username string,
	password string,
	email string) error {

	passwordObject, err := userDomain.NewPassword(password)

	if err != nil {
		return err
	}

	usernameObject, err := userDomain.NewUsername(username)

	if err != nil {
		return err
	}


	emailObject, err := userDomain.NewEmail(email)

	if err != nil {
		return err
	}

	user, err := userDomain.NewUser(usernameObject, emailObject, passwordObject)

	if err != nil {
		return err
	}

	exists, err := u.dao.NewUserRepository().UserExists(user.GetUsername())

	if err != nil {
		return err
	}
	
	if exists {
		return errors.New("username already exists")
	}

	err = u.dao.NewUserRepository().CreateUser(user)

	return err
}