package user

import (
	"context"
	"google.golang.org/grpc"
)

type ServiceImpl struct {

}

func (s *ServiceImpl) FindAll(context.Context, *ListUserRequest) (*User, error) {
	return &User{
		Id:       1,
		Username: "Vitor",
	}, nil
}

func NewServiceImpl() UserServiceServer {
	return &ServiceImpl{}
}

func ImplementUserService(server *grpc.Server) {
	RegisterUserServiceServer(server, NewServiceImpl())
}