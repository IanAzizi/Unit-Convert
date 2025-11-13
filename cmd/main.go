package main

import (
	"unit-convert/api/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.UnitRouter(r)
	r.Run(":8080")
}
