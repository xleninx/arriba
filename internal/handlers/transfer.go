package handlers

import (
	"arriba/internal/domain"
	"arriba/internal/domain/external"
	"arriba/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Transfer(c echo.Context, ctx *domain.ArribaContext) error {
	request := new(external.TransferRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	err := services.Transfer(
		ctx,
		request.FromUserID,
		request.ToUserID,
		request.Amount,
		request.Currency,
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, ctx.Users)

}
