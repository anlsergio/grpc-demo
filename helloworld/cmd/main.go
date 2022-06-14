package cmd

import (
	"github.com/anlsergio/grpc-demo/helloworld/internal/svc"
	greetingv1 "github.com/anlsergio/grpc-demo/helloworld/pkg/pb/greeting/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	listener net.Listener
	server   *grpc.Server
	logger   *zap.Logger
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	initListener()

	server = grpc.NewServer()

	greetingv1.RegisterGreeterServiceServer(server, &svc.GreeterService{})
	logger.Info("The service handlers have been registered")

	go signalsListener(server)

	logger.Info("Starting the gRPC server...")
	if err := server.Serve(listener); err != nil {
		logger.Panic("Failed to start the gRPC server", zap.Error(err))
	}
}

func initListener() {
	var err error
	addr := "localhost:50051"

	listener, err = net.Listen("tcp", addr)
	if err != nil {
		logger.Panic("Failed to listen",
			zap.String("address", addr),
			zap.Error(err),
		)
	}

	logger.Info("Started listening...", zap.String("address", addr))
}

func signalsListener(s *grpc.Server) {
	waitc := make(chan os.Signal, 1)
	signal.Notify(waitc, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-waitc

	logger.Info("Gracefully shutting down server...")
	server.GracefulStop()
}
