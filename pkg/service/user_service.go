package service

import (
	"context"
	"errors"
	"sync"
	pb "user_management/api/proto"
)

type UserService struct {
	mu     sync.Mutex
	users  map[int32]*pb.UserResponse
	nextID int32
}

func NewUserService() *UserService {
	return &UserService{
		users:  make(map[int32]*pb.UserResponse),
		nextID: 1,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// Normally check if email exists — skipped for simplicity
	user := &pb.UserResponse{
		Id:    s.nextID,
		Name:  req.Name,
		Email: req.Email,
	}

	s.users[user.Id] = user
	s.nextID++

	return user, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, ok := s.users[req.Id]
	if !ok {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (s *UserService) ListUsers(ctx context.Context, _ *pb.Empty) (*pb.ListUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	users := make([]*pb.UserResponse, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return &pb.ListUserResponse{Users: users}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	user, ok := s.users[req.Id]
	if !ok {
		return nil, errors.New("user not found")
	}

	user.Name = req.Name
	user.Email = req.Email
	s.users[req.Id] = user
	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.users[req.Id]
	if !ok {
		return nil, errors.New("user not found")
	}
	delete(s.users, req.Id)

	return &pb.Empty{}, nil
}
