package order

import (
	"github.com/gofiber/fiber/v2"
	"peerswap/order/core"
	"peerswap/order/escrow"
	"peerswap/order/mongo"
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
}

func (m Module) service() *core.Service {
	return core.NewService(mongo.NewAdapter(), escrow.NewEscrow(), m.event)
}
