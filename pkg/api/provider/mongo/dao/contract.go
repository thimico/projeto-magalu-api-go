package dao

import (
	"context"
	"fmt"
	"projeto-magalu-api-go/pkg/api/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Contract struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoContract(client *mongo.Client, db *mongo.Database) *Contract {
	return &Contract{client: client, collection: db.Collection("contracts")}
}

func (p *Contract) Create(ctx context.Context, contract *model.Contract) (*model.Contract, error) {
	if contract == nil {
		return nil, fmt.Errorf("Error on parsing")
	}
	one, err := p.collection.InsertOne(ctx, contract)
	if err != nil {
		return nil, err
	}
	result := p.collection.FindOne(ctx, bson.M{"_id": one.InsertedID.(primitive.ObjectID)})
	var contractOut model.Contract
	err = result.Decode(&contractOut)
	if err != nil {
		return nil, err
	}
	return &contractOut, nil
}

func (p *Contract) GetByID(ctx context.Context, id string) (*model.Contract, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result := p.collection.FindOne(ctx, bson.M{"_id": oID})
	var contract model.Contract
	err = result.Decode(&contract)
	if err != nil {
		return nil, err
	}
	return &contract, nil
}
