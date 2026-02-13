package controllers

import (
	"net/http"
	"strconv"

	"banking_system/models"
	"banking_system/services"

	"github.com/gin-gonic/gin"
)

type BankController struct {
	service *services.BankService
}

func NewBankController(service *services.BankService) *BankController {
	return &BankController{service: service}
}

func (c *BankController) CreateBank(ctx *gin.Context) {
	var bank models.Bank
	if err := ctx.ShouldBindJSON(&bank); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if bank.Name == "" || bank.Code == "" || bank.Location == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name, code, and location are required"})
		return
	}

	if err := c.service.Create(&bank); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, bank)
}

func (c *BankController) GetBankByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid bank id"})
		return
	}

	bank, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "bank not found"})
		return
	}

	ctx.JSON(http.StatusOK, bank)
}

func (c *BankController) GetAllBanks(ctx *gin.Context) {
	banks, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, banks)
}

func (c *BankController) UpdateBank(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid bank id"})
		return
	}

	bank, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "bank not found"})
		return
	}

	if err := ctx.ShouldBindJSON(bank); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if bank.Name == "" || bank.Code == "" || bank.Location == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name, code, and location are required"})
		return
	}

	if err := c.service.Update(bank); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, bank)
}

func (c *BankController) DeleteBank(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid bank id"})
		return
	}

	if err := c.service.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

