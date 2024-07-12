package handler

import (
	"context"
	"kitchen/services/common/genproto/orders"
	"kitchen/services/orders/types"

	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	orderService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrderService(grpc *grpc.Server, orderService types.OrderService) {
	grpcHandler := &OrdersGrpcHandler{
		orderService: orderService,
	}

	//register the orderService
	 orders.RegisterOrderServiceServer(grpc, grpcHandler)
}

func (h *OrdersGrpcHandler) GetOrder(ctx context.Context, req *orders.GetOrderRequest) (*orders.GetOrderResponse, error ){
	o := h.orderService.GetOrder(ctx)
	res := &orders.GetOrderResponse{
		Orders: o,
	}

	return res, nil
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error ){
	order := &orders.Order{
		OrderID: 42,
		CustomerID: 2,
		ProductID: 1,
		Quantity: 10,
	}

	err := h.orderService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	return res, nil
}

