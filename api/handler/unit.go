package handler

import (
	"net/http"
	"unit-convert/service"

	"github.com/gin-gonic/gin"
)

type UnitHandler struct{}

func NewUnitHandler() *UnitHandler {
	return &UnitHandler{}
}

func (u *UnitHandler) ConvertUnit(c *gin.Context) {
	var req service.Unit
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := service.Convert(req)

	c.JSON(http.StatusOK, gin.H{
		"from":   req.From,
		"to":     req.To,
		"value":  req.Value,
		"result": result,
	})
}
