package api

import (
	"github.com/RianNegreiros/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	store *db.Store
}

func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

type ResourceResponse struct {
	Total int `json:"total"`
	Data  any `json:"data"`
	Page  int `json:"page"`
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var pagination db.Pagination
	if err := c.QueryParser(&pagination); err != nil {
		return ErrInvalidQueryParams()
	}

	hotels, err := h.store.Hotel.GetAll(c.Context(), nil, &pagination)
	if err != nil {
		return ErrResourceNotFound("hotels")
	}

	resp := ResourceResponse{
		Total: len(hotels),
		Data:  hotels,
		Page:  int(pagination.Page),
	}

	return c.JSON(resp)
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
		return ErrResourceNotFound("hotel")
	}
	return c.JSON(rooms)
}

func (h *HotelHandler) HandleGetHotel(c *fiber.Ctx) error {
	id := c.Params("id")
	hotel, err := h.store.Hotel.GetByID(c.Context(), id)
	if err != nil {
		return ErrResourceNotFound("hotel with id: " + id)
	}

	return c.JSON(hotel)
}
