package service

import (
	"context"

	"github.com/zakiry/golang_api_sample/database"
	pb "github.com/zakiry/golang_api_sample/proto"
)

func NewUserService(user database.User) pb.UserService {
	return &userService{user: user}
}

type userService struct {
	user database.User
}

func (svc *userService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user := svc.user.GetById(req.GetId())
	return &pb.GetUserResponse{Id: user.Id, Name: user.Name, Age: user.Age}, nil
}
