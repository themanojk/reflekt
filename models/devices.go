package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	MacAddress string             `bson:"mac_address"`
	CreatedAt  *time.Time         `bson:"created_at"`
	UpdatedAt  *time.Time         `bson:"updated_at"`
}
