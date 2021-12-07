package constants

type ClientError string

const (
	InsufficientFounds     ClientError = "you do not have sufficient funds for the operation"
	CurrencyInvalid        ClientError = "currency is invalid for the operation"
	InvalidDestinationUser ClientError = "the destination user id is invalid"
	InvalidIdempotentKey   ClientError = "you must include idempotent_key on the headers and should be unique"
)
