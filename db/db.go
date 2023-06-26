package db

const (
	DBNAME = "hotel-reservation"
	TESTDB = "hotel-reservation-test"
	DBURI  = "mongodb://localhost:27018"
)

type Store struct {
	User  UserStore
	Hotel HotelStore
	Room  RoomStore
}
