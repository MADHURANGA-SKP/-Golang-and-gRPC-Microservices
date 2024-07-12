package handler

import (
	"kitchen/services/common/genproto/orders"
	"kitchen/services/common/genproto/orders/common/util"
	"kitchen/services/orders/types"
	"net/http"
)

type OrderHttphandler struct {
	orderService types.OrderService
}

func NewHttpOrderHandler(orderService types.OrderService) *OrderHttphandler {
	handler := &OrderHttphandler{
		orderService: orderService,
	}

	return handler
}

func (h *OrderHttphandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func ( h *OrderHttphandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := &orders.Order{
		OrderID: 42,
		CustomerID: req.CustomerID,
		ProductID: req.ProductID,
		Quantity: req.Quantity,
	}

	err = h.orderService.CreateOrder(r.Context(), order)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}
	util.WriteJSON(w,http.StatusOK, res)
}