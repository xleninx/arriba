package domain

import (
	"arriba/internal/domain/constants"
)

type BalanceHandler struct {
	User          User
	AssetProvider map[constants.AssetID]Asset
}

func NewBalanceHandler(user User, assetProvider map[constants.AssetID]Asset) BalanceHandler {
	return BalanceHandler{
		User:          user,
		AssetProvider: assetProvider,
	}
}
