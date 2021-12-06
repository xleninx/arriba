package handlers

import (
	"arriba/internal/domain"
	"arriba/internal/domain/constants"
	"arriba/internal/domain/external"
	"arriba/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Sell(c echo.Context, ctx *domain.ArribaContext) error {
	request := new(external.SellRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	amountToDeposit := ctx.AssetsProvider[request.Currency].Price * request.Amount

	_, err := services.AddMovement(ctx, request.UserID, request.Amount, request.Currency, constants.Sell)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	balance, _ := services.AddMovement(ctx, request.UserID, amountToDeposit, constants.USD, constants.Deposit)

	return c.JSON(200, balance)
}
