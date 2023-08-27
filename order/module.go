package order

import (
	"github.com/gofiber/fiber/v2"
	"peerswap/escrow"
)

type Module struct {
	app *fiber.App
}

func NewModule(app *fiber.App) *Module {
	return &Module{app: app}
}

func (m Module) Register() {
	NewController(m.app, NewService(NewAdapterMgm(), escrow.NewEscrow())).RegisterRoute()
}
