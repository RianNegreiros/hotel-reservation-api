package api

import (
	"github.com/RianNegreiros/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HotelHandler struct {
	store *db.Store
}

func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	page := 1
	limit := 10
	opts := options.FindOptions{}
	opts.SetLimit(int64(limit))
	opts.SetSkip(int64((page - 1) * limit))
	hotels, err := h.store.Hotel.GetAll(c.Context(), nil, &opts)
	if err != nil {
		return ErrNotResourceNotFound("hotels")
	}
	return c.JSON(hotels)
}

func (h *HotelHandler) HandleGetRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID()
	}

	filter := bson.M{"hotelID": oid}
	rooms, err := h.store.Room.GetAll(c.Context(), filter)
	if err != nil {
		return ErrNotResourceNotFound("hotel")
	}
	return c.JSON(rooms)
}

func (h *HotelHandler) HandleGetHotel(c *fiber.Ctx) error {
	id := c.Params("id")
	hotel, err := h.store.Hotel.GetByID(c.Context(), id)
	if err != nil {
		return ErrNotResourceNotFound("hotel with id: " + id)
	}

	return c.JSON(hotel)
}
