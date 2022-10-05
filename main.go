package main

import (
	"github.com/ainmtsn1999/go-api-get-time/controllers/timecontroller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/time", timecontroller.Index)
	r.Run()

}
