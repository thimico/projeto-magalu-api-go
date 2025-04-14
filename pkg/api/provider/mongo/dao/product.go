package dao

import (
	"context"
	"fmt"
	"projeto-magalu-api-go/pkg/api/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Product struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoProduct(client *mongo.Client, db *mongo.Database) *Product {
	return &Product{client: client, collection: db.Collection("products")}
}

func (p *Product) Create(ctx context.Context, product *model.Product) (*model.Product, error) {
	if product == nil {
		return nil, fmt.Errorf("Error on parsing")
	}
	one, err := p.collection.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}
	result := p.collection.FindOne(ctx, bson.M{"_id": one.InsertedID.(primitive.ObjectID)})
	var productOut model.Product
	err = result.Decode(&productOut)
	if err != nil {
		return nil, err
	}
	return &productOut, nil
}

func (p *Product) GetByID(ctx context.Context, id string) (*model.Product, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result := p.collection.FindOne(ctx, bson.M{"_id": oID})
	var product model.Product
	err = result.Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
