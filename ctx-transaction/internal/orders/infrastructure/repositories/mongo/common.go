package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

type commonRepository struct {
}

func (r *commonRepository) NextID() any {
	return primitive.NewObjectID()
}
