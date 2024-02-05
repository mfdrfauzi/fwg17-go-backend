package admin

import (
	"github.com/gin-gonic/gin"
	admin_controllers "github.com/mfdrfauzi/fwg17-go-backend/src/controllers/admin"
)

func UserRouter(r *gin.RouterGroup) {
	r.GET("", admin_controllers.ListAllUsers)
	r.GET("/:id", admin_controllers.DetailUser)
	r.POST("", admin_controllers.CreateUser)
	r.PATCH("/:id", admin_controllers.UpdateUser)
	r.DELETE("/:id", admin_controllers.DeleteUser)
}
