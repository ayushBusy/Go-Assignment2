package controllers

import (
	"net/http"
	"strconv"

	"banking_system/services"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	service *services.TransactionService
}

func NewTransactionController(service *services.TransactionService) *TransactionController {
	return &TransactionController{service: service}
}

type CreateTransactionRequest struct {
	AccountID   uint   `json:"account_id" binding:"required"`
	Type        string `json:"transaction_type" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
	Description string `json:"description"`
}

func (c *TransactionController) CreateTransaction(ctx *gin.Context) {
	var req CreateTransactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	txn, err := c.service.ProcessTransaction(req.AccountID, req.Type, req.Amount, req.Description)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, txn)
}

func (c *TransactionController) GetTransactionByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid transaction id"})
		return
	}

	txn, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "transaction not found"})
		return
	}

	ctx.JSON(http.StatusOK, txn)
}

func (c *TransactionController) GetAllTransactions(ctx *gin.Context) {
	txs, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, txs)
}

func (c *TransactionController) GetAccountTransactions(ctx *gin.Context) {
	accountID, err := strconv.Atoi(ctx.Param("accountId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid account id"})
		return
	}

	txs, err := c.service.GetAccountTransactions(uint(accountID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, txs)
}

