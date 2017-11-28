package basic

import (
	"context"
	"fmt"

	"./proto"

	"google.golang.org/grpc"
)

var _ context.Context

//GreeterClient client
type GreeterClient struct {
	conn   *grpc.ClientConn
	client proto.GreeterClient
}

//Init init connection
func (greeterClient *GreeterClient) Init(address string) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("did not connect: %v", err)
	}
	greeterClient.conn = conn
	greeterClient.client = proto.NewGreeterClient(conn)
	return err
}

//Close close connection
func (greeterClient *GreeterClient) Close() error {
	if greeterClient != nil {
		return greeterClient.conn.Close()
	}
	return nil
}

//SayHello call SayHello
func (greeterClient *GreeterClient) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
    return greeterClient.client.SayHello(ctx, in)
}
