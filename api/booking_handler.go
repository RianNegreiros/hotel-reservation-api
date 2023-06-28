package api

import (
	"github.com/RianNegreiros/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
)

type BookingHandler struct {
	store *db.Store
}

func NewBookingHandler(store *db.Store) *BookingHandler {
	return &BookingHandler{store: store}
}

func (h *BookingHandler) HandleGetBooking(c *fiber.Ctx) error {
	return nil
}

func (h *BookingHandler) HandleGetBookings(c *fiber.Ctx) error {
	return nil
}
