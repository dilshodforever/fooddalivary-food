package service

import (
	"context"
	"log"

	pb "github.com/dilshodforever/fooddalivary-food/genprotos"
	s "github.com/dilshodforever/fooddalivary-food/storage"
)

type ProductService struct {
	stg s.InitRoot
	pb.UnimplementedProductServiceServer
}

func NewProductService(stg s.InitRoot) *ProductService {
	return &ProductService{stg: stg}
}

// CreateProduct handles the creation of a new product
func (s *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.ProductResponse, error) {
	resp, err := s.stg.Product().CreateProduct(ctx, req)
	if err != nil {
		log.Print(err)
		return &pb.ProductResponse{Message: "Failed to create product", Success: false}, err
	}
	return &pb.ProductResponse{Message: "Product created successfully", Success: true}, nil
}

// GetProduct retrieves a product by its ID
func (s *ProductService) GetProduct(ctx context.Context, req *pb.ProductIdRequest) (*pb.GetProductResponse, error) {
	resp, err := s.stg.Product().GetProduct(ctx, req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

// UpdateProduct updates an existing product
func (s *ProductService) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.ProductResponse, error) {
	resp, err := s.stg.Product().UpdateProduct(ctx, req)
	if err != nil {
		log.Print(err)
		return &pb.ProductResponse{Message: "Failed to update product", Success: false}, err
	}
	return &pb.ProductResponse{Message: "Product updated successfully", Success: true}, nil
}

// DeleteProduct removes a product by its ID
func (s *ProductService) DeleteProduct(ctx context.Context, req *pb.ProductIdRequest) (*pb.ProductResponse, error) {
	resp, err := s.stg.Product().DeleteProduct(ctx, req)
	if err != nil {
		log.Print(err)
		return &pb.ProductResponse{Message: "Failed to delete product", Success: false}, err
	}
	return &pb.ProductResponse{Message: "Product deleted successfully", Success: true}, nil
}

// ListProducts lists all products
func (s *ProductService) ListProducts(ctx context.Context, req *pb.GetAllProductRequest) (*pb.GetAllProductResponse, error) {
	resp, err := s.stg.Product().ListProducts(ctx, req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}
