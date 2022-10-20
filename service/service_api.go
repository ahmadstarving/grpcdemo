package service

import (
	"context"

	"github.com/ahmadstarving/grpcdemo/business"
	"github.com/go-kit/log"
)

type ServiceApi interface {
	CreateUser(ctx context.Context, username string, full_name string, email string, password string) (business.CreateUserResponse, error)
	LoginUser(ctx context.Context, username string, password string) (business.LoginUserResponse, error)
}

type service struct {
	logger log.Logger
}

// CreateUser implements ServiceApi
func (*service) CreateUser(ctx context.Context, username string, full_name string, email string, password string) (business.CreateUserResponse, error) {
	return business.CreateUserResponse{
		User: business.User{Username: username, Fullname: full_name, Email: email, Password: password},
	}, nil

}

// LoginUser implements ServiceApi
func (*service) LoginUser(ctx context.Context, username string, password string) (business.LoginUserResponse, error) {
	return business.LoginUserResponse{
		User:      business.User{Username: username, Password: password},
		SessionId: "Generated Session ID",
	}, nil

}

func NewService(logger log.Logger) ServiceApi {
	return &service{
		logger: logger,
	}
}
