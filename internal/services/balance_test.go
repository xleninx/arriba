package services

import (
	"arriba/internal/domain"
	"arriba/internal/domain/constants"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCalculateBalance(t *testing.T) {
	tests := []struct {
		name string
		args domain.Account
		want []domain.Balance
	}{
		{
			name: "Calculate one asset",
			args: domain.Account{
				Movements: map[constants.AssetID][]domain.Transaction{
					constants.BTC: {
						{
							TransactionDate: time.Now(),
							TransactionType: constants.Buy,
							FromUser:        0,
							ToUser:          0,
							Asset: domain.Asset{
								ID:     constants.BTC,
								Name:   "BTC",
								Price:  100,
								Amount: 2,
							},
						},
					},
				},
			},
			want: []domain.Balance{
				{
					AssetID: constants.BTC,
					Total:   2,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateBalance(tt.args)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestGetTotalAsset(t *testing.T) {
	type args struct {
		Account domain.Account
		AssetId constants.AssetID
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Calculate total with 2 transactions",
			args: args{
				Account: domain.Account{
					Movements: map[constants.AssetID][]domain.Transaction{
						constants.BTC: {
							{
								TransactionDate: time.Now(),
								TransactionType: constants.Buy,
								FromUser:        0,
								ToUser:          0,
								Asset: domain.Asset{
									ID:     constants.BTC,
									Name:   "BTC",
									Price:  100,
									Amount: 2,
								},
							},
							{
								TransactionDate: time.Now(),
								TransactionType: constants.Buy,
								FromUser:        0,
								ToUser:          0,
								Asset: domain.Asset{
									ID:     constants.BTC,
									Name:   "BTC",
									Price:  110,
									Amount: 2,
								},
							},
						},
					},
				},
				AssetId: constants.BTC,
			},
			want: 4,
		},
		{
			name: "Calculate total without transactions",
			args: args{
				Account: domain.Account{
					Balance:   nil,
					Movements: nil,
				},
				AssetId: constants.BTC,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetTotalAsset(tt.args.Account, tt.args.AssetId)
			assert.Equal(t, got, tt.want)
		})
	}
}
