package ad

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
	NewSearchController(m.app, NewMgmService()).RegisterRoute()
	NewActiveController(m.app, NewMgmService()).RegisterRoute()
	NewController(m.app, NewMgmService()).RegisterRoute()
	NewFundingController(m.app, escrow.NewEscrow(), NewServiceMdmAdapter()).RegisterRoute()
}
