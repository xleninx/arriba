package domain

import "arriba/internal/domain/constants"

type Account struct {
	Balance   []Balance
	Movements map[constants.AssetID][]Transaction
}
