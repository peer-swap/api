package commodity

import (
	"github.com/gofiber/fiber/v2"
	"peerswap/reusable"
)

type Module struct {
	app   *fiber.App
	event reusable.Event
}

func (m Module) Register() {

}

func NewModule(app *fiber.App, event reusable.Event) *Module {
	return &Module{app: app, event: event}
}
