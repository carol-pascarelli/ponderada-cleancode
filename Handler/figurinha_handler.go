package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ponderada-cleancode/domain"
	"ponderada-cleancode/service"
)

type FigurinhaHandler struct {
	svc service.FigurinhaService
}

func NewFigurinhaHandler(svc service.FigurinhaService) *FigurinhaHandler {
	return &FigurinhaHandler{svc: svc}
}

func (h *FigurinhaHandler) respond(c *gin.Context, data interface{}, err error) {
	if err == nil {
		if c.Request.Method == http.MethodPost {
			c.JSON(http.StatusCreated, data)
		} else if c.Request.Method == http.MethodDelete {
			c.Status(http.StatusNoContent)
		} else {
			c.JSON(http.StatusOK, data)
		}
		return
	}

	if errors.Is(err, service.ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else if errors.Is(err, service.ErrInvalidTipo) ||
		errors.Is(err, service.ErrInvalidPosicao) ||
		errors.Is(err, service.ErrCamposObrigatorios) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro interno"})
	}
}

func (h *FigurinhaHandler) Create(c *gin.Context) {
	var req domain.CreateFigurinhaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	figurinha, err := h.svc.Create(req)
	h.respond(c, figurinha, err)
}

func (h *FigurinhaHandler) List(c *gin.Context) {
	posicao := c.Query("posicao")
	tipo := c.Query("tipo")
	figurinhas, err := h.svc.List(posicao, tipo)
	h.respond(c, figurinhas, err)
}

func (h *FigurinhaHandler) GetByID(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	figurinha, err := h.svc.GetByID(id)
	h.respond(c, figurinha, err)
}

func (h *FigurinhaHandler) Update(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var req domain.UpdateFigurinhaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	figurinha, err := h.svc.Update(id, req)
	h.respond(c, figurinha, err)
}

func (h *FigurinhaHandler) Delete(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	err = h.svc.Delete(id)
	h.respond(c, nil, err)
}

func parseID(c *gin.Context) (uint, error) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	return uint(id), err
}
