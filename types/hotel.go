package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID       primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string               `json:"name" bson:"name"`
	Location string               `json:"location" bson:"location"`
	Rooms    []primitive.ObjectID `json:"rooms" bson:"rooms"`
	Rating   float64              `json:"rating" bson:"rating"`
}

type RoomType int

const (
	_ RoomType = iota
	SingleRoomType
	DoubleRoomType
	SeaSideRoomType
	DeluxeRoomType
)

type Room struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Size      string             `json:"size" bson:"size"`
	Single    bool               `json:"single" bson:"single"`
	Double    bool               `json:"double" bson:"double"`
	SeaSide   bool               `json:"seaSide" bson:"seaSide"`
	Deluxe    bool               `json:"deluxe" bson:"deluxe"`
	BasePrice float64            `json:"basePrice" bson:"basePrice"`
	Price     float64            `json:"price" bson:"price"`
	HotelID   primitive.ObjectID `json:"hotelId" bson:"hotelId"`
	Avaliable bool               `json:"-" bson:"avaliable"`
}
