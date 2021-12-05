package domain

import "arriba/internal/domain/constants"

type Balance struct {
	AssetID constants.AssetID
	Total   int64
}
