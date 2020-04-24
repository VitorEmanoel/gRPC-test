package services

import (
	"google.golang.org/grpc"
	"grpc-test/services/user"
)

func NewServices(server *grpc.Server) {
	user.ImplementUserService(server)
}
