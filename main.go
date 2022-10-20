package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ahmadstarving/grpcdemo/endpoints"
	"github.com/ahmadstarving/grpcdemo/interceptors"
	"github.com/ahmadstarving/grpcdemo/pb"
	"github.com/ahmadstarving/grpcdemo/service"
	"github.com/ahmadstarving/grpcdemo/transport"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func accessiblePage() map[string][]string {
	const servicePath = "/pb.GrpcDemo/"

	return map[string][]string{
		servicePath + "CreateUser": {"admin"},
	}
}

func main() {
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	addService := service.NewService(logger)
	addEndpoint := endpoints.MakeEndpoints(addService)
	grpcServer := transport.NewGRPCServer(addEndpoint, logger)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grcpListener, err := net.Listen("tcp", ":8383")
	if err != nil {
		logger.Log("During", "Listen", "Error", err)
		os.Exit(1)
	}

	//define interceptor
	interceptors := interceptors.NewAuthInterceptor(logger, accessiblePage())

	go func() {
		baseServer := grpc.NewServer(
			grpc.UnaryInterceptor(interceptors.UnaryInterceptor()),
		)
		pb.RegisterGrpcDemoServer(baseServer, grpcServer)
		level.Info(logger).Log("msg", "Server started successfully")
		reflection.Register(baseServer)
		baseServer.Serve(grcpListener)
	}()

	level.Error(logger).Log("Exit", <-errs)
}
