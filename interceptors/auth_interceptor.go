package interceptors

import (
	"context"

	"github.com/go-kit/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthInterceptor struct {
	logger         log.Logger
	accessiblePage map[string][]string
}

func NewAuthInterceptor(logger log.Logger, accessiblePage map[string][]string) *AuthInterceptor {
	return &AuthInterceptor{
		logger:         logger,
		accessiblePage: accessiblePage,
	}
}

func (interceptor *AuthInterceptor) UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		interceptor.logger.Log("Auth Interceptor --------->", info.FullMethod)
		err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) error {
	accessibleRoles, ok := interceptor.accessiblePage[method]
	if !ok {
		// everyone can access
		return nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	for _, role := range accessibleRoles {
		interceptor.logger.Log(role)
	}

	return status.Error(codes.PermissionDenied, "no permission to access this RPC")
}
