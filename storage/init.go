package storage

import (
	pb "github.com/dilshodforever/fooddalivary-food/genprotos"
)

// InitRoot provides access to the different storage interfaces.
type InitRoot interface {
	Product() ProductStorage
}

// ProductStorage defines the methods required for product-related operations.
type ProductStorage interface {
	CreateProduct(req *pb.CreateProductRequest) (*pb.ProductResponse, error)
	GetProduct(req *pb.ProductIdRequest) (*pb.GetProductResponse, error)
	UpdateProduct(req *pb.UpdateProductRequest) (*pb.ProductResponse, error)
	DeleteProduct(req *pb.ProductIdRequest) (*pb.ProductResponse, error)
	ListProducts(req *pb.GetAllProductRequest) (*pb.GetAllProductResponse, error)
}
