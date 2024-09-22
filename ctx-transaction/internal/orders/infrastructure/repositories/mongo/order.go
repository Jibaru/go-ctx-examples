package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/jibaru/ctx-transaction/internal/orders/domain"
	"github.com/jibaru/ctx-transaction/internal/shared/app"
)

type MongoOrderRepository struct {
	collection *mongo.Collection
}

func NewMongoOrderRepository(db *mongo.Database) *MongoOrderRepository {
	return &MongoOrderRepository{
		collection: db.Collection("orders"),
	}
}

func (r *MongoOrderRepository) Save(ctx context.Context, order *domain.Order) error {
	session, ok := ctx.Value(app.SessionKey).(mongo.SessionContext)
	if ok {
		_, err := r.collection.InsertOne(session, order, options.InsertOne())
		return err
	}

	_, err := r.collection.InsertOne(ctx, order, options.InsertOne())
	return err
}
