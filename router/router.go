package router

import (
	"github.com/gin-gonic/gin"
	"gogin/collector"
)

func Router(e *gin.Engine) {
	r := e.Group("api")
	{
		r.GET("/", collector.Test)
	}
}
