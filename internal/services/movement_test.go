package services

import (
	"arriba/internal/domain"
	"arriba/internal/domain/constants"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddMovement(t *testing.T) {
	type args struct {
		context         domain.ArribaContext
		userId          int64
		transactionType constants.TransactionType
		assetId         constants.AssetID
		amount          int64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		errMsg  constants.ClientError
		wantErr bool
	}{
		{
			name: "Add a deposit",
			args: args{
				context:         buildArribaContext(),
				userId:          1,
				transactionType: constants.Deposit,
				assetId:         constants.USD,
				amount:          1000,
			},
			want:    1000,
			wantErr: false,
		},
		{
			name: "Error sell without founds",
			args: args{
				context:         buildArribaContext(),
				userId:          1,
				transactionType: constants.Sell,
				assetId:         constants.BTC,
				amount:          1,
			},
			errMsg:  constants.InsufficientFounds,
			wantErr: true,
		},
		{
			name: "Error invalid currency",
			args: args{
				context:         buildArribaContext(),
				userId:          1,
				transactionType: constants.Buy,
				assetId:         "AVAX",
				amount:          1,
			},
			errMsg:  constants.CurrencyInvalid,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AddMovement(
				&tt.args.context,
				tt.args.userId,
				tt.args.amount,
				tt.args.assetId,
				tt.args.transactionType,
			)
			if tt.wantErr {
				assert.Equal(t, err.Error(), string(tt.errMsg))
			}
			for _, asset := range got {
				if asset.AssetID == tt.args.assetId {
					assert.Equal(t, asset.Total, tt.want)
				}
			}
		})
	}
}

func buildArribaContext() domain.ArribaContext {
	return domain.ArribaContext{
		Users: map[int64]domain.User{
			1: {
				Id:      1,
				Account: domain.Account{},
			},
			2: {
				Id: 2,
				Account: domain.Account{Movements: map[constants.AssetID][]domain.Transaction{
					constants.USD: {
						{
							TransactionDate: time.Time{},
							TransactionType: constants.Deposit,
							FromUser:        0,
							ToUser:          0,
							Asset: domain.Asset{
								ID:     constants.USD,
								Name:   "",
								Amount: 1000,
							},
						},
					},
				}},
			},
		},
		AssetsProvider: map[constants.AssetID]domain.Asset{
			constants.BTC: {
				ID:    constants.BTC,
				Name:  "Bitcoin",
				Price: 100,
			},
			constants.ETH: {
				ID:    constants.ETH,
				Name:  "Ether",
				Price: 10,
			},
		},
	}
}
