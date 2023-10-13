package controller

import (
	"Loco/errors"
	"Loco/model"
	"Loco/repository/transaction"
	"Loco/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// NewTransactionController initialises transaction repository and service and then initialises API routes for transaction model.
func NewTransactionController(db *gorm.DB, r *gin.RouterGroup) {
	transactionOps := transaction.NewTransactionOps(db)
	transactionService := service.NewTransactionService(transactionOps, db)
	tc := TransactionController{transactionService: transactionService}
	transactionServiceRoute := r.Group("/transactionservice")
	{
		transactionServiceRoute.PUT("/transaction/:transactionId", validateTransactionIDFromPath, tc.CreateUpdateTransaction)
		transactionServiceRoute.GET("/transaction/:transactionId", validateTransactionIDFromPath, tc.GetTransaction)
		transactionServiceRoute.GET("/type/:type", tc.GetTransactionIDsByType)
		transactionServiceRoute.GET("/sum/:transactionId", validateTransactionIDFromPath, tc.GetSum)
	}
}

type TransactionController struct {
	transactionService service.TransactionService
}

// createUpdateTransactionRequest contains the create/update transaction request struct
type createUpdateTransactionRequest struct {
	Amount   float64 `json:"amount" binding:"required,min=1"`
	Type     string  `json:"type" binding:"required"`
	ParentID uint    `json:"parent_id"`
}

// Controller to validate createTransaction request and then call service function
func (t TransactionController) CreateUpdateTransaction(c *gin.Context) {
	// Validate the request payload
	var reqData createUpdateTransactionRequest
	if err := c.ShouldBindJSON(&reqData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": errors.FormatErrors(err)})
		return
	}
	transactionID := c.GetUint("transaction_id")
	transaction := model.Transaction{
		ID:       transactionID,
		Amount:   reqData.Amount,
		Type:     reqData.Type,
		ParentID: reqData.ParentID,
	}
	statusCode, err := t.transactionService.CreateTransaction(transaction)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{"err": err.Error()})
		return
	}
	c.JSON(statusCode, gin.H{"status": "ok"})
}

// Controller to get details of a transaction
func (t TransactionController) GetTransaction(c *gin.Context) {
	transactionID := c.GetUint("transaction_id")
	transaction, statusCode, err := t.transactionService.GetTransaction(transactionID)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{"err": err.Error()})
		return
	}
	c.JSON(statusCode, gin.H{"transaction": transaction})
}

// Controller to get all transactions belonging to a type
func (t TransactionController) GetTransactionIDsByType(c *gin.Context) {
	transactionType := c.Param("type")
	transactionIDs, statusCode, err := t.transactionService.GetAllTransactionIDsByType(transactionType)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{"err": err.Error()})
		return
	}
	c.JSON(statusCode, gin.H{"transactionIDs": transactionIDs})
}

// Controller to get sum of all transactions that are child/successor of a transaction.
func (t TransactionController) GetSum(c *gin.Context) {
	transactionID := c.GetUint("transaction_id")
	sum, statusCode, err := t.transactionService.GetSumOfSuccessorTransactions(transactionID)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{"err": err.Error()})
		return
	}
	c.JSON(statusCode, gin.H{"sum": sum})
}
