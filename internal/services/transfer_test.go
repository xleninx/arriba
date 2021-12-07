package services

import (
	"arriba/internal/domain"
	"arriba/internal/domain/constants"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransfer(t *testing.T) {
	type args struct {
		context    domain.ArribaContext
		fromUserId int64
		toUserId   int64
		assetId    constants.AssetID
		amount     int64
	}
	tests := []struct {
		name    string
		args    args
		errMsg  constants.ClientError
		wantErr bool
	}{
		{
			name: "Make a success transfer",
			args: args{
				context:    buildArribaContext(),
				fromUserId: 2,
				toUserId:   1,
				assetId:    constants.USD,
				amount:     100,
			},
		},
		{
			name: "Invalid Destination",
			args: args{
				context:    buildArribaContext(),
				fromUserId: 2,
				toUserId:   2,
				assetId:    constants.USD,
				amount:     100,
			},
			wantErr: true,
			errMsg:  constants.InvalidDestinationUser,
		},
		{
			name: "Insufficient founds",
			args: args{
				context:    buildArribaContext(),
				fromUserId: 2,
				toUserId:   1,
				assetId:    constants.USD,
				amount:     10000,
			},
			wantErr: true,
			errMsg:  constants.InsufficientFounds,
		},
		{
			name: "Invalid currency",
			args: args{
				context:    buildArribaContext(),
				fromUserId: 2,
				toUserId:   1,
				assetId:    "ADA",
				amount:     100,
			},
			wantErr: true,
			errMsg:  constants.CurrencyInvalid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Transfer(
				&tt.args.context,
				tt.args.fromUserId,
				tt.args.toUserId,
				tt.args.amount,
				tt.args.assetId,
			)
			if tt.wantErr {
				assert.Equal(t, err.Error(), string(tt.errMsg))
			} else {
				assert.Nil(t, err)
			}

		})
	}
}
