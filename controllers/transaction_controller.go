package controllers

import (
	"net/http"
	"strconv"

	"banking_system/models"
	"banking_system/services"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	service *services.TransactionService
}

func NewTransactionController(service *services.TransactionService) *TransactionController {
	return &TransactionController{service: service}
}

func (c *TransactionController) CreateTransaction(ctx *gin.Context) {
	var txn models.Transaction
	if err := ctx.ShouldBindJSON(&txn); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.Create(&txn); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

func (c *TransactionController) UpdateTransaction(ctx *gin.Context) {
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

	if err := ctx.ShouldBindJSON(txn); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.Update(txn); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, txn)
}

func (c *TransactionController) DeleteTransaction(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid transaction id"})
		return
	}

	if err := c.service.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

