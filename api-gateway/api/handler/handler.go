package handler

import (
	pb "github.com/dilshodforever/fooddalivary-api/genprotos"
)


type Handler struct {
	Order      pb.OrderServiceClient
	OrderItems pb.OrderItemServiceClient
	Product    pb.ProductServiceClient
}


func NewHandler(order pb.OrderServiceClient, 
	orderItems pb.OrderItemServiceClient, 
	product pb.ProductServiceClient) *Handler {
	return &Handler{
		Order:      order,
		OrderItems: orderItems,
		Product:    product,
	}
}
