package external

import "arriba/internal/domain/constants"

type BuyRequest struct {
	UserID   int64             `json:"user_id"`
	Currency constants.AssetID `json:"currency"`
	Amount   int64             `json:"amount"`
}
