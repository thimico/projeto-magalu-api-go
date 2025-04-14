package dao

import (
	"context"
	"fmt"
	"projeto-magalu-api-go/pkg/api/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Tenant struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoTenant(client *mongo.Client, db *mongo.Database) *Tenant {
	return &Tenant{client: client, collection: db.Collection("tenants")}
}

func (p *Tenant) Create(ctx context.Context, tenant *model.Tenant) (*model.Tenant, error) {
	if tenant == nil {
		return nil, fmt.Errorf("Error on parsing")
	}
	one, err := p.collection.InsertOne(ctx, tenant)
	if err != nil {
		return nil, err
	}
	result := p.collection.FindOne(ctx, bson.M{"_id": one.InsertedID.(primitive.ObjectID)})
	var tenantOut model.Tenant
	err = result.Decode(&tenantOut)
	if err != nil {
		return nil, err
	}
	return &tenantOut, nil
}

func (p *Tenant) GetByID(ctx context.Context, id string) (*model.Tenant, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result := p.collection.FindOne(ctx, bson.M{"_id": oID})
	var tenant model.Tenant
	err = result.Decode(&tenant)
	if err != nil {
		return nil, err
	}
	return &tenant, nil
}
