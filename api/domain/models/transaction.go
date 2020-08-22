package models

import "github.com/google/uuid"

type TransactionType string

const (
	TransactionTypeDebit  TransactionType = "debit"
	TransactionTypeCredit TransactionType = "credit"
)

type Transaction struct {
	Id              string
	TransactionType TransactionType
	Amount          float64
}

func NewTransaction(transactionType TransactionType, amount float64) *Transaction {
	return &Transaction{Id: uuid.New().String(), TransactionType: transactionType, Amount: amount}
}
