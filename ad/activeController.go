package ad

import "github.com/gofiber/fiber/v2"

type ActiveController struct {
	app     *fiber.App
	service *Service
}

func NewActiveController(app *fiber.App, service *Service) *ActiveController {
	return &ActiveController{app: app, service: service}
}

func (s ActiveController) RegisterRoute() {
	s.app.Patch("ad/:id/active", s.update())

}

func (s ActiveController) update() fiber.Handler {
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

		ad, err := s.service.UpdateActive(ctx.Params("id"), active)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "db error")
		}

		return ctx.JSON(map[string]interface{}{
			"data": ad,
		})
	}
}
