package domain

import "arriba/internal/domain/constants"

type ArribaContext struct {
	Users          map[int64]User
	AssetsProvider map[constants.AssetID]Asset
}
