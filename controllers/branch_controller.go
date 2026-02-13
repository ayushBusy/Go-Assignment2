package controllers

import (
	"net/http"
	"strconv"

	"banking_system/models"
	"banking_system/services"

	"github.com/gin-gonic/gin"
)

type BranchController struct {
	service *services.BranchService
}

func NewBranchController(service *services.BranchService) *BranchController {
	return &BranchController{service: service}
}

func (c *BranchController) CreateBranch(ctx *gin.Context) {
	var branch models.Branch
	if err := ctx.ShouldBindJSON(&branch); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.Create(&branch); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, branch)
}

func (c *BranchController) GetBranchByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid branch id"})
		return
	}

	branch, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "branch not found"})
		return
	}

	ctx.JSON(http.StatusOK, branch)
}

func (c *BranchController) GetAllBranches(ctx *gin.Context) {
	branches, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, branches)
}

func (c *BranchController) UpdateBranch(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid branch id"})
		return
	}

	branch, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "branch not found"})
		return
	}

	if err := ctx.ShouldBindJSON(branch); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.Update(branch); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, branch)
}

func (c *BranchController) DeleteBranch(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid branch id"})
		return
	}

	if err := c.service.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

