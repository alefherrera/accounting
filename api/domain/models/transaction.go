package models

type TransactionType string

const (
	TransactionTypeDebit  TransactionType = "debit"
	TransactionTypeCredit TransactionType = "credit"
)

type Transaction struct {
	TransactionType TransactionType
}
