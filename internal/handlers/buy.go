package handlers

import (
	"arriba/internal/domain"
	"arriba/internal/domain/constants"
	"arriba/internal/domain/external"
	"arriba/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Buy(c echo.Context, ctx *domain.ArribaContext) error {
	request := new(external.BuyRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	amountToDebit := ctx.AssetsProvider[request.Currency].Price * request.Amount

	_, err := services.AddMovement(ctx, request.UserID, amountToDebit, constants.USD, constants.Debit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	balance, err := services.AddMovement(ctx, request.UserID, request.Amount, request.Currency, constants.Buy)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(200, balance)
}
