package postgres

import (
	"context"
	"log"

	game "github.com/dilshodforever/4-oyimtixon-game-service/genprotos/product"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// OrderItemService implementation

func (g *GameStorage) CreateOrderItem(ctx context.Context, req *game.OrderItem) (*game.OrderItem, error) {
	coll := g.db.Collection("order_items")
	_, err := coll.InsertOne(ctx, bson.M{
		"id":        req.Id,
		"order_id":  req.OrderId,
		"product_id": req.ProductId,
		"quantity":  req.Quantity,
		"price":     req.Price,
	})
	if err != nil {
		log.Printf("Failed to create order item: %v", err)
		return nil, err
	}

	return req, nil
}

func (g *GameStorage) GetOrderItem(ctx context.Context, req *game.OrderItemIdRequest) (*game.OrderItem, error) {
	coll := g.db.Collection("order_items")
	var orderItem game.OrderItem
	err := coll.FindOne(ctx, bson.M{"id": req.Id}).Decode(&orderItem)
	if err != nil {
		log.Printf("Failed to get order item: %v", err)
		return nil, err
	}

	return &orderItem, nil
}

func (g *GameStorage) UpdateOrderItem(ctx context.Context, req *game.OrderItem) (*game.OrderItem, error) {
	coll := g.db.Collection("order_items")
	_, err := coll.UpdateOne(ctx, bson.M{"id": req.Id}, bson.M{
		"$set": bson.M{
			"order_id":   req.OrderId,
			"product_id": req.ProductId,
			"quantity":   req.Quantity,
			"price":      req.Price,
		},
	})
	if err != nil {
		log.Printf("Failed to update order item: %v", err)
		return nil, err
	}

	return req, nil
}

func (g *GameStorage) DeleteOrderItem(ctx context.Context, req *game.OrderItemIdRequest) (*game.Empty, error) {
	coll := g.db.Collection("order_items")
	_, err := coll.DeleteOne(ctx, bson.M{"id": req.Id})
	if err != nil {
		log.Printf("Failed to delete order item: %v", err)
		return nil, err
	}

	return &game.Empty{}, nil
}

func (g *GameStorage) ListOrderItems(ctx context.Context, req *game.ListRequest) (*game.OrderItemListResponse, error) {
	coll := g.db.Collection("order_items")
	cursor, err := coll.Find(ctx, bson.M{}, &mongo.FindOptions{
		Limit: int64(req.Limit),
		Skip:  int64(req.Offset),
	})
	if err != nil {
		log.Printf("Failed to list order items: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var orderItems []*game.OrderItem
	for cursor.Next(ctx) {
		var orderItem game.OrderItem
		if err := cursor.Decode(&orderItem); err != nil {
			log.Printf("Failed to decode order item: %v", err)
			return nil, err
		}
		orderItems = append(orderItems, &orderItem)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
		return nil, err
	}

	return &game.OrderItemListResponse{OrderItems: orderItems}, nil
}
