package domain

import (
	"arriba/internal/domain/constants"
)

type BalanceHandler struct {
	User          User
	AssetProvider map[constants.AssetID]Asset
}
