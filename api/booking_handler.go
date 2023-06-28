package api

import (
	"github.com/RianNegreiros/hotel-reservation/db"
	"github.com/RianNegreiros/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

type BookingHandler struct {
	store *db.Store
}

func NewBookingHandler(store *db.Store) *BookingHandler {
	return &BookingHandler{store: store}
}

func (h *BookingHandler) HandleGetBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := h.store.Booking.GetByID(c.Context(), id)
	if err != nil {
		return err
	}

	user, ok := c.Context().Value("user").(*types.User)
	if !ok {
		return fiber.ErrUnauthorized
	}

	if booking.UserID != user.ID {
		return fiber.ErrUnauthorized
	}

	return c.JSON(booking)
}

func (h *BookingHandler) HandleGetBookings(c *fiber.Ctx) error {
	bookings, err := h.store.Booking.GetAll(c.Context(), nil)
	if err != nil {
		return err
	}
	return c.JSON(bookings)
}
