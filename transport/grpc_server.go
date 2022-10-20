package transport

import (
	"context"

	"github.com/ahmadstarving/grpcdemo/business"
	"github.com/ahmadstarving/grpcdemo/endpoints"
	"github.com/ahmadstarving/grpcdemo/pb"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type GRPCServer struct {
	pb.UnimplementedGrpcDemoServer
	createUser grpc.Handler
	loginUser  grpc.Handler
}

func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.GrpcDemoServer {
	return &GRPCServer{
		createUser: grpc.NewServer(
			endpoints.CreateUser,
			decodeCreateUser,
			encodeCreateUser,
		),
		loginUser: grpc.NewServer(
			endpoints.LoginUser,
			decodeLoginUser,
			encodeLoginUser,
		),
	}
}

func decodeCreateUser(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateUserRequest)
	return business.CreateUserRequest{
		Username: req.Username,
		Fullname: req.FullName,
		Email:    req.Email,
		Password: req.Password,
	}, nil
}

func encodeCreateUser(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(business.CreateUserResponse)
	return &pb.CreateUserResponse{
		User: &pb.User{Username: resp.User.Username, FullName: resp.User.Fullname, Email: resp.User.Email},
	}, nil
}

func decodeLoginUser(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.LoginUserRequest)
	return business.LoginUserRequest{
		Username: req.Username,
		Password: req.Password,
	}, nil
}

func encodeLoginUser(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(business.LoginUserResponse)
	return &pb.LoginUserResponse{
		User:      &pb.User{Username: resp.User.Username, FullName: resp.User.Fullname},
		SessionId: resp.SessionId,
	}, nil
}

// See: Real implementation of service_grpc.pb.gp
func (s *GRPCServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	_, resp, err := s.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CreateUserResponse), nil
}

// See: Real implementation of service_grpc.pb.gp
func (s *GRPCServer) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	_, resp, err := s.loginUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.LoginUserResponse), nil
}
