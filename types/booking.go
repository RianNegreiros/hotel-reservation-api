package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID     primitive.ObjectID `bson:"userID,omitempty" json:"userID,omitempty"`
	RoomID     primitive.ObjectID `bson:"roomID,omitempty" json:"roomID,omitempty"`
	NumPersons int                `bson:"numPersons,omitempty" json:"numPersons,omitempty"`
	FromDate   time.Time          `bson:"fromDate,omitempty" json:"fromDate,omitempty"`
	ToDate     time.Time          `bson:"toDate,omitempty" json:"toDate,omitempty"`
	Canceled   bool               `bson:"canceled,omitempty" json:"canceled,omitempty"`
}
