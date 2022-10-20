package endpoints

import (
	"context"

	"github.com/ahmadstarving/grpcdemo/business"
	"github.com/ahmadstarving/grpcdemo/service"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	LoginUser  endpoint.Endpoint
}

func MakeEndpoints(s service.ServiceApi) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		LoginUser:  makeLoginUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s service.ServiceApi) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(business.CreateUserRequest)
		result, _ := s.CreateUser(ctx, req.Username, req.Fullname, req.Email, req.Password)
		return result, nil
	}
}

func makeLoginUserEndpoint(s service.ServiceApi) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(business.LoginUserRequest)
		result, _ := s.LoginUser(ctx, req.Username, req.Password)
		return result, nil
	}
}
