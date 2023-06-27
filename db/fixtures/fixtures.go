package fixtures

import (
	"context"
	"fmt"
	"log"

	"github.com/RianNegreiros/hotel-reservation/db"
	"github.com/RianNegreiros/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddRoom(store *db.Store, size string, ss bool, price float64, hid primitive.ObjectID) *types.Room {
	room := &types.Room{
		Size:    size,
		Price:   price,
		HotelID: hid,
	}
	insertedRoom, err := store.Room.Insert(context.Background(), room)
	if err != nil {
		log.Fatal(err)
	}
	return insertedRoom
}

func AddHotel(store *db.Store, name string, loc string, rating float64, rooms []primitive.ObjectID) *types.Hotel {
	var roomIDS = rooms
	if rooms == nil {
		roomIDS = []primitive.ObjectID{}
	}
	hotel := types.Hotel{
		Name:     name,
		Location: loc,
		Rooms:    roomIDS,
		Rating:   rating,
	}
	insertedHotel, err := store.Hotel.Insert(context.TODO(), &hotel)
	if err != nil {
		log.Fatal(err)
	}
	return insertedHotel
}

func AddUser(store *db.Store, fn, ln string) *types.User {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		Email:     fmt.Sprintf("%s@%s.com", fn, ln),
		FirstName: fn,
		LastName:  ln,
		Password:  fmt.Sprintf("%s_%s", fn, ln),
	})
	if err != nil {
		log.Fatal(err)
	}
	insertedUser, err := store.User.InsertUser(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	return insertedUser
}
