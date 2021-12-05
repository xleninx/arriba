package domain

import (
	"arriba/internal/domain/constants"
	"time"
)

type Transaction struct {
	TransactionDate time.Time
	TransactionType constants.TransactionType
	Asset           Asset
}
