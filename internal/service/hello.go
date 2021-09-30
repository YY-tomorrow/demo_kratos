package service

import (
	"context"
	v1 "demo_kratos/api/helloworld/v1"

	pb "demo_kratos/api/helloworld"
)

type HelloService struct {
	pb.UnimplementedHelloServer
}

func NewHelloService() *HelloService {
	return &HelloService{}
}

func (s *HelloService) CreateHello(ctx context.Context, req *pb.CreateHelloRequest) (*pb.CreateHelloReply, error) {
	return &pb.CreateHelloReply{}, nil
}
func (s *HelloService) UpdateHello(ctx context.Context, req *pb.UpdateHelloRequest) (*pb.UpdateHelloReply, error) {
	return &pb.UpdateHelloReply{}, nil
}
func (s *HelloService) DeleteHello(ctx context.Context, req *pb.DeleteHelloRequest) (*pb.DeleteHelloReply, error) {
	return &pb.DeleteHelloReply{}, nil
}
func (s *HelloService) GetHello(ctx context.Context, req *pb.GetHelloRequest) (*pb.GetHelloReply, error) {
	return &pb.GetHelloReply{}, nil
}
func (s *HelloService) ListHello(ctx context.Context, req *pb.ListHelloRequest) (*pb.ListHelloReply, error) {
	return &pb.ListHelloReply{}, nil
}
func (s *HelloService) HelloDemo(ctx context.Context, req *pb.HelloDemoRequest) (*pb.HelloDemoReply, error) {
	if req.GetName() == "error" {
		return nil, v1.ErrorUserNotFound("user not found: %s", req.GetName())
	}
	return &pb.HelloDemoReply{Message: "Hello " + req.GetName()}, nil
}
