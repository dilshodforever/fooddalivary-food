package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/dilshodforever/fooddalivary-api/genprotos"
)

// AddItems adds items to an order
// @Summary      Add Items
// @Description  Add items to an existing order
// @Tags         OrderItem
// @Accept       json
// @Produce      json
// @Param        items body pb.AddItemsRequest true "Items data"
// @Success      200 {object} pb.AddItemsResponse "Items added successfully"
// @Failure      400 {string} string "Invalid request body"
// @Failure      500 {string} string "Internal Server Error"
// @Router       /orderitems [post]
func (h *Handler) AddItems(ctx *gin.Context) {
	var req pb.AddItemsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	res, err := h.OrderItems.AddItems(context.Background(), &req)
	if err != nil {
		log.Println("AddItems error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// DeleteItems removes items from an order
// @Summary      Delete Items
// @Description  Remove items from an existing order
// @Tags         OrderItem
// @Accept       json
// @Produce      json
// @Param        items body pb.DeleItemsRequest true "Items deletion data"
// @Success      200 {object} pb.DeleteProductResponse "Items deleted successfully"
// @Failure      400 {string} string "Invalid request body"
// @Failure      500 {string} string "Internal Server Error"
// @Router       /orderitems [delete]
func (h *Handler) DeleteItems(ctx *gin.Context) {
	var req pb.DeleItemsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	res, err := h.OrderItems.DeleteItems(context.Background(), &req)
	if err != nil {
		log.Println("DeleteItems error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// ListOrderItems lists all items in an order by user ID
// @Summary      List Order Items
// @Description  List all items in an order for a given user
// @Tags         OrderItem
// @Accept       json
// @Produce      json
// @Param        userId query string true "User ID"
// @Success      200 {object} pb.GetAllItems "List of order items"
// @Failure      400 {string} string "Missing or invalid user ID"
// @Failure      500 {string} string "Internal Server Error"
// @Router       /orderitems [get]
func (h *Handler) ListOrderItems(ctx *gin.Context) {
	userId := ctx.Query("userId")
	if userId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid user ID"})
		return
	}

	req := pb.GetByUseridItems{UserId: userId}
	res, err := h.OrderItems.ListOrderItems(context.Background(), &req)
	if err != nil {
		log.Println("ListOrderItems error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
