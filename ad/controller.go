package ad

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"peerswap/ad/core/dto"
	"peerswap/ad/core/service"
	"peerswap/reusable"
)

type ControllerStoreInputDto struct {
	Type            reusable.TransactionType `json:"type"`
	Asset           string                   `json:"asset"`
	Fiat            string                   `json:"fiat"`
	Price           float64                  `json:"price"`
	Supply          float64                  `json:"supply"`
	PaymentMethods  []string                 `json:"payment_methods"`
	OrderLowerLimit float64                  `json:"order_lower_limit"`
	OrderUpperLimit float64                  `json:"order_upper_limit"`
	ChainId         uint                     `json:"chain_id"`
	AssetType       reusable.AssetType       `json:"asset_type"`
}

func (d ControllerStoreInputDto) toStoreInputDto() dto.StoreInputDto {
	return dto.StoreInputDto{
		Type:            d.Type,
		Asset:           d.Asset,
		AssetType:       d.AssetType,
		Fiat:            d.Fiat,
		Price:           d.Price,
		Supply:          d.Supply,
		PaymentMethods:  d.PaymentMethods,
		OrderLowerLimit: d.OrderLowerLimit,
		OrderUpperLimit: d.OrderUpperLimit,
		ChainId:         d.ChainId,
	}
}

type Controller struct {
	app     *fiber.App
	service *service.Service
}

func NewController(app *fiber.App, service *service.Service) *Controller {
	return &Controller{app: app, service: service}
}

func (c Controller) RegisterRoute() {
	c.app.Get("/ad/search", c.search())
	c.app.Get("/ad/list", c.list())
	c.app.Post("/ad/store", c.store())
	c.app.Patch("/ad/:id/active", c.updateActive())
	c.app.Get("/ad/:id", c.show())
}

func (c Controller) list() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		inputDto := dto.ServiceListInputDto{
			Type:           reusable.TransactionType(ctx.Query("type")),
			Amount:         ctx.QueryFloat("amount"),
			Fiat:           ctx.Query("fiat"),
			Asset:          ctx.Query("asset"),
			PaymentMethods: []string{ctx.Query("payment_methods")},
			ChainId:        uint(ctx.QueryInt("chain_id")),
		}
		ads, err := c.service.List(inputDto)
		if ve, ok := err.(*reusable.ValidationErrorsMapper); ok {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "The service has encountered an error.", "data": ve.ErrorMap()})
		} else if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "The service has encountered an error.")
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": ads,
		})
	}
}

func (c Controller) store() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var input = &ControllerStoreInputDto{}
		if err := ctx.BodyParser(input); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "There was a parsing error")
		}

		ad, err := c.service.Store(input.toStoreInputDto())
		if err != nil {
			if validationErrors, ok := err.(validator.ValidationErrors); ok {
				mapper := reusable.ValidationErrorsMapper{ValidationErrors: validationErrors}
				return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "The service has encountered an error.", "data": mapper.ErrorMap()})
			} else {
				return fiber.NewError(fiber.StatusInternalServerError, "The service has encountered an error.")
			}
		}

		return ctx.Status(201).JSON(fiber.Map{
			"data": ad,
		})
	}
}

func (c Controller) updateActive() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		body := make(map[string]interface{})
		if err := ctx.BodyParser(&body); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "couldn't parse input")
		}
		unknownActive := body["active"]
		active, ok := unknownActive.(bool)
		if !ok {
			return fiber.NewError(fiber.StatusBadRequest, "Item not found in request body")
		}

		ad, err := c.service.UpdateActive(ctx.Params("id"), active)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "db error")
		}

		return ctx.JSON(map[string]interface{}{
			"data": ad,
		})
	}
}

func (c Controller) show() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ad, err := c.service.Find(ctx.Params("id"))
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "The service has encountered an error.")
		}
		return ctx.Status(200).JSON(fiber.Map{
			"data": ad,
		})
	}
}

func (c Controller) search() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		inputDto := dto.ServiceListInputDto{
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
