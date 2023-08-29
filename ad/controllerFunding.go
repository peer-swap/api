package ad

import (
	"github.com/gofiber/fiber/v2"
	"peerswap/ad/core/dto"
	"peerswap/ad/core/service"
	"peerswap/reusable"
)

type FundingController struct {
	app     *fiber.App
	service *service.Funding
}

func NewFundingController(app *fiber.App, service *service.Funding) *FundingController {
	return &FundingController{app: app, service: service}
}

func (c FundingController) RegisterRoute() {
	c.app.Get("/ad/:id/fund/token", c.addToken())
	c.app.Get("/ad/:id/fund/coin", c.addErc20())
}

func (c FundingController) addToken() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var input = &dto.AdFundingAddTokenInput{}
		err := ctx.BodyParser(input)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "There was a parsing error")
		}

		ad, err := c.service.AddToken(ctx.Params("id"), input)
		return c.handleResponse(ctx, ad, err)
	}
}

func (c FundingController) addErc20() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var input = &dto.AdFundingAddErc20Input{}
		err := ctx.BodyParser(input)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "There was a parsing error")
		}

		ad, err := c.service.AddErc20(ctx.Params("id"), input)
		return c.handleResponse(ctx, ad, err)
	}
}

func (c FundingController) handleResponse(ctx *fiber.Ctx, ad *dto.Ad, err error) error {
	if err == service.AdFindError {
		return fiber.NewError(fiber.StatusNotFound, "Ad not found")
	} else if mapper, ok := err.(*reusable.ValidationErrorsMapper); ok {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(mapper.ErrorMessageMap())
	} else if err == service.FundingTransactionNotFoundError {
		return fiber.NewError(fiber.StatusNotFound, "Transaction not found")
	} else if err == service.FundingTransactionAmountError {
		return fiber.NewError(fiber.StatusForbidden, "Amount does not match supply")
	} else if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "The service has encountered an error.")
	}

	return ctx.Status(200).JSON(ad)
}
