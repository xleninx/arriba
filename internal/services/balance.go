package services

import (
	"arriba/internal/domain"
	"arriba/internal/domain/constants"
)

func CalculateBalance(account domain.Account) []domain.Balance {
	var assets []domain.Balance
	for key, transactions := range account.Movements {
		var total int64 = 0
		for _, transaction := range transactions {
			total = total + transaction.Asset.Amount
		}
		assets = append(assets, domain.Balance{
			AssetID: key,
			Total:   total,
		})
	}

	return assets
}

func GetTotalAsset(account domain.Account, assetId constants.AssetID) int64 {
	balances := CalculateBalance(account)
	for _, asset := range balances {
		if asset.AssetID == assetId {
			return asset.Total
		}
	}
	return 0
}
