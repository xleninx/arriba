package domain

import "arriba/internal/domain/constants"

type Asset struct {
	ID     constants.AssetID
	Name   string
	Price  int64
	Amount int64
}
