package service

import (
	"context"
	"kitchen/services/common/genproto/orders"
)

var orderDB = make([]*orders.Order, 0)

type OrderService struct {
	//
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func(s *OrderService) CreateOrder(ctx context.Context, res *orders.Order) error {
	orderDB = append(orderDB, res)
	return nil
}

func(s *OrderService) GetOrder(ctx context.Context) []*orders.Order  {
	return orderDB
}