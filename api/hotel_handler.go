package api

import (
	"github.com/RianNegreiros/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	hotelStore db.HotelStore
	roomStore  db.RoomStore
}

func NewHotelHandler(hotelStore db.HotelStore, roomStore db.RoomStore) *HotelHandler {
	return &HotelHandler{hotelStore: hotelStore, roomStore: roomStore}
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var (
		filter = bson.M{}
	)

	hotels, err := h.hotelStore.All(c.Context(), filter)
	if err != nil {
		return err
	}

	return c.JSON(hotels)
}

func (h *HotelHandler) HandleGetRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": oid}
	rooms, err := h.roomStore.All(c.Context(), filter)
	if err != nil {
		return err
	}

	return c.JSON(rooms)
}
