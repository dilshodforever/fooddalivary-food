package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/dilshodforever/fooddalivary-api/genprotos"
)

// CreateProduct handles the creation of a new product
// @Summary      Create Product
// @Description  Create a new product with the given details
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        product body pb.CreateProductRequest true "Product data"
// @Success      200 {object} pb.ProductResponse "Product created successfully"
// @Failure      400 {string} string "Invalid request body"
// @Failure      500 {string} string "Internal Server Error"
// @Router       /products [post]
func (h *Handler) CreateProduct(ctx *gin.Context) {
	var req pb.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	res, err := h.Product.CreateProduct(context.Background(), &req)
	if err != nil {
		log.Println("CreateProduct error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// GetProduct retrieves a product by its ID
// @Summary      Get Product By ID
// @Description  Retrieve a product's details by its ID
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "Product ID"
// @Success      200 {object} pb.GetProductResponse "Product details"
// @Failure      500 {string} string "Internal Server Error"
// @Router       /products/{id} [get]
func (h *Handler) GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	req := &pb.ProductIdRequest{ProductId: id}
	res, err := h.Product.GetProduct(context.Background(), req)
	if err != nil {
		log.Println("GetProduct error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// UpdateProduct updates an existing product
// @Summary      Update Product
// @Description  Update the details of an existing product
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        product body pb.UpdateProductRequest true "Updated product data"
// @Success      200 {object} pb.ProductResponse "Product updated successfully"
// @Failure      400 {string} string "Invalid request body"
// @Failure      500 {string} string "Internal Server Error"
// @Router       /products [put]
func (h *Handler) UpdateProduct(ctx *gin.Context) {
	var req pb.UpdateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	res, err := h.Product.UpdateProduct(context.Background(), &req)
	if err != nil {
		log.Println("UpdateProduct error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// DeleteProduct removes a product by its ID
// @Summary      Delete Product
// @Description  Delete a product by its ID
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "Product ID"
// @Success      200 {object} pb.ProductResponse "Product deleted successfully"
// @Failure      500 {string} string "Internal Server Error"
// @Router       /products/{id} [delete]
func (h *Handler) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	req := &pb.ProductIdRequest{ProductId: id}
	res, err := h.Product.DeleteProduct(context.Background(), req)
	if err != nil {
		log.Println("DeleteProduct error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// ListProducts lists all products
// @Summary      List Products
// @Description  List all available products
// @Tags         Product
// @Accept       json
// @Produce      json
// @Success      200 {object} pb.GetAllProductResponse "List of products"
// @Failure      500 {string} string "Internal Server Error"
// @Router       /products [get]
func (h *Handler) ListProducts(ctx *gin.Context) {
	req := &pb.GetAllProductRequest{}

	res, err := h.Product.ListProducts(context.Background(), req)
	if err != nil {
		log.Println("ListProducts error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
