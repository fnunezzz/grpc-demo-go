package repository

import (
	"context"
	"fmt"
	"os"
	"time"

	userDomain "github.com/fnunezzz/grpc-demo-go/internal/domain/user"
	"github.com/google/uuid"
)

type roleModel struct {
	ID int
	Description string
}

type userModel struct {
	ID string
	Username string
	Email string
	Password string
	Created_at time.Time
	Updated_at time.Time
}
type UserRepository interface {
	FindUser(username string) (*userDomain.User, error)
	CreateUser(*userDomain.User) error
	UserExists(username string) (bool, error)
}

type userRepository struct{}

func (u *userRepository) CreateUser(userDomain *userDomain.User) error {
	var model = toModel(userDomain)

	_, err := DB.Exec(context.Background(), `insert into demo_grpc.users 
	(id, username, email, password, created_at, updated_at)
	values
	($1, $2, $3, $4, $5, $6)`,model.ID, model.Username, model.Email, model.Password, model.Created_at, model.Updated_at)

	return err
}

func (u *userRepository) UserExists(username string) (bool, error) {
	var exists bool
	err := DB.QueryRow(context.Background(), "select exists(select id from demo_grpc.users where username = $1)", username).Scan(&exists)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return false, err
	}
	return exists, nil
}

func (u *userRepository) FindUser(username string) (*userDomain.User, error) {
	var model userModel
	err := DB.QueryRow(context.Background(), "select u.id, u.username, u.email, u.password from demo_grpc.users u where u.username = $1", username).Scan(
			&model.ID,
			&model.Username,
			&model.Email,
			&model.Password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}
	
	return toDomain(model)
}





func toModel(domain *userDomain.User) *userModel {
	return &userModel{
		ID: domain.GetID(),
		Username: domain.GetUsername(),
		Email: domain.GetEmail().GetValue(),
		Password: domain.GetPassword().String(),
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}
}

func toDomain(model userModel) (*userDomain.User, error) {
	passwordObject, err := userDomain.LoadPassword(model.Password)

	if err != nil {
		return nil, err
	}


	usernameObject, err := userDomain.LoadUsername(model.Username)

	if err != nil {
		return nil, err
	}


	emailObject, err := userDomain.LoadEmail(model.Email)

	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(model.ID)

	if err != nil {
		return nil, err
	}

	user, err := userDomain.LoadUser(&id, usernameObject, emailObject, passwordObject)

	if err != nil {
		return nil, err
	}

	return user, nil
}