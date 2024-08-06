package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/dilshodforever/fooddalivary-api/genprotos"
)

// CreateOrder handles the creation of a new order
// @Summary      Create Order
// @Description  Create a new order with the given details
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param        order body pb.CreateOrderRequest true "Order data"
// @Success      200 {object} pb.CreateOrderResponse "Order created successfully"
// @Failure      400 {string} string "Invalid request body"
// @Failure      500 {string} string "Internal Server Error"
// @Router       /orders [post]
func (h *Handler) CreateOrder(ctx *gin.Context) {
	var req pb.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	res, err := h.Order.CreateOrder(context.Background(), &req)
	if err != nil {
		log.Println("CreateOrder error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	

	ctx.JSON(http.StatusOK, res)
}

// GetOrderById retrieves an order by its ID
// @Summary      Get Order By ID
// @Description  Retrieve an order's details by its ID
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "Order ID"
// @Success      200 {object} pb.GetAllOrders "Order details"
// @Failure      500 {string} string "Internal Server Error"
// @Router       /orders/{id} [get]
func (h *Handler) GetOrderById(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.GetByIdRequest{OrderId: id}
	res, err := h.Order.GetOrderByid(context.Background(), req)
	if err != nil {
		log.Println("GetOrderById error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// UpdateOrder updates an existing order
// @Summary      Update Order
// @Description  Update the details of an existing order
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param        order body pb.UpdateOrderRequest true "Updated order data"
// @Success      200 {object} pb.UpdateStatusResponse "Order updated successfully"
// @Failure      400 {string} string "Invalid request body"
// @Failure      500 {string} string "Internal Server Error"
// @Router       /orders [put]
func (h *Handler) UpdateOrder(ctx *gin.Context) {
	var req pb.UpdateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	res, err := h.Order.UpdateOrder(context.Background(), &req)
	if err != nil {
		log.Println("UpdateOrder error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// DeleteOrder removes an order by its ID
// @Summary      Delete Order
// @Description  Delete an order by its ID
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "Order ID"
// @Success      200 {object} pb.DeleteOrdersByidResponse "Order deleted successfully"
// @Failure      500 {string} string "Internal Server Error"
// @Router       /orders/{id} [delete]
func (h *Handler) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	req := &pb.DeleteOrdersByidRequest{OrderId	: id}
	res, err := h.Order.DeleteOrder(context.Background(), req)
	if err != nil {
		log.Println("DeleteOrder error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// ListOrders lists all orders
// @Summary      List Orders
// @Description  List all orders with optional filtering
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param        status query string false "Order status"
// @Success      200 {object} pb.GetAllOrdersList "List of orders"
// @Failure      500 {string} string "Internal Server Error"
// @Router       /orders [get]
func (h *Handler) ListOrders(ctx *gin.Context) {
	req := pb.GetAllRequest{
		// Extracting the status query parameter if present
		Status: ctx.Query("status"),
	}

	res, err := h.Order.ListOrders(context.Background(), &req)
	if err != nil {
		log.Println("ListOrders error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// UpdateStatus updates the status of an order
// @Summary      Update Order Status
// @Description  Update the status of an existing order
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param        status body pb.UpdateStatusRequest true "Status update data"
// @Success      200 {object} pb.UpdateStatusResponse "Order status updated successfully"
// @Failure      400 {string} string "Invalid request body"
// @Failure      500 {string} string "Internal Server Error"
// @Router       /orders/status [patch]
func (h *Handler) UpdateStatus(ctx *gin.Context) {
	var req pb.UpdateStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	res, err := h.Order.UpdateStatus(context.Background(), &req)
	if err != nil {
		log.Println("UpdateStatus error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}






