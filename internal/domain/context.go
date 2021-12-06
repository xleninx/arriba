package domain

import (
	"arriba/internal/domain/constants"
	"sync"
)

type ArribaContext struct {
	Users          map[int64]User              `json:"users"`
	AssetsProvider map[constants.AssetID]Asset `json:"assets_provider"`
	Mutex          sync.RWMutex                `json:"-"`
}
