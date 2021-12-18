package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CovidStat struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	StateCode     string             `bson:"state_code" json:"state_code"`
	StateCount    int                `bson:"state_count" json:"state_count"`
	LastUpdatedAt time.Time          `bson:"last_updated_at" json:"last_updated_at"`
}
