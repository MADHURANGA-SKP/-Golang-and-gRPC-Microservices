package main

import (
	"kitchen/services/orders/handler"
	"kitchen/services/orders/service"
	"log"
	"net/http"
)

type httpServer struct{
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	orderService := service.NewOrderService()
	orderHandler := handler.NewHttpOrderHandler(orderService)
	orderHandler.RegisterRouter(router)

	log.Println("stating server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}