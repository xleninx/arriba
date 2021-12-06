package services

import (
	"arriba/internal/domain"
	"arriba/internal/domain/constants"
	"errors"
	"time"
)

func Transfer(ctx *domain.ArribaContext, FromUserID int64, ToUserID int64,
	amount int64, assetID constants.AssetID) error {

	if !assetID.IsValid() {
		return errors.New(string(constants.CurrencyInvalid))
	}

	if FromUserID == ToUserID {
		return errors.New(string(constants.InvalidDestinationUser))
	}

	usersIds := []int64{
		FromUserID,
		ToUserID,
	}
	for _, userID := range usersIds {
		user := GetUser(ctx, userID)
		account := user.Account

		transaction := domain.Transaction{
			TransactionDate: time.Now(),
			TransactionType: constants.Transfer,
			FromUser:        FromUserID,
			ToUser:          ToUserID,
			Asset: domain.Asset{
				ID:     assetID,
				Price:  ctx.AssetsProvider[assetID].Price,
				Amount: amount,
			},
		}

		if userID == FromUserID {
			total := GetTotalAsset(user.Account, assetID)
			if total < amount {
				return errors.New(string(constants.InsufficientFounds))
			}

			transaction.Asset.Amount = -amount
		}

		if account.Movements == nil {
			account.Movements = make(map[constants.AssetID][]domain.Transaction)
		}

		account.Movements[assetID] = append(account.Movements[assetID], transaction)
		account.Balance = CalculateBalance(account)

		UpdateUserAccount(ctx, userID, account)
	}

	return nil
}
