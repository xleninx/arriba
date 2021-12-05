package handlers

import (
	"arriba/internal/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Balance(c echo.Context, ctx domain.ArribaContext) error {
	return c.JSON(http.StatusOK, ctx.Users)
}
