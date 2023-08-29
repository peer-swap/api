package ad

import (
	"github.com/gofiber/fiber/v2"
	"peerswap/ad/core/service"
	"peerswap/ad/escrow"
	"peerswap/ad/mongo"
	"peerswap/reusable"
)

type Module struct {
	app   *fiber.App
	event reusable.Event
}

func NewModule(app *fiber.App, event reusable.Event) *Module {
	return &Module{app: app, event: event}
}

func (m Module) Register() {
	NewController(m.app, m.service()).RegisterRoute()
	NewFundingController(m.app, service.NewFunding(escrow.NewAdapter(), mongo.NewAdapter(), m.event)).RegisterRoute()
}

func (m Module) service() *service.Service {
	return service.NewService(mongo.NewAdapter(), m.event)
}
