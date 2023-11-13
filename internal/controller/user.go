package controller

import (
	"context"

	proto "github.com/fnunezzz/grpc-demo-go/internal/proto/gen"
	"github.com/fnunezzz/grpc-demo-go/internal/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type loginProps struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}
type createUserProps struct {
	Username string `json:"userName"`
	Password string `json:"password"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Role string `json:"role"`
}

type userServerContract struct {
	login usecase.LoginUserUseCase
	create usecase.CreateUserUseCase
	proto.UnimplementedUserServer
}


func NewUserController(server *grpc.Server, login usecase.LoginUserUseCase, createUserUseCase usecase.CreateUserUseCase) {
	contract := &userServerContract{login: login, create: createUserUseCase}
	proto.RegisterUserServer(server, contract)
}

func (s userServerContract) SignIn(ctx context.Context, req *proto.SignInRequest) (*proto.SignInResponse, error) {
	// var props *loginProps
	
	dto, err := s.login.Handle(req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.SignInResponse{
		Id: dto.ID,
		Token: dto.Token,
		RefreshToken: dto.Token,
	}, nil
}
func (s userServerContract) SignUp(ctx context.Context, req *proto.SignUpRequest) (*proto.SignUpResponse, error) {
	err := s.create.Handle(req.Username, 
	req.Password,
	req.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.SignUpResponse{
		Success: true,
	}, nil
}