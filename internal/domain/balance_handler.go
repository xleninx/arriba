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

//func (bh BalanceHandler) deposit(amount int64) error {
//	asset := bh.AssetProvider[constants.USD]
//	asset.Amount = amount
//
//	balance := domain.Transaction{AssetID: asset, TransactionDate: time.Now(), TransactionType: constants.Deposit}
//	bh.User.Account.Movements[constants.USD] = append(bh.User.Account.Movements[constants.USD], balance)
//	return nil
//}
//
//func (bh BalanceHandler) withdraw(amount int64) error {
//	asset := bh.AssetProvider[constants.USD]
//	asset.Amount = amount
//
//	balance := domain.Transaction{AssetID: asset, TransactionDate: time.Now(), TransactionType: constants.Withdraw}
//	bh.User.Account.Movements[constants.USD] = append(bh.User.Account.Movements[constants.USD], balance)
//
//	return nil
//}
