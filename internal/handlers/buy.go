package handlers

import (
	"arriba/internal/domain"
	"arriba/internal/domain/constants"
	"arriba/internal/domain/external"
	"arriba/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Buy(c echo.Context, ctx domain.ArribaContext) error {
	request := new(external.BuyRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	if !request.Currency.IsValid() {
		return echo.NewHTTPError(http.StatusBadRequest, string(constants.CurrencyInvalid))
	}

	amountToDebit := ctx.AssetsProvider[request.Currency].Price * request.Amount
	account := ctx.Users[request.UserID].Account
	total := services.GetTotalAsset(account, constants.USD)

	if total < amountToDebit {
		return echo.NewHTTPError(http.StatusBadRequest, string(constants.InsufficientFounds))
	}

	services.AddMovement(ctx, request.UserID, amountToDebit, constants.USD, constants.Debit)
	balance := services.AddMovement(ctx, request.UserID, request.Amount, request.Currency, constants.Buy)

	return c.JSON(200, balance)
}
