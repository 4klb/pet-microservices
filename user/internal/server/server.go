package server

import (
	"context"

	pb "github.com/4klb/pet-microservices/user/proto/user"
)

//SignInServer ..
type SignInServer struct {
	pb.UnimplementedSignInServiceServer
}

//InsertUser ..
func (s *SignInServer) InsertUser(ctx context.Context, user *pb.User) (*pb.StatusResp, error) {
	//TODO
	return &pb.StatusResp{
		Code:        0,
		Description: "User was added",
	}, nil
}

// func (s *Server) Run() error {
// 	return nil
// }
