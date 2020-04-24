package user

import (
	"context"
	"google.golang.org/grpc"
	"testing"
)

func TestNewUserServiceClient(t *testing.T) {
	clientConn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Error("Error in connect to server. Error: ", err.Error())
		return
	}
	var client = NewUserServiceClient(clientConn)
	var ctx = context.Background()
	user, err := client.FindAll(ctx, &ListUserRequest{})
	if err != nil {
		t.Error("Error in execute findAll. Error: ", err.Error())
		return
	}
	t.Log(user)
}