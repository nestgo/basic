package server

import (
	"context"

	"../proto"
	"google.golang.org/grpc"
)

var _ context.Context

//GreeterServer server struct
type GreeterServer struct {
}

//Register register service
func (greeter *GreeterServer) Register(s *grpc.Server) {
	proto.RegisterGreeterServer(s, greeter)
}

//SayHello call SayHello
func (greeter *GreeterServer) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
    //TODO 
    return nil,nil
}
