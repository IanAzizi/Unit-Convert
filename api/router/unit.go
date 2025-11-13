package router

import (
	"unit-convert/api/handler"

	"github.com/gin-gonic/gin"
)

func UnitRouter(r *gin.Engine) {
	h := handler.NewUnitHandler()
	api := r.Group("/api")
	{
		api.POST("/convert", h.ConvertUnit)
	}
}
