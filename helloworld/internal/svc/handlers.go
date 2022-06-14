package svc

import (
	"context"
	"fmt"
	pb "github.com/anlsergio/grpc-demo/helloworld/pkg/pb/greeting/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GreeterService holds all handlers regarding
// the gRPC service definition
type GreeterService struct {
}

// Greet returns a full greeting based upon the name and greeting kind
// parsed from the Message being passed in
func (gs GreeterService) Greet(ctx context.Context, req *pb.GreetRequest) (res *pb.GreetResponse, err error) {
	if req.Msg == nil {
		err = status.New(codes.InvalidArgument, "Message can not be empty").Err()
		return
	}

	helloMsg := fmt.Sprintf("%s, %s", req.Msg.Greeting.String(), req.Msg.Name)

	res = &pb.GreetResponse{
		Resp: helloMsg,
	}

	return
}
