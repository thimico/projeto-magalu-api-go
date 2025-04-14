package dao

import (
	"context"
	"fmt"
	"math"
	"projeto-magalu-api-go/pkg/api/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Transaction struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoTransaction(client *mongo.Client, db *mongo.Database) *Transaction {
	return &Transaction{client: client, collection: db.Collection("transactions")}
}

func (p *Transaction) Create(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error) {
	if transaction == nil {
		return nil, fmt.Errorf("Error on parsing")
	}
	one, err := p.collection.InsertOne(ctx, transaction)
	if err != nil {
		return nil, err
	}
	result := p.collection.FindOne(ctx, bson.M{"_id": one.InsertedID.(primitive.ObjectID)})
	var transactionOut model.Transaction
	err = result.Decode(&transactionOut)
	if err != nil {
		return nil, err
	}
	return &transactionOut, nil
}

func (p *Transaction) BalancePayment(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error) {
	tenantID, err := primitive.ObjectIDFromHex(transaction.TenantID.Hex())

	filter := bson.M{
		"tenant_id": tenantID,
		"balance":   bson.M{"$lt": 0},
	}

	cursor, err := p.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("erro base filtro transaction: %v", err)
	}

	remainningAmount := transaction.Amount

	for cursor.Next(ctx) {
		var transaction model.Transaction
		if err := cursor.Decode(&transaction); err != nil {
			return nil, fmt.Errorf("erro base filtro transaction: %v", err)
		}

		if remainningAmount <= 0 {
			break
		}

		amountAtual := math.Min(-transaction.Balance, remainningAmount)
		transaction.Balance += amountAtual
		remainningAmount -= amountAtual

		if _, err = p.collection.UpdateOne(ctx, bson.M{"_id": transaction.ID}, bson.M{"$set": bson.M{"balance": transaction.Balance}}); err != nil {
			return nil, fmt.Errorf("erro update transaction: %v", err)
		}
	}

	transaction.Balance = remainningAmount
	transaction.OperationTypeID = 4

	one, err := p.collection.InsertOne(ctx, transaction)
	if err != nil {
		return nil, err
	}
	result := p.collection.FindOne(ctx, bson.M{"_id": one.InsertedID.(primitive.ObjectID)})
	var transactionOut model.Transaction
	err = result.Decode(&transactionOut)
	if err != nil {
		return nil, err
	}
	return &transactionOut, nil
}

func (p *Transaction) Check(ctx context.Context) error {
	ctx, _ = context.WithTimeout(ctx, time.Second)
	return p.client.Ping(ctx, nil)
}
