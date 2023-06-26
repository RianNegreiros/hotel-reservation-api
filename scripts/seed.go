package main

import (
	"context"
	"log"

	"github.com/RianNegreiros/hotel-reservation/db"
	"github.com/RianNegreiros/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	hotelStore db.HotelStore
	roomStore  db.RoomStore
	ctx        = context.Background()
)

func seedHotel(name, location string, rating float64) {
	hotel := &types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
		Rating:   rating,
	}

	rooms := []*types.Room{
		{
			Type:      types.SingleRoomType,
			BasePrice: 100,
		},
		{
			Type:      types.DoubleRoomType,
			BasePrice: 200,
		},
		{
			Type:      types.SeaSideRoomType,
			BasePrice: 300,
		},
	}

	insertedHotel, err := hotelStore.Insert(ctx, hotel)
	if err != nil {
		log.Fatal(err)
	}
	for _, room := range rooms {
		room.HotelID = insertedHotel.ID
		_, err := roomStore.InsertRoom(ctx, room)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	seedHotel("Hotel A", "Location A", 4.5)
	seedHotel("Hotel B", "Location B", 3.5)
	seedHotel("Hotel C", "Location C", 2.5)
}

func init() {
	var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
}
