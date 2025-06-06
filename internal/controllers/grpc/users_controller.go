package grpc

import (
	"context"
	"errors"

	desc "github.com/GP-Hacks/proto/pkg/api/user"
	"github.com/GP-Hacks/users/internal/models"
	"github.com/GP-Hacks/users/internal/services"
	"github.com/GP-Hacks/users/internal/services/user_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserController struct {
	desc.UnimplementedUserServiceServer
	usersService *user_service.UserService
}

func NewUserController(us *user_service.UserService) *UserController {
	return &UserController{
		usersService: us,
	}
}

func (c *UserController) GetMe(ctx context.Context, req *desc.GetMeRequest) (*desc.GetMeResponse, error) {
	u, err := c.usersService.GetMe(ctx, req.Token)
	if err != nil {
		return nil, err
	}

	var st desc.UserStatus
	if u.Status == models.DefaultUser {
		st = desc.UserStatus_DEFAULT
	} else if u.Status == models.AdminUser {
		st = desc.UserStatus_ADMIN
	} else {
		return nil, status.Error(codes.Internal, "not valid user status")
	}

	res := desc.GetMeResponse{
		User: &desc.User{
			Email:       u.Email,
			FirstName:   u.FirstName,
			LastName:    u.LastName,
			Surname:     u.Surname,
			DateOfBirth: timestamppb.New(u.DateOfBirth),
		},
		AvatarURL: u.AvatarURL,
		Status:    st,
	}

	return &res, nil
}

func (c *UserController) Update(ctx context.Context, req *desc.UpdateUserRequest) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Internal, "Not implemented")
}

func (c *UserController) Create(ctx context.Context, req *desc.CreateRequest) (*emptypb.Empty, error) {
	if err := c.usersService.CreateUser(ctx, &models.User{
		ID:          req.Id,
		Email:       req.User.Email,
		FirstName:   req.User.FirstName,
		LastName:    req.User.LastName,
		Surname:     req.User.Surname,
		DateOfBirth: req.User.DateOfBirth.AsTime(),
		Status:      models.DefaultUser,
	}); err != nil {
		if errors.Is(err, services.AlreadyExists) {
			return &emptypb.Empty{}, status.Error(codes.AlreadyExists, "users with data already exists")
		} else {
			return &emptypb.Empty{}, status.Error(codes.Internal, "internal error")
		}
	}

	return &emptypb.Empty{}, nil
}

func (c *UserController) UploadAvatar(ctx context.Context, req *desc.UploadAvatarRequest) (*desc.UploadAvatarResponse, error) {
	url, err := c.usersService.UpdateAvatar(ctx, req.Token, req.Photo)
	if err != nil {
		return nil, err
	}

	return &desc.UploadAvatarResponse{
		Url: url,
	}, nil
}
