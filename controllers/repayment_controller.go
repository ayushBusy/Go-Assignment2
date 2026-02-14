package controllers

import (
	"net/http"
	"strconv"

	"banking_system/models"
	"banking_system/services"

	"github.com/gin-gonic/gin"
)

type RepaymentController struct {
	service *services.RepaymentService
}

func NewRepaymentController(service *services.RepaymentService) *RepaymentController {
	return &RepaymentController{service: service}
}

func (c *RepaymentController) CreateRepayment(ctx *gin.Context) {
	var repayment models.Repayment
	if err := ctx.ShouldBindJSON(&repayment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.Create(&repayment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, repayment)
}

func (c *RepaymentController) GetRepaymentByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid repayment id"})
		return
	}

	repayment, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "repayment not found"})
		return
	}

	ctx.JSON(http.StatusOK, repayment)
}

func (c *RepaymentController) GetAllRepayments(ctx *gin.Context) {
	repayments, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, repayments)
}

