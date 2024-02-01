package routers

import (
	"github.com/gin-gonic/gin"
	controller "github.com/mfdrfauzi/fwg17-go-backend/src/controllers"
)

func UserRouter(r *gin.RouterGroup) {
	r.GET("", controller.ListAllUsers)
	r.GET("/:id", controller.DetailUser)
	r.POST("", controller.CreateUser)
	// r.PATCH("/:id", controller.UpdateUser)
	// r.DELETE("/:id", controller.DeleteUser)
}
