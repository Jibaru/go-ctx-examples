package mongo

import (
	"context"

	"github.com/jibaru/ctx-transaction/internal/shared/app"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTransactional struct {
	client *mongo.Client
}

func NewMongoTransactional(client *mongo.Client) *MongoTransactional {
	return &MongoTransactional{client: client}
}

func (t *MongoTransactional) InTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	session, err := t.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	mongoCtx := context.WithValue(ctx, app.SessionKey, session)
	_, err = session.WithTransaction(mongoCtx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		return nil, fn(sessCtx)
	})

	return err
}
