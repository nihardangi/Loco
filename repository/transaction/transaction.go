package transaction

import (
	"Loco/model"

	"gorm.io/gorm"
)

type TransactionOps struct {
	db *gorm.DB
}

func NewTransactionOps(db *gorm.DB) TransactionOps {
	return TransactionOps{db: db}
}

// CreateTransaction adds a new transaction object in DB.
func (t TransactionOps) CreateTransaction(transaction model.Transaction) error {
	if err := t.db.Create(&transaction).Error; err != nil {
		return err
	}
	return nil
}

// GetTransaction fetches details of a transaction.
func (t TransactionOps) GetTransaction(transactionID uint) (model.Transaction, error) {
	transaction := model.Transaction{ID: transactionID}
	if err := t.db.First(&transaction).Error; err != nil {
		return model.Transaction{}, err
	}
	return transaction, nil
}

// GetAllTransactionIDsByType fetches the transactionIDs of all transactions belonging to a particular type.
func (t TransactionOps) GetAllTransactionIDsByType(transactionType string) ([]uint, error) {
	var transactionIDs []uint
	if err := t.db.Model(&model.Transaction{}).Where("type = ?", transactionType).Pluck("id", &transactionIDs).Error; err != nil {
		return []uint{}, err
	}
	return transactionIDs, nil
}

// GetSumOfSuccessorTransactions gets the sum of all transactions that are children/successor of input transactionID
func (t TransactionOps) GetSumOfSuccessorTransactions(transactionID uint) (float64, error) {
	transaction, err := t.GetTransaction(transactionID)
	if err != nil {
		return 0, err
	}

	var totalSum float64
	t.db.Raw(
		`WITH RECURSIVE transaction_sum AS (
			SELECT
				t.id,
				t.amount
			FROM
				transactions t
			WHERE
				id = ?
			UNION ALL
			SELECT
				t.id,
				t.amount
			FROM
				transaction_sum ts
				JOIN transactions t ON ts.id = t.parent_id
		)
		SELECT
			SUM(amount)
		FROM
			transaction_sum;`, transactionID).Scan(&totalSum)

	// At this stage, totalSum holds base transaction amount + amount of all child transactions.
	// Subtracting base transaction's amount as we only need sum of child transactions.
	totalSum -= transaction.Amount
	return totalSum, nil
}
