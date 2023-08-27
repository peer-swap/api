package ad

import (
	"github.com/gofiber/fiber/v2"
	"peerswap/escrow"
)

type FundingServiceEscrowInterface interface {
	ConfirmTransactionAndAmount(escrow.FundControllerAddTokenInputDto) (bool, error)
	ConfirmAllowanceMatchAmount(escrow.FundControllerAddCoinInputDto) (bool, error)
	DepositCoin(escrow.DepositCoinInputDto) error
}

type FundingServiceDbInterface interface {
	Find(id string) (*Dto, error)
	UpdateBalance(id string, amount float64) (*Dto, error)
}

type FundingController struct {
	app    *fiber.App
	escrow FundingServiceEscrowInterface
	db     FundingServiceDbInterface
}

func NewFundingController(app *fiber.App, escrow FundingServiceEscrowInterface, db FundingServiceDbInterface) *FundingController {
	return &FundingController{app: app, escrow: escrow, db: db}
}

func (c FundingController) RegisterRoute() {
	c.app.Get("ad/:id/fund/token", c.addToken())
	c.app.Get("ad/:id/fund/coin", c.addCoin())
}

func (c FundingController) addToken() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var input escrow.FundControllerAddTokenInputDto
		err := ctx.BodyParser(input)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "There was a parsing error")
		}

		ad, err := c.db.Find(ctx.Params("id"))
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "The requested resource was not found.")
		} else if ad.Balance != 0 {
			return fiber.NewError(fiber.StatusForbidden, "Forbidden")
		}

		input.Amount = ad.Amount
		confirmed, err := c.escrow.ConfirmTransactionAndAmount(input)
		if !confirmed {
			return fiber.NewError(fiber.StatusForbidden, "Forbidden")
		} else if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}

		ad, err = c.db.UpdateBalance(ad.Id, ad.Amount)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}

		return ctx.Status(200).JSON(map[string]interface{}{
			"data": ad,
		})
	}
}

func (c FundingController) addCoin() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var input escrow.FundControllerAddCoinInputDto
		err := ctx.BodyParser(input)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "There was a parsing error")
		}

		ad, err := c.db.Find(ctx.Params("id"))
		if ad.Balance != 0 {
			return fiber.NewError(fiber.StatusForbidden, "Forbidden")
		} else if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "The requested resource was not found.")
		}

		input.Amount = ad.Amount
		confirmed, err := c.escrow.ConfirmAllowanceMatchAmount(input)
		if !confirmed {
			return fiber.NewError(fiber.StatusForbidden, "Forbidden")
		} else if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}
		err = c.escrow.DepositCoin(escrow.DepositCoinInputDto{
			ChainId:         input.ChainId,
			MerchantAddress: input.MerchantAddress,
			TokenAddress:    input.TokenAddress,
			Amount:          input.Amount,
		})
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}

		ad, err = c.db.UpdateBalance(ad.Id, ad.Amount)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}

		return ctx.Status(200).JSON(ad)
	}
}
