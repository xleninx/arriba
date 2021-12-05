package handlers

import (
	"arriba/internal/domain"
	"arriba/internal/domain/constants"
	"arriba/internal/domain/external"
	"arriba/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Sell(c echo.Context, ctx domain.ArribaContext) error {
	request := new(external.SellRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	if !request.Currency.IsValid() {
		return echo.NewHTTPError(http.StatusBadRequest, string(constants.CurrencyInvalid))
	}

	amountToDeposit := ctx.AssetsProvider[request.Currency].Price * request.Amount
	account := ctx.Users[request.UserID].Account
	total := services.GetTotalAsset(account, request.Currency)

	if total < request.Amount {
		return echo.NewHTTPError(http.StatusBadRequest, string(constants.InsufficientFounds))
	}

	services.AddMovement(ctx, request.UserID, request.Amount, request.Currency, constants.Sell)
	balance := services.AddMovement(ctx, request.UserID, amountToDeposit, constants.USD, constants.Deposit)

	return c.JSON(200, balance)
}
