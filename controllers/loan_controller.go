package controllers

import (
	"net/http"
	"strconv"
	"time"

	"banking_system/models"
	"banking_system/services"

	"github.com/gin-gonic/gin"
)

type LoanController struct {
	service *services.LoanService
}

func NewLoanController(service *services.LoanService) *LoanController {
	return &LoanController{service: service}
}

func (c *LoanController) CreateLoan(ctx *gin.Context) {
	var loan models.Loan
	if err := ctx.ShouldBindJSON(&loan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.Create(&loan); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, loan)
}

func (c *LoanController) GetLoanByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan id"})
		return
	}

	loan, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "loan not found"})
		return
	}

	ctx.JSON(http.StatusOK, loan)
}

func (c *LoanController) GetAllLoans(ctx *gin.Context) {
	loans, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, loans)
}

func (c *LoanController) UpdateLoan(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan id"})
		return
	}

	loan, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "loan not found"})
		return
	}

	if err := ctx.ShouldBindJSON(loan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.Update(loan); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, loan)
}

func (c *LoanController) DeleteLoan(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan id"})
		return
	}

	if err := c.service.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *LoanController) GetLoanDetails(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan id"})
		return
	}

	details, err := c.service.GetDetails(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, details)
}

type RepayRequest struct {
	Amount      float64 `json:"amount"`
	PaymentDate string  `json:"payment_date"`
}

func (c *LoanController) RepayLoan(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan id"})
		return
	}

	var req RepayRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var paymentDate time.Time
	if req.PaymentDate == "" {
		paymentDate = time.Now()
	} else {
		paymentDate, err = time.Parse(time.RFC3339, req.PaymentDate)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid payment_date, must be RFC3339"})
			return
		}
	}

	repayment, err := c.service.Repay(uint(id), req.Amount, paymentDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, repayment)
}

