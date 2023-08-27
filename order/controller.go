package order

import (
	"github.com/gofiber/fiber/v2"
	"peerswap/order/dto"
	"peerswap/reusable"
)

type Controller struct {
	*fiber.App
	service *Service
}

func (c Controller) RegisterRoute() {
	c.Post("order/store", c.store())
}

func (c Controller) store() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var input = &dto.ServiceStoreInput{}
		if err := ctx.BodyParser(input); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "There was a parsing error")
		}

		order, _, err := c.service.store(input)
		if errors, ok := err.(*reusable.ValidationErrorsMapper); ok {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(errors.ErrorMessageMap())
		} else if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "The service has encountered an error.")
		}

		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"data": order})
	}
}

func NewController(app *fiber.App, service *Service) *Controller {
	return &Controller{App: app, service: service}
}
