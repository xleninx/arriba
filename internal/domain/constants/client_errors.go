package constants

type ClientError string

const (
	InsufficientFounds ClientError = "You do not have sufficient funds for the operation"
	CurrencyInvalid    ClientError = "Currency is invalid for the operation"
)
