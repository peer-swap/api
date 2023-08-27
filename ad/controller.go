package ad

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
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

func (d ControllerStoreInputDto) toStoreInputDto() StoreInputDto {
	return StoreInputDto{
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
	service *Service
}

func NewController(app *fiber.App, service *Service) *Controller {
	return &Controller{app: app, service: service}
}

func (c Controller) RegisterRoute() {
	c.app.Post("ad/store", c.store())
	c.app.Get("ad/list", c.list())
	c.app.Get("ad/:id", c.show())
}

func (c Controller) list() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		dto := ServiceListInputDto{
			Type:           reusable.TransactionType(ctx.Query("type")),
			Amount:         ctx.QueryFloat("amount"),
			Fiat:           ctx.Query("fiat"),
			Asset:          ctx.Query("asset"),
			PaymentMethods: []string{ctx.Query("payment_methods")},
			ChainId:        uint(ctx.QueryInt("chain_id")),
		}
		ads, err := c.service.List(dto)
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
