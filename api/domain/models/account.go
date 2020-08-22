package models

type Account struct {
	Transactions []Transaction
}

func (a Account) GetBalance() float64 {
	return 0
}

func (a *Account) CommitTransaction(transaction Transaction) {
	a.Transactions = append(a.Transactions, transaction)
}
