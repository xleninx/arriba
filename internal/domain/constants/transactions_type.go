package constants

type TransactionType string

const (
	Deposit  TransactionType = "Deposit"
	Withdraw TransactionType = "Withdraw"
	Debit    TransactionType = "Debit"
	Buy      TransactionType = "Buy"
	Sell     TransactionType = "Sell"
)
