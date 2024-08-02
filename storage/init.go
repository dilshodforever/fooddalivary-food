package storage

import (
	pb "github.com/dilshodforever/fooddalivary-food/genprotos"
)

// InitRoot provides access to the different storage interfaces.
type InitRoot interface {
	Product() ProductStorage
	Order() OrderStorage
	OrderItem() OrderItemStorage
}

// ProductStorage defines the methods required for product-related operations.
type ProductStorage interface {
	CreateProduct(req *pb.CreateProductRequest) (*pb.ProductResponse, error)
	GetProduct(req *pb.ProductIdRequest) (*pb.GetProductResponse, error)
	UpdateProduct(req *pb.UpdateProductRequest) (*pb.ProductResponse, error)
	DeleteProduct(req *pb.ProductIdRequest) (*pb.ProductResponse, error)
	ListProducts(req *pb.GetAllProductRequest) (*pb.GetAllProductResponse, error)
}

// OrderStorage defines the methods required for order-related operations.
type OrderStorage interface {
	CreateOrder(req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error)
	GetOrderById(req *pb.GetByIdRequest) (*pb.GetAllOrders, error)
	UpdateOrder(req *pb.UpdateOrderRequest) (*pb.UpdateStatusResponse, error)
	DeleteOrder(req *pb.DeleteOrdersByidRequest) (*pb.DeleteOrdersByidResponse, error)
	ListOrders(req *pb.GetAllRequest) (*pb.GetAllOrdersList, error)
	UpdateStatus(req *pb.UpdateStatusRequest) (*pb.UpdateStatusResponse, error)
}

// OrderItemStorage defines the methods required for order item-related operations.
type OrderItemStorage interface {
	AddItems(req *pb.AddItemsRequest) (*pb.AddItemsResponse, error)
	DeleteItems(req *pb.DeleItemsRequest) (*pb.DeleteProductResponse, error)
	ListOrderItems(req *pb.GetByUseridItems) (*pb.GetAllItems, error)
}
