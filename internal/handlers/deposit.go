package handlers

import (
	"arriba/internal/domain"
	"arriba/internal/domain/constants"
	"arriba/internal/domain/external"
	"arriba/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Deposit(c echo.Context, ctx *domain.ArribaContext) error {
	u := new(external.FiatRequest)
	if err := c.Bind(u); err != nil {
		return err
	}

	balance, err := services.AddMovement(ctx, u.UserID, u.Amount, constants.USD, constants.Deposit)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(200, balance)
}
