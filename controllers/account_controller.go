package controllers

import (
	"net/http"
	"strconv"

	"banking_system/models"
	"banking_system/services"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	service *services.AccountService
}

func NewAccountController(service *services.AccountService) *AccountController {
	return &AccountController{service: service}
}

func (c *AccountController) CreateAccount(ctx *gin.Context) {
	var account models.Account
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.Create(&account); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, account)
}

func (c *AccountController) GetAccountByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid account id"})
		return
	}

	accountDetail, err := c.service.GetAccountDetail(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}

	ctx.JSON(http.StatusOK, accountDetail)
}

func (c *AccountController) GetAllAccounts(ctx *gin.Context) {
	accounts, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}

func (c *AccountController) UpdateAccount(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid account id"})
		return
	}

	account, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}

	if err := ctx.ShouldBindJSON(account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.Update(account); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (c *AccountController) DeleteAccount(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid account id"})
		return
	}

	if err := c.service.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *AccountController) AddCustomerToAccount(ctx *gin.Context) {
	accountID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid account id"})
		return
	}

	customerID, err := strconv.Atoi(ctx.Param("customerId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer id"})
		return
	}

	accountDetail, err := c.service.AddCustomer(uint(accountID), uint(customerID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, accountDetail)
}

func (c *AccountController) RemoveCustomerFromAccount(ctx *gin.Context) {
	accountID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid account id"})
		return
	}

	customerID, err := strconv.Atoi(ctx.Param("customerId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer id"})
		return
	}

	if err := c.service.RemoveCustomer(uint(accountID), uint(customerID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *AccountController) GetAccountTransactions(ctx *gin.Context) {
	accountID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid account id"})
		return
	}

	txs, err := c.service.GetTransactions(uint(accountID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, txs)
}

type DepositRequest struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

func (c *AccountController) Deposit(ctx *gin.Context) {
	accountID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid account id"})
		return
	}

	var req DepositRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	txRecord, err := c.service.Deposit(uint(accountID), req.Amount, req.Description)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, txRecord)
}

func (c *AccountController) Withdraw(ctx *gin.Context) {
	accountID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid account id"})
		return
	}

	var req DepositRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	txRecord, err := c.service.Withdraw(uint(accountID), req.Amount, req.Description)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, txRecord)
}

