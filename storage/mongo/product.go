package postgres

import (
	"context"
	"log"

	game "github.com/dilshodforever/4-oyimtixon-game-service/genprotos/product"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GameStorage struct {
	db *mongo.Database
}

func NewProductStorage(db *mongo.Database) *GameStorage {
	return &GameStorage{db: db}
}

// CreateProduct inserts a new product into the database
func (g *GameStorage) CreateProduct(ctx context.Context, req *game.Product) (*game.Product, error) {
	coll := g.db.Collection("products")
	_, err := coll.InsertOne(ctx, bson.M{
		"id":          req.Id,
		"name":        req.Name,
		"description": req.Description,
		"price":       req.Price,
		"image_url":   req.ImageUrl,
		"created_at":  req.CreatedAt.AsTime(),
		"updated_at":  req.UpdatedAt.AsTime(),
	})
	if err != nil {
		log.Printf("Failed to create product: %v", err)
		return nil, err
	}

	return req, nil
}

// GetProduct retrieves a product by its ID
func (g *GameStorage) GetProduct(ctx context.Context, req *game.ProductIdRequest) (*game.Product, error) {
	coll := g.db.Collection("products")
	var product game.Product
	err := coll.FindOne(ctx, bson.M{"id": req.Id}).Decode(&product)
	if err != nil {
		log.Printf("Failed to get product: %v", err)
		return nil, err
	}

	return &product, nil
}

// UpdateProduct updates an existing product
func (g *GameStorage) UpdateProduct(ctx context.Context, req *game.Product) (*game.Product, error) {
	coll := g.db.Collection("products")
	_, err := coll.UpdateOne(ctx, bson.M{"id": req.Id}, bson.M{
		"$set": bson.M{
			"name":        req.Name,
			"description": req.Description,
			"price":       req.Price,
			"image_url":   req.ImageUrl,
			"updated_at":  req.UpdatedAt.AsTime(),
		},
	})
	if err != nil {
		log.Printf("Failed to update product: %v", err)
		return nil, err
	}

	return req, nil
}

// DeleteProduct removes a product by its ID
func (g *GameStorage) DeleteProduct(ctx context.Context, req *game.ProductIdRequest) (*game.Empty, error) {
	coll := g.db.Collection("products")
	_, err := coll.DeleteOne(ctx, bson.M{"id": req.Id})
	if err != nil {
		log.Printf("Failed to delete product: %v", err)
		return nil, err
	}

	return &game.Empty{}, nil
}

// ListProducts lists all products with pagination
func (g *GameStorage) ListProducts(ctx context.Context, req *game.ListRequest) (*game.ProductListResponse, error) {
	coll := g.db.Collection("products")
	cursor, err := coll.Find(ctx, bson.M{}, &mongo.FindOptions{
		Limit: int64(req.Limit),
		Skip:  int64(req.Offset),
	})
	if err != nil {
		log.Printf("Failed to list products: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*game.Product
	for cursor.Next(ctx) {
		var product game.Product
		if err := cursor.Decode(&product); err != nil {
			log.Printf("Failed to decode product: %v", err)
			return nil, err
		}
		products = append(products, &product)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
		return nil, err
	}

	return &game.ProductListResponse{Products: products}, nil
}
