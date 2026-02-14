package controllers

import (
	"net/http"
	"strconv"

	"banking_system/models"
	"banking_system/services"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	service *services.CustomerService
}

func NewCustomerController(service *services.CustomerService) *CustomerController {
	return &CustomerController{service: service}
}

func (c *CustomerController) CreateCustomer(ctx *gin.Context) {
	var customer models.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if customer.FirstName == "" || customer.LastName == "" || customer.Email == "" || customer.Phone == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "first_name, last_name, email, and phone_number are required"})
		return
	}

	if err := c.service.Create(&customer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, customer)
}

func (c *CustomerController) GetCustomerByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer id"})
		return
	}

	customer, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

func (c *CustomerController) GetAllCustomers(ctx *gin.Context) {
	customers, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customers)
}

func (c *CustomerController) UpdateCustomer(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer id"})
		return
	}

	customer, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
		return
	}

	if err := ctx.ShouldBindJSON(customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if customer.FirstName == "" || customer.LastName == "" || customer.Email == "" || customer.Phone == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "first_name, last_name, email, and phone_number are required"})
		return
	}

	if err := c.service.Update(customer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

func (c *CustomerController) DeleteCustomer(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer id"})
		return
	}

	if err := c.service.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
