package external

import "arriba/internal/domain/constants"

type TransferRequest struct {
	FromUserID int64             `json:"from_user_id"`
	ToUserID   int64             `json:"to_user_id"`
	Currency   constants.AssetID `json:"currency"`
	Amount     int64             `json:"amount"`
}
