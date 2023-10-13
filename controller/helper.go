package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Validates transactionID received in path parameter. Throws error if type is not int
// or when transactionID < 0
func validateTransactionIDFromPath(c *gin.Context) {
	transactionIDStr := c.Param("transactionId")
	transactionID, err := strconv.Atoi(transactionIDStr)
	if err != nil {
		err = fmt.Errorf("only integer allowed for transaction_id, received: %s", transactionIDStr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	} else if transactionID <= 0 {
		// ID cannot be negative
		err = fmt.Errorf("negative values not allowed for category_id, received: %s", transactionIDStr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.Set("transaction_id", uint(transactionID))
	c.Next()
}
