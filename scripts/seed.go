package main

import (
	"context"
	"fmt"
	"log"

	"github.com/RianNegreiros/hotel-reservation/api"
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
	userStore  db.UserStore
	ctx        = context.Background()
)

func seedUser(firstName, lastName, email, password string, isAdmin bool) {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  "123456",
	})

	if err != nil {
		log.Fatal(err)
	}

	user.IsAdmin = isAdmin
	_, err = userStore.InsertUser(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	token := api.CreateTokenFromUser(user)

	fmt.Printf("User %s created with token: %s\n", user.Email, token)
}

func seedHotel(name, location string, rating float64) {
	hotel := &types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
		Rating:   rating,
	}

	rooms := []*types.Room{
		{
			Size:      "small",
			BasePrice: 100.0,
		},
		{
			Size:      "medium",
			BasePrice: 200.0,
		},
		{
			Size:      "large",
			BasePrice: 300.0,
		},
	}

	insertedHotel, err := hotelStore.Insert(ctx, hotel)
	if err != nil {
		log.Fatal(err)
	}
	for _, room := range rooms {
		room.HotelID = insertedHotel.ID
		_, err := roomStore.Insert(ctx, room)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	seedHotel("Hotel A", "Location A", 4.5)
	seedHotel("Hotel B", "Location B", 3.5)
	seedHotel("Hotel C", "Location C", 2.5)
	seedUser("John", "Doe", "johndoe@mail.com", "123456", true)
	seedUser("James", "Doe", "jamesdoe@mail.com", "123456", false)
}

func init() {
	var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
	userStore = db.NewMongoUserStore(client)
}
