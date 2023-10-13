package service

import (
	"Loco/errors"
	"Loco/model"
	"Loco/repository"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type TransactionService struct {
	transactionRepo repository.TransactionRepo
	db              *gorm.DB
}

func NewTransactionService(transactionRepo repository.TransactionRepo, db *gorm.DB) TransactionService {
	return TransactionService{transactionRepo: transactionRepo, db: db}
}

// Service function to create a new transaction.
// Throws 409 conflict status code if transactionID already exists.
// In case of other errors, throws 500 internal server error
// Return 200 OK on successful transaction creation.
func (t TransactionService) CreateTransaction(transaction model.Transaction) (int, error) {
	if err := t.transactionRepo.CreateTransaction(transaction); err != nil {
		if errors.IsDuplicateKeyGormError(err) {
			return http.StatusConflict, fmt.Errorf("transactionId: %d already exists in DB", transaction.ID)
		}
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

// Fetch details of a transaction
// Throws 404 when transactionID is not found
func (t TransactionService) GetTransaction(transactionID uint) (model.Transaction, int, error) {
	transaction, err := t.transactionRepo.GetTransaction(transactionID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.Transaction{}, http.StatusNotFound, fmt.Errorf("transaction_id: %d not found", transactionID)
		}
		return model.Transaction{}, http.StatusInternalServerError, fmt.Errorf("error in fetching transaction details, err: %v", err)
	}
	return transaction, http.StatusOK, nil
}

// Fetch all transactionIDs belonging to a type.
// Throws 404 when no transactionIDs found
func (t TransactionService) GetAllTransactionIDsByType(transactionType string) ([]uint, int, error) {
	transactionIDs, err := t.transactionRepo.GetAllTransactionIDsByType(transactionType)
	if err != nil {
		return []uint{}, http.StatusInternalServerError, err
	}
	if len(transactionIDs) == 0 {
		return []uint{}, http.StatusNotFound, fmt.Errorf("no transactions found")
	}
	return transactionIDs, http.StatusOK, nil
}

// Returns sum of all child/successor transactions for a transactionID.
// // Throws 404 when transactionID does not exist.
func (t TransactionService) GetSumOfSuccessorTransactions(transactionID uint) (float64, int, error) {
	sum, err := t.transactionRepo.GetSumOfSuccessorTransactions(transactionID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, http.StatusNotFound, fmt.Errorf("transaction_id: %d not found", transactionID)
		}
	}
	return sum, http.StatusOK, nil
}
