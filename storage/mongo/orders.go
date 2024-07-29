package postgres

import (
	"context"
	"log"

	"github.com/dilshodforever/4-oyimtixon-game-service/genprotos/game"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// OrderService implementation

func (g *GameStorage) CreateOrder(ctx context.Context, req *game.Order) (*game.Order, error) {
	coll := g.db.Collection("orders")
	_, err := coll.InsertOne(ctx, bson.M{
		"id":               req.Id,
		"user_id":          req.UserId,
		"courier_id":       req.CourierId,
		"status":           req.Status,
		"total_amount":     req.TotalAmount,
		"delivery_address": req.DeliveryAddress,
		"created_at":       req.CreatedAt.AsTime(),
		"updated_at":       req.UpdatedAt.AsTime(),
	})
	if err != nil {
		log.Printf("Failed to create order: %v", err)
		return nil, err
	}

	return req, nil
}

func (g *GameStorage) GetOrder(ctx context.Context, req *game.OrderIdRequest) (*game.Order, error) {
	coll := g.db.Collection("orders")
	var order game.Order
	err := coll.FindOne(ctx, bson.M{"id": req.Id}).Decode(&order)
	if err != nil {
		log.Printf("Failed to get order: %v", err)
		return nil, err
	}

	return &order, nil
}

func (g *GameStorage) UpdateOrder(ctx context.Context, req *game.Order) (*game.Order, error) {
	coll := g.db.Collection("orders")
	_, err := coll.UpdateOne(ctx, bson.M{"id": req.Id}, bson.M{
		"$set": bson.M{
			"user_id":          req.UserId,
			"courier_id":       req.CourierId,
			"status":           req.Status,
			"total_amount":     req.TotalAmount,
			"delivery_address": req.DeliveryAddress,
			"updated_at":       req.UpdatedAt.AsTime(),
		},
	})
	if err != nil {
		log.Printf("Failed to update order: %v", err)
		return nil, err
	}

	return req, nil
}

func (g *GameStorage) DeleteOrder(ctx context.Context, req *game.OrderIdRequest) (*game.Empty, error) {
	coll := g.db.Collection("orders")
	_, err := coll.DeleteOne(ctx, bson.M{"id": req.Id})
	if err != nil {
		log.Printf("Failed to delete order: %v", err)
		return nil, err
	}

	return &game.Empty{}, nil
}

func (g *GameStorage) ListOrders(ctx context.Context, req *game.ListRequest) (*game.OrderListResponse, error) {
	coll := g.db.Collection("orders")
	cursor, err := coll.Find(ctx, bson.M{}, &mongo.FindOptions{
		Limit: int64(req.Limit),
		Skip:  int64(req.Offset),
	})
	if err != nil {
		log.Printf("Failed to list orders: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []*game.Order
	for cursor.Next(ctx) {
		var order game.Order
		if err := cursor.Decode(&order); err != nil {
			log.Printf("Failed to decode order: %v", err)
			return nil, err
		}
		orders = append(orders, &order)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
		return nil, err
	}

	return &game.OrderListResponse{Orders: orders}, nil
}
