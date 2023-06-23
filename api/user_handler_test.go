package api

import (
	"context"
	"testing"

	"github.com/RianNegreiros/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testdburi = "mongodb://localhost:27018"
	dbname    = "hotel-reservation"
)

type testdb struct {
	db.UserStore
}

func (db *testdb) teardown(t *testing.T) {
	if err := db.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testdburi))
	if err != nil {
		t.Fatal(err)
	}

	return &testdb{
		UserStore: db.NewMongoUserStore(client, dbname),
	}
}

func TestPostUser(t *testing.T) {
	db := setup(t)
	defer db.teardown(t)
}
