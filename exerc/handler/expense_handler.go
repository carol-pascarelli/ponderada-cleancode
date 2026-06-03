package handler

import (
	"errors"
	"net/http"
	"ponderada3/domain"
	"ponderada3/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExpenseHandler struct {
	svc service.ExpenseService
}

func NewExpenseHandler(svc service.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{svc: svc}
}

func (h *ExpenseHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/expenses", h.create)
	r.GET("/expenses", h.list)
	r.GET("/expenses/:id", h.getByID)
	r.PATCH("/expenses/:id", h.update)
	r.DELETE("/expenses/:id", h.delete)
}

func (h *ExpenseHandler) create(c *gin.Context) {
	var req domain.CreateExpenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	expense, err := h.svc.Create(req)
	if err != nil {
		c.JSON(mapError(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, expense)
}

func (h *ExpenseHandler) list(c *gin.Context) {
	category := c.Query("category")
	expenses, err := h.svc.List(category)
	if err != nil {
		c.JSON(mapError(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, expenses)
}

func (h *ExpenseHandler) getByID(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}
	expense, err := h.svc.GetByID(id)
	if err != nil {
		c.JSON(mapError(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, expense)
}

func (h *ExpenseHandler) update(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}
	var req domain.UpdateExpenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	expense, err := h.svc.Update(id, req)
	if err != nil {
		c.JSON(mapError(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, expense)
}

func (h *ExpenseHandler) delete(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}
	if err := h.svc.Delete(id); err != nil {
		c.JSON(mapError(err), gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func parseID(c *gin.Context) (uint, error) {
	n, err := strconv.ParseUint(c.Param("id"), 10, 64)
	return uint(n), err
}

func mapError(err error) int {
	switch {
	case errors.Is(err, service.ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, service.ErrInvalidAmount), errors.Is(err, service.ErrInvalidCategory):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
