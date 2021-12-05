package app

import (
	"arriba/internal/domain"
	"arriba/internal/handlers"

	"github.com/labstack/echo/v4"
)

func BuildRoutes(e *echo.Echo, ctx domain.ArribaContext) {
	e.GET("/", func(c echo.Context) error {
		return handlers.Balance(c, ctx)
	})
	e.POST("/deposit", func(c echo.Context) error {
		return handlers.Deposit(c, ctx)
	})
	e.POST("/withdraw", func(c echo.Context) error {
		return handlers.Withdraw(c, ctx)
	})

	e.POST("/buy", func(c echo.Context) error {
		return handlers.Buy(c, ctx)
	})
	e.POST("/sell", func(c echo.Context) error {
		return handlers.Sell(c, ctx)
	})
}
