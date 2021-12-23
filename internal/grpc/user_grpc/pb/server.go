package pb

import (
	"google.golang.org/grpc"
)

type GrpcServer struct {
	Server *grpc.Server
}

func NewGrpcServer(opts ...grpc.ServerOption) *GrpcServer {
	return &GrpcServer{
		Server: grpc.NewServer(opts...),
	}
}

func (gs *GrpcServer) Register() {
	RegisterUserServer(gs.Server, NewUserGrpc())
}
