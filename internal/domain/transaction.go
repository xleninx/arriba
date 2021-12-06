package domain

import (
	"arriba/internal/domain/constants"
	"time"
)

type Transaction struct {
	TransactionDate time.Time
	TransactionType constants.TransactionType
	FromUser        int64
	ToUser          int64
	Asset           Asset
}
