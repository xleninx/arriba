package main

import (
	"arriba/internal/app"
	"arriba/internal/domain"
	"arriba/internal/domain/constants"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	users := map[int64]domain.User{
		1: {Id: 1, Account: domain.Account{}},
		2: {Id: 2, Account: domain.Account{}},
		3: {Id: 3, Account: domain.Account{}},
	}
	assetsProvider := map[constants.AssetID]domain.Asset{
		constants.ETH: {
			ID:    constants.ETH,
			Name:  "Ether",
			Price: 450231,
		},
		constants.BTC: {
			ID:    constants.BTC,
			Name:  "Bitcoin",
			Price: 5650020,
		},
		constants.USD: {
			ID:    constants.USD,
			Name:  "Dollar",
			Price: 1,
		},
	}
	ctx := domain.ArribaContext{
		Users:          users,
		AssetsProvider: assetsProvider,
	}

	e := echo.New()
	e.Use(app.IdempotentCheck(&ctx.Cache))

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody []byte, resBody []byte) {
		app.SetIdempotentValue(c, &ctx, resBody)
	}))

	app.BuildRoutes(e, &ctx)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
