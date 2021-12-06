package services

import (
	"arriba/internal/domain"
	"arriba/internal/domain/constants"
	"errors"
	"time"
)

func AddMovement(ctx *domain.ArribaContext, userId int64, amount int64,
	assetID constants.AssetID, transactionType constants.TransactionType) ([]domain.Balance, error) {

	if !assetID.IsValid() {
		return nil, errors.New(string(constants.CurrencyInvalid))
	}

	user := GetUser(ctx, userId)
	account := user.Account

	switch transactionType {
	case constants.Withdraw, constants.Sell, constants.Debit:
		total := GetTotalAsset(user.Account, assetID)

		if total < amount {
			return nil, errors.New(string(constants.InsufficientFounds))
		}

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

	UpdateUserAccount(ctx, userId, account)

	return account.Balance, nil
}
