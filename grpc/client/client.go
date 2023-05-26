package client

import (
	"bicycle/bicycle_api_gateway/config"
	"bicycle/bicycle_api_gateway/genproto/order_service"
	"bicycle/bicycle_api_gateway/genproto/sms_service"
	"bicycle/bicycle_api_gateway/genproto/user_service"

	"google.golang.org/grpc"
)

type ServiceManagerI interface {
	OrderService() order_service.OrderServiceClient
	ProductService() order_service.ProductServiceClient
	UserService() user_service.UserServiceClient
	SmsService() sms_service.SmsServiceClient
}

type grpcClients struct {
	orderService   order_service.OrderServiceClient
	productService order_service.ProductServiceClient
	userService    user_service.UserServiceClient
	smsService     sms_service.SmsServiceClient
}

func NewGrpcClient(cfg config.Config) (ServiceManagerI, error) {
	connOrderService, err := grpc.Dial(
		cfg.OrderServiceHost+cfg.OrderServicePort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	connUserService, err := grpc.Dial(
		cfg.UserServiceHost+cfg.UserServicePort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	connSmsService, err := grpc.Dial(
		cfg.SmsServiceHost+cfg.SmsServicePort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		orderService: order_service.NewOrderServiceClient(connOrderService),
		userService:  user_service.NewUserServiceClient(connUserService),
		smsService:   sms_service.NewSmsServiceClient(connSmsService),
	}, nil
}

func (g *grpcClients) OrderService() order_service.OrderServiceClient {
	return g.orderService
}

func (g *grpcClients) UserService() user_service.UserServiceClient {
	return g.userService
}

func (g *grpcClients) SmsService() sms_service.SmsServiceClient {
	return g.smsService
}

func (g *grpcClients) ProductService() order_service.ProductServiceClient {
	return g.productService
}
