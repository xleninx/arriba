package app

import (
	"arriba/internal/domain"
	"arriba/internal/domain/constants"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

// SetIdempotentValue For now the function set this value without TTL but
// in a production app should be implemented with a cache storage with a TTL
func SetIdempotentValue(c echo.Context, ctx *domain.ArribaContext, resBody []byte) {
	key := c.Request().Header.Get(string(constants.IdempotentKey))
	if c.Request().Method == "POST" && key != "" && c.Response().Status == http.StatusOK {
		var dat interface{}
		if err := json.Unmarshal(resBody, &dat); err != nil {
			println(err.Error())
		}
		key := c.Request().Header.Get(string(constants.IdempotentKey))
		ctx.Cache.Store(key, dat)
	}
}

func IdempotentCheck(store *sync.Map) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			key := c.Request().Header.Get(string(constants.IdempotentKey))
			if c.Request().Method == "POST" {
				if key == "" {
					return echo.NewHTTPError(http.StatusBadRequest, string(constants.InvalidIdempotentKey))
				}

				resp, ok := store.Load(key)
				if ok {
					return c.JSON(http.StatusOK, resp)
				}
			}
			return next(c)
		}
	}
}
