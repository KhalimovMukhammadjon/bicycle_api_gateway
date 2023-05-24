package client

import (
	"bicycle/api-gateway/config"
	"bicycle/api-gateway/genproto/order_service"
	"bicycle/api-gateway/genproto/user_service"

	"google.golang.org/grpc"
)

type ServiceManagerI interface {
	OrderService() order_service.OrderServiceClient
	ProductService() order_service.ProductServiceClient
	UserService() user_service.UserServiceClient
}

type grpcClients struct {
	orderService   order_service.OrderServiceClient
	productService order_service.ProductServiceClient
	userService    user_service.UserServiceClient
}

func NewGrpcClient(cfg config.Config) (ServiceManagerI, error) {
	connUserService, err := grpc.Dial(
		cfg.OrderServiceHost+cfg.OrderServicePort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		orderService: order_service.NewOrderServiceClient(connUserService),
	}, nil
}

func (g *grpcClients) OrderService() order_service.OrderServiceClient {
	return g.orderService
}

func (g *grpcClients) UserService() user_service.UserServiceClient {
	return g.userService
}

func (g *grpcClients) ProductService() order_service.ProductServiceClient {
	return g.productService
}
