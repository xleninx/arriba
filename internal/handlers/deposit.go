package handlers

import (
	"arriba/internal/domain"
	"arriba/internal/domain/constants"
	"arriba/internal/domain/external"
	"arriba/internal/services"

	"github.com/labstack/echo/v4"
)

func Deposit(c echo.Context, ctx domain.ArribaContext) error {
	u := new(external.FiatRequest)
	if err := c.Bind(u); err != nil {
		return err
	}

	balance := services.AddMovement(ctx, u.UserID, u.Amount, constants.USD, constants.Deposit)
	return c.JSON(200, balance)
}
