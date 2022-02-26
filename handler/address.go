package handler

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gofiber/fiber/v2"
)

type Address struct {
	Id       int    `json:"id" fake:"{number:1000000,99999999}"`
	Uuid     string `json:"uuid" fake:"{uuid}"`
	Location struct {
		Country     string `json:"country" fake:"{country}"`
		City        string `json:"city" fake:"{city}"`
		State       string `json:"state" fake:"{state}"`
		Street      string `json:"street" fake:"{street}"`
		ZipCode     int    `json:"zipCode" fake:"{zip}"`
		Coordinates struct {
			Latitude  float64 `json:"latitude" fake:"{latitude}"`
			Longitude float64 `json:"longitude" fake:"{longitude}"`
		} `json:"coordinates"`
		TimeZone string `json:"timeZone" fake:"{timezoneregion}"`
	} `json:"location"`
	CreatedAt time.Time `json:"createdAt"`
}

// @Tags Generate
// @Produce json
// @Success 200 {object} Address
// @Failure 500 {object} Response
// @Router /v1/address [get]
func GenerateAddress(c *fiber.Ctx) error {
	var address Address
	if err := gofakeit.Struct(&address); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&Response{
			Status:  fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	// CreatedAt
	address.CreatedAt = time.Now()

	return c.Status(fiber.StatusOK).JSON(&Response{
		Status: fiber.StatusOK,
		Data:   &address,
	})
}
