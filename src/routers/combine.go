package routers

import "github.com/gin-gonic/gin"

func Combine(r *gin.Engine) {
	UserRouter(r.Group("/users"))
}
