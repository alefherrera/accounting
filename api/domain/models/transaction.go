package models

type TransactionType string

const (
	TransactionTypeDebit  TransactionType = "debit"
	TransactionTypeCredit TransactionType = "credit"
)

type Transaction struct {
	TransactionType TransactionType
	Amount          float64
}

func NewTransaction(transactionType TransactionType, amount float64) *Transaction {
	return &Transaction{TransactionType: transactionType, Amount: amount}
}
