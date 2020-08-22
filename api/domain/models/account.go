package models

type Account struct {
	Transactions []Transaction
}

func (a Account) GetBalance() float64 {
	var result float64

	for _, transaction := range a.Transactions {
		if transaction.TransactionType == TransactionTypeDebit {
			result -= transaction.Amount
		} else {
			result += transaction.Amount
		}
	}

	return result
}

func (a *Account) CommitTransaction(transaction Transaction) {
	a.Transactions = append(a.Transactions, transaction)
}
