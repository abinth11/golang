package movies

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	MovieName string             `json:"movieName"`
	Duration  float32            `json:"duration"`
	Rating    float32            `json:"rating"`
}
