// Now we will bind the service methods to the grpc server
package handler

import (
	"context"
	pb "user_management/api/proto"
	"user_management/pkg/service"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	return h.service.CreateUser(ctx, req)
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	return h.service.GetUser(ctx, req)
}

func (h *UserHandler) ListUsers(ctx context.Context, req *pb.Empty) (*pb.ListUserResponse, error) {
	return h.service.ListUsers(ctx, req)
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	return h.service.UpdateUser(ctx, req)
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.Empty, error) {
	return h.service.DeleteUser(ctx, req)
}
