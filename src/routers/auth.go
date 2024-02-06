package routers

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/mfdrfauzi/fwg17-go-backend/src/controllers/auth"
)

func AuthRouter(r *gin.RouterGroup) {
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)
	r.POST("/forgot-password", controllers.ForgotPassword)
}
