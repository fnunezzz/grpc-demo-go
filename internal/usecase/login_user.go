package usecase

import (
	"errors"
	"os"
	"time"

	"github.com/fnunezzz/grpc-demo-go/internal/repository"
	"github.com/golang-jwt/jwt/v5"
)

type jwtDto struct {
	Token string `json:"token,omitempty"`
	ID string `json:"ID,omitempty"`
	Error string `json:"error,omitempty"`
}

type loginUserUseCase struct {
	dao repository.DAO
}

type LoginUserUseCase interface {
	Handle(userName string, password string) (*jwtDto, error)
}

func NewLoginUserUseCase(dao repository.DAO) LoginUserUseCase {
	return &loginUserUseCase{dao: dao}
}

func (u *loginUserUseCase) Handle(userName string, password string) (*jwtDto, error) {
	user, err := u.dao.NewUserRepository().FindUser(userName)
	if err != nil {
		return nil, errors.New("username or password are invalid")
	}

	err = user.ValidateCredentials(password)
	if err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{
		"ID":  user.GetID(),
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return nil, err
	}
	return &jwtDto{Token: t, ID: user.GetID()}, nil
}
