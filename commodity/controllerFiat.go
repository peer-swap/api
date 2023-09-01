package commodity

import (
	"github.com/gofiber/fiber/v2"
	"peerswap/commodity/core/dto"
	"peerswap/commodity/core/service"
	"peerswap/reusable"
)

type FiatController struct {
	app     *fiber.App
	service *service.Fiat
}

func (c FiatController) RegisterRoute() {
	c.app.Post("/fiat", c.store())
	c.app.Get("/fiat/list", c.list())
}

func (c FiatController) store() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		input := &dto.FiatAddInput{}
		if err := ctx.BodyParser(input); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "There was a parsing error")
		}

		fiat, err := c.service.Add(input)
		if errors, ok := err.(*reusable.ValidationErrorsMapper); ok {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(errors.ErrorMessageMap())
		} else if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "The service has encountered an error.")
		}

		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"data": fiat})
	}
}

func (c FiatController) list() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		fiats, err := c.service.List(ctx.Query("query"))
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "The service has encountered an error.")
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": fiats,
		})
	}
}
