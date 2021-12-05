package handlers

import (
	"arriba/internal/domain"
	"arriba/internal/domain/constants"
	"arriba/internal/domain/external"
	"arriba/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Withdraw(c echo.Context, ctx domain.ArribaContext) error {
	u := new(external.FiatRequest)
	if err := c.Bind(u); err != nil {
		return err
	}

	account := ctx.Users[u.UserID].Account
	total := services.GetTotalAsset(account, constants.USD)

	if total < u.Amount {
		return echo.NewHTTPError(http.StatusBadRequest, string(constants.InsufficientFounds))
	}

	balance := services.AddMovement(ctx, u.UserID, u.Amount, constants.USD, constants.Withdraw)
	return c.JSON(200, balance)
}
