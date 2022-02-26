package handler

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gofiber/fiber/v2"
)

type CreditCard struct {
	Id         int    `json:"id" fake:"{number:1000000,99999999}"`
	Uuid       string `json:"uuid" fake:"{uuid}"`
	CreditCard struct {
		Type   string `json:"type" fake:"{creditcardtype}"`
		Number string `json:"number" fake:"{creditcardnumber}"`
		Exp    string `json:"exp" fake:"{creditcardexp}"`
		Cvv    string `json:"cvv" fake:"{creditcardcvv}"`
	} `json:"credit_card"`
	CreatedAt time.Time `json:"createdAt"`
}

// @Tags Generate
// @Produce json
// @Success 200 {object} CreditCard
// @Failure 500 {object} Response
// @Router /v1/credit_card [get]
func GenerateCreditCard(c *fiber.Ctx) error {
	var creditcard CreditCard
	if err := gofakeit.Struct(&creditcard); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	// CreatedAt
	creditcard.CreatedAt = time.Now()

	return c.Status(fiber.StatusOK).JSON(Response{
		Status: fiber.StatusOK,
		Data:   &creditcard,
	})
}
