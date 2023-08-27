package ad

import (
	"github.com/gofiber/fiber/v2"
	"peerswap/reusable"
)

type SearchController struct {
	app     *fiber.App
	service *Service
}

func NewSearchController(app *fiber.App, service *Service) *SearchController {
	return &SearchController{app: app, service: service}
}

func (c SearchController) RegisterRoute() {
	c.app.Get("ad/search", c.search())
}

func (c SearchController) search() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		inputDto := ServiceListInputDto{
			Type:           reusable.TransactionType(ctx.Query("type")),
			Amount:         ctx.QueryFloat("amount"),
			Fiat:           ctx.Query("fiat"),
			Asset:          ctx.Query("asset"),
			PaymentMethods: []string{ctx.Query("payment_methods")},
			ChainId:        uint(ctx.QueryInt("chain_id")),
		}
		ads, err := c.service.Search(inputDto)
		if ve, ok := err.(*reusable.ValidationErrorsMapper); ok {
			data := fiber.Map{"message": "The service has encountered an error.", "data": ve.ErrorMap()}
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(data)
		} else if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "The service has encountered an error.")
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": ads,
		})
	}
}
