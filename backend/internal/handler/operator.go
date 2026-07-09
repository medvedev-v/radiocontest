package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/medvedev-v/radiocontest/internal/model"
	"github.com/medvedev-v/radiocontest/internal/repository"
)

type OperatorHandler struct {
	repo *repository.OperatorRepository
}

func NewOperatorHandler(repo *repository.OperatorRepository) *OperatorHandler {
	return &OperatorHandler{repo: repo}
}

func (h *OperatorHandler) CreateOperator(c *gin.Context) {
    var operator model.Operator
    if err := c.ShouldBindJSON(&operator); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request body: " + err.Error(),
        })
        return
    }

    id, err := h.repo.Create(operator)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *OperatorHandler) GetOperator(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "error": "Invalid user ID",
        })
        return
    }

    user, err := h.repo.GetByID(id)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, user)
}

func (h *OperatorHandler) UpdateOperator(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "error": "Invalid operator ID",
        })
        return
    }

    var operator model.Operator
    if err := c.ShouldBindJSON(&operator); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request body: " + err.Error(),
        })
        return
    }
    operator.ID = id

    if err := h.repo.Update(operator); err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.Status(http.StatusOK)
}

func (h *OperatorHandler) DeleteOperator(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "error": "Invalid user ID",
        })
        return
    }

    if err := h.repo.Delete(id); err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.Status(http.StatusNoContent)
}
