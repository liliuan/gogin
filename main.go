package main

import (
	"github.com/gin-gonic/gin"
	"gogin/router"
)

func main() {
	e := gin.Default()

	router.Router(e)

	e.Run(":3000")
}
