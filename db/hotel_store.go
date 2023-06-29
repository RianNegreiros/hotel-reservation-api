package db

import (
	"context"
	"os"

	"github.com/RianNegreiros/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelStore interface {
	Insert(context.Context, *types.Hotel) (*types.Hotel, error)
	Update(context.Context, Map, Map) error
	GetAll(context.Context, Map) ([]*types.Hotel, error)
	GetByID(context.Context, string) (*types.Hotel, error)
}

type MongoHotelStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client) *MongoHotelStore {
	dbname := os.Getenv(MongoDBNameEnvName)
	return &MongoHotelStore{
		client:     client,
		collection: client.Database(dbname).Collection("hotels"),
	}
}

func (s *MongoHotelStore) Insert(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	resp, err := s.collection.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	hotel.ID = resp.InsertedID.(primitive.ObjectID)
	return hotel, nil
}

func (s *MongoHotelStore) Update(ctx context.Context, filter Map, update Map) error {
	_, err := s.collection.UpdateOne(ctx, filter, update)
	return err
}

func (s *MongoHotelStore) GetAll(ctx context.Context, filter Map) ([]*types.Hotel, error) {
	resp, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var hotels []*types.Hotel
	if err := resp.All(ctx, &hotels); err != nil {
		return nil, err
	}
	return hotels, nil
}

func (s *MongoHotelStore) GetByID(ctx context.Context, id string) (*types.Hotel, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var hotel *types.Hotel
	if err := s.collection.FindOne(ctx, Map{"_id": oid}).Decode(&hotel); err != nil {
		return nil, err
	}
	return hotel, nil
}
