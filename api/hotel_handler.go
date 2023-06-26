package api

import (
	"github.com/RianNegreiros/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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
