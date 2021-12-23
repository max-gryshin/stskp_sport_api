package main

import (
	"context"
	"fmt"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/grpc/user_grpc/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

func timingInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	fmt.Printf(`--
	call=%v
	req=%#v
	reply=%#v
	time=%v
	err=%v
	`, method, req, reply, time.Since(start), err)
	return err
}

type tokenAuth struct {
	Token string
}

func (t *tokenAuth) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"access-token": t.Token,
	}, nil
}

func (c *tokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	grpcConn, err := grpc.Dial(
		"localhost:8081",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(timingInterceptor),
		grpc.WithPerRPCCredentials(&tokenAuth{"123"}),
	)
	if err != nil {
		log.Print(err)
	}
	defer grpcConn.Close()

	userClient := pb.NewUserClient(grpcConn)
	ctx := context.Background()
	response, errResponse := userClient.Get(ctx, &pb.Request{ID: 1})
	if errResponse != nil {
		log.Print(errResponse.Error())
	}

	fmt.Println(response)
}
