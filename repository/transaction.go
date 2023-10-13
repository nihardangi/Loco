package repository

import "Loco/model"

// Holds all repository functions related to transactions.
// Provides abstraction b/w transaction service and transaction repository functions.
type TransactionRepo interface {
	CreateTransaction(transaction model.Transaction) error
	GetTransaction(transactionID uint) (model.Transaction, error)
	GetAllTransactionIDsByType(transactionType string) ([]uint, error)
	GetSumOfSuccessorTransactions(transactionID uint) (float64, error)
}
