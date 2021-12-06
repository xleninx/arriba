package constants

type ClientError string

const (
	InsufficientFounds     ClientError = "you do not have sufficient funds for the operation"
	CurrencyInvalid        ClientError = "currency is invalid for the operation"
	InvalidDestinationUser ClientError = "the destination user id is invalid"
)
