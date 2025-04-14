package dao

import (
	"context"
	"fmt"
	"projeto-magalu-api-go/pkg/api/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Pulse struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (p *Pulse) GetCurrentMonthConsumption(ctx context.Context, tenant string, sku string) (float64, error) {
	now := time.Now()
	firstOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	filter := bson.M{
		"tenant":      tenant,
		"product_sku": sku,
		"timestamp": bson.M{
			"$gte": firstOfMonth,
		},
	}

	cursor, err := p.collection.Database().Collection("aggregated_pulses").Find(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("failed to find aggregated pulses: %w", err)
	}
	defer cursor.Close(ctx)

	var total float64

	for cursor.Next(ctx) {
		var result struct {
			TotalAmount float64 `bson:"total_amount"`
		}
		if err := cursor.Decode(&result); err != nil {
			return 0, fmt.Errorf("failed to decode document: %w", err)
		}
		total += result.TotalAmount
	}

	return total, nil
}

func (p *Pulse) GetAllResourcesConsumption(ctx context.Context, tenant string) (map[string]float64, error) {
	now := time.Now()
	firstOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	filter := bson.M{
		"tenant": tenant,
		"timestamp": bson.M{
			"$gte": firstOfMonth,
		},
	}

	cursor, err := p.collection.Database().Collection("aggregated_pulses").Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find aggregated pulses: %w", err)
	}
	defer cursor.Close(ctx)

	resultMap := make(map[string]float64)

	for cursor.Next(ctx) {
		var result struct {
			ProductSKU  string  `bson:"product_sku"`
			TotalAmount float64 `bson:"total_amount"`
		}
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("failed to decode document: %w", err)
		}
		resultMap[result.ProductSKU] += result.TotalAmount
	}

	return resultMap, nil
}

func NewMongoPulse(client *mongo.Client, db *mongo.Database) *Pulse {
	return &Pulse{client: client, collection: db.Collection("pulses")}
}

func (p *Pulse) Create(ctx context.Context, pulse *model.Pulse) (*model.Pulse, error) {
	if pulse == nil {
		return nil, fmt.Errorf("Error on parsing")
	}
	one, err := p.collection.InsertOne(ctx, pulse)
	if err != nil {
		return nil, err
	}
	result := p.collection.FindOne(ctx, bson.M{"_id": one.InsertedID})
	var pulseOut model.Pulse
	err = result.Decode(&pulseOut)
	if err != nil {
		return nil, err
	}
	return &pulseOut, nil
}

func (p *Pulse) AggregateAndStore(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	pipeline := mongo.Pipeline{
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: bson.D{
				{Key: "tenant", Value: "$tenant"},
				{Key: "product_sku", Value: "$product_sku"},
				{Key: "use_unity", Value: "$use_unity"},
			}},
			{Key: "total_amount", Value: bson.D{{Key: "$sum", Value: "$used_amount"}}},
		}}},
	}

	cursor, err := p.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return fmt.Errorf("aggregation failed: %w", err)
	}
	defer cursor.Close(ctx)

	type AggregatedPulse struct {
		ID struct {
			Tenant     string `bson:"tenant"`
			ProductSKU string `bson:"product_sku"`
			UseUnity   string `bson:"use_unity"`
		} `bson:"_id"`
		TotalAmount float64   `bson:"total_amount"`
		Timestamp   time.Time `bson:"timestamp"`
	}

	var results []interface{}

	now := time.Now()

	for cursor.Next(ctx) {
		var result AggregatedPulse
		if err := cursor.Decode(&result); err != nil {
			return fmt.Errorf("failed to decode aggregation result: %w", err)
		}
		result.Timestamp = now

		results = append(results, bson.D{
			{Key: "tenant", Value: result.ID.Tenant},
			{Key: "product_sku", Value: result.ID.ProductSKU},
			{Key: "use_unity", Value: result.ID.UseUnity},
			{Key: "total_amount", Value: result.TotalAmount},
			{Key: "timestamp", Value: result.Timestamp},
		})
	}

	if len(results) > 0 {
		aggCollection := p.collection.Database().Collection("aggregated_pulses")
		_, err = aggCollection.InsertMany(ctx, results)
		if err != nil {
			return fmt.Errorf("failed to store aggregated data: %w", err)
		}
	}

	return nil
}

func (p *Pulse) Check(ctx context.Context) error {
	ctx, _ = context.WithTimeout(ctx, time.Second)
	return p.client.Ping(ctx, nil)
}
