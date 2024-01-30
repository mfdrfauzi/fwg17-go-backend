package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mfdrfauzi/fwg17-go-backend/src/routers"
)

func main() {
	r := gin.Default()
	routers.Combine(r)

	r.Run()
}
