package controllers

import (
	"net/http"

	"github.com/KEINOS/go-argonize"
	"github.com/gin-gonic/gin"
	"github.com/mfdrfauzi/fwg17-go-backend/src/models"
	"github.com/mfdrfauzi/fwg17-go-backend/src/services"
)

func Login(c *gin.Context) {
	form := models.User{}
	err := c.ShouldBind(&form)

	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Invalid",
		})
		return
	}

	user, err := models.GetOneUserByEmail(form.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, &services.ResponseOnly{
			Success: false,
			Message: "Wrong Email or Password",
		})
		return
	}
	decoded, err := argonize.DecodeHashStr(user.Password)

	plainPass := []byte(form.Password)
	if decoded.IsValidPassword(plainPass) {
		c.JSON(http.StatusOK, &services.ResponseOnly{
			Success: false,
			Message: "Wrong Email or Password",
		})
		return
	} else {
		c.JSON(http.StatusUnauthorized, &services.ResponseOnly{
			Success: false,
			Message: "Wrong Email or Password",
		})
		return
	}
}
