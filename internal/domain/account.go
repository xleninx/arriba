package domain

import (
	"arriba/internal/domain/constants"
)

type Account struct {
	Balance   []Balance                           `json:"balance"`
	Movements map[constants.AssetID][]Transaction `json:"-"`
}
