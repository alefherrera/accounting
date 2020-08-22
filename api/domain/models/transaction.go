package models

import "github.com/google/uuid"

type TransactionType string

const (
	TransactionTypeDebit  TransactionType = "debit"
	TransactionTypeCredit TransactionType = "credit"
)

type Transaction struct {
	Id              string          `json:"id"`
	TransactionType TransactionType `json:"transaction_type"`
	Amount          float64         `json:"amount"`
}

func NewTransaction(transactionType TransactionType, amount float64) *Transaction {
	return &Transaction{Id: uuid.New().String(), TransactionType: transactionType, Amount: amount}
}
