package commodity

import (
	"github.com/gofiber/fiber/v2"
	"peerswap/commodity/core/dto"
	"peerswap/commodity/core/service"
	"peerswap/reusable"
)

type TokenController struct {
	app     *fiber.App
	service *service.Token
}

func (c TokenController) RegisterRoute() {
	c.app.Post("/token", c.store())
	c.app.Get("/token/list", c.list())

}

func (c TokenController) store() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		input := &dto.TokenAddInput{}
		if err := ctx.BodyParser(input); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "There was a parsing error")
		}

		token, err := c.service.Add(input)
		if errors, ok := err.(*reusable.ValidationErrorsMapper); ok {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(errors.ErrorMessageMap())
		} else if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "The service has encountered an error.")
		}

		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"data": token})
	}
}

func (c TokenController) list() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokens, err := c.service.List(dto.TokenListFilter{
			Query:   ctx.Query("query"),
			ChainId: ctx.QueryInt("chain_id"),
		})
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "The service has encountered an error.")
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": tokens,
		})

	}
}
