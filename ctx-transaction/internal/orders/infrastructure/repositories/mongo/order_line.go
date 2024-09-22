package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/jibaru/ctx-transaction/internal/orders/domain"
	"github.com/jibaru/ctx-transaction/internal/shared/app"
)

type MongoOrderLineRepository struct {
	collection *mongo.Collection
}

func NewMongoOrderLineRepository(db *mongo.Database) *MongoOrderLineRepository {
	return &MongoOrderLineRepository{
		collection: db.Collection("order_lines"),
	}
}

func (r *MongoOrderLineRepository) Save(ctx context.Context, orderLine *domain.OrderLine) error {
	session, ok := ctx.Value(app.SessionKey).(mongo.SessionContext)
	if ok {
		_, err := r.collection.InsertOne(session, orderLine, options.InsertOne())
		return err
	}

	_, err := r.collection.InsertOne(ctx, orderLine, options.InsertOne())
	return err
}
