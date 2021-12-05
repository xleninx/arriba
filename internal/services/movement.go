package services

import (
	"arriba/internal/domain"
	"arriba/internal/domain/constants"
	"time"
)

func AddMovement(ctx domain.ArribaContext, userId int64, amount int64,
	assetID constants.AssetID, transactionType constants.TransactionType) []domain.Balance {
	account := ctx.Users[userId].Account

	switch transactionType {
	case constants.Withdraw, constants.Sell, constants.Debit:
		amount = -amount
	}

	transaction := domain.Transaction{
		TransactionDate: time.Now(),
		TransactionType: transactionType,
		Asset: domain.Asset{
			ID:     assetID,
			Price:  ctx.AssetsProvider[assetID].Price,
			Amount: amount,
		},
	}

	if account.Movements == nil {
		account.Movements = make(map[constants.AssetID][]domain.Transaction)
	}
	account.Movements[assetID] = append(account.Movements[assetID], transaction)
	account.Balance = CalculateBalance(account)

	if a, ok := ctx.Users[userId]; ok {
		a.Account = account
		ctx.Users[userId] = a
	}

	return account.Balance
}
