package order

import (
	"github.com/gofiber/fiber/v2"
	"peerswap/order/core"
	"peerswap/order/core/dto"
	"peerswap/reusable"
)

type Controller struct {
	*fiber.App
	service *core.Service
}

func (c Controller) RegisterRoute() {
	c.Post("order/store", c.store())
	c.Patch("order/:id/payment/sent", c.paymentSent())
	c.Patch("order/:id/payment/received", c.paymentReceived())
	c.Patch("order/:id/appeal", c.appeal())
	c.Patch("order/:id/cancel", c.cancel())
}

func (c Controller) store() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var input = &dto.ServiceStoreInput{}
		if err := ctx.BodyParser(input); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "There was a parsing error")
		}

		order, ad, err := c.service.Store(input)
		if errors, ok := err.(*reusable.ValidationErrorsMapper); ok {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(errors.ErrorMessageMap())
		} else if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "The service has encountered an error.")
		}

		data := fiber.Map{"data": map[string]interface{}{"order": order, "ad": ad}}
		return ctx.Status(fiber.StatusCreated).JSON(data)
	}
}

func (c Controller) paymentSent() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		order, err := c.service.PaymentSent(ctx.Params("id"))
		if err == core.NotFoundError {
			return fiber.ErrNotFound
		} else if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "The service has encountered an error.")
		}

		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"data": order})
	}
}

func (c Controller) paymentReceived() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		order, err := c.service.PaymentReceived(ctx.Params("id"))
		if err == core.NotFoundError {
			return fiber.ErrNotFound
		} else if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "The service has encountered an error.")
		}

		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"data": order})
	}
}

func (c Controller) appeal() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var input = &dto.ServiceAppealInput{}
		if err := ctx.BodyParser(input); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "There was a parsing error")
		}

		order, err := c.service.Appeal(ctx.Params("id"), input)
		if err == core.NotFoundError {
			return fiber.ErrNotFound
		} else if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "The service has encountered an error.")
		}

		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"data": order})
	}
}

func (c Controller) cancel() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var input = &dto.ServiceCancelInput{}
		if err := ctx.BodyParser(input); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "There was a parsing error")
		}

		order, err := c.service.Cancel(ctx.Params("id"), input)
		if err == core.NotFoundError {
			return fiber.ErrNotFound
		} else if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "The service has encountered an error.")
		}

		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"data": order})
	}
}

func NewController(app *fiber.App, service *core.Service) *Controller {
	return &Controller{App: app, service: service}
}
