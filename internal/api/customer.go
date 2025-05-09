package api

import (
	"PerpusGo/domain"
	"PerpusGo/dto"
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CustomerAPI struct {
	customerService domain.CustomerService
}

func NewCustomer(app *fiber.App, cutomerService domain.CustomerService) {
	ca := CustomerAPI{
		customerService: cutomerService,
	}

	app.Get("/customer", ca.Index)
}

func (ca *CustomerAPI) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := ca.customerService.Index(c)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.JSON(dto.CreateResponseSuccess(res))
}
