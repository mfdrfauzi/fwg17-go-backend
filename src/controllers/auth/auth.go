package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/KEINOS/go-argonize"
	"github.com/gin-gonic/gin"
	"github.com/mfdrfauzi/fwg17-go-backend/src/lib"
	"github.com/mfdrfauzi/fwg17-go-backend/src/models"
	"github.com/mfdrfauzi/fwg17-go-backend/src/services"
)

type ResetForm struct {
	Email           string `json:"email" form:"email"`
	Otp             string `json:"otp" form:"otp"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirmPassword" form:"confirmPassword"`
}

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
	decoded, _ := argonize.DecodeHashStr(user.Password)

	plainPass := []byte(form.Password)
	if decoded.IsValidPassword(plainPass) {
		c.JSON(http.StatusOK, &services.ResponseOnly{
			Success: true,
			Message: "Login Success",
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

func Register(c *gin.Context) {
	form := models.User{}

	err := c.ShouldBind(&form)

	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Invalid",
		})
		return
	}

	plainPass := []byte(form.Password)
	hashedPass, err := argonize.Hash(plainPass)
	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Failed to generate Hash",
		})
	}

	form.Password = hashedPass.String()

	newUser, err := models.CreateUser(form)
	if err != nil {
		if strings.HasPrefix(err.Error(), "pq: duplicate key value") {
			c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
				Success: false,
				Message: "Email already exists",
			})
			return
		}

		log.Fatalln(err)
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Failed to Register.",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Register Success",
		Results: newUser,
	})
}

func ForgotPassword(c *gin.Context) {
	form := ResetForm{}
	c.ShouldBind(&form)
	if form.Email != "" {
		user, _ := models.GetOneUserByEmail(form.Email)
		if user.Id != 0 {
			resetForm := models.ResetForm{
				Otp:   lib.RandomNumber(6),
				Email: user.Email,
			}
			models.CreateResetPassword(resetForm)

			fmt.Println("OTP: ", resetForm.Otp)

			c.JSON(http.StatusOK, &services.ResponseOnly{
				Success: true,
				Message: "OTP has been sent to your email",
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Failed to reset",
			})
			return
		}
	}
	if form.Otp != "" {
		user, _ := models.ResetByOtp(form.Otp)
		if user.Id != 0 {
			if form.Password == form.ConfirmPassword {
				userFound, _ := models.GetOneUserByEmail(user.Email)
				data := models.User{
					Id: userFound.Id,
				}

				hashed, _ := argonize.Hash([]byte(form.Password))
				data.Password = hashed.String()

				updated, _ := models.UpdateUser(data)
				message := fmt.Sprintf("Reset password for %v success", updated.Email)
				c.JSON(http.StatusOK, &services.ResponseOnly{
					Success: true,
					Message: message,
				})
				models.DeleteResetPassword(user.Id)
				return
			} else {
				c.JSON(http.StatusBadRequest, &services.ResponseOnly{
					Success: false,
					Message: "Confirm password doesn't match",
				})
			}
		}
	}
	c.JSON(http.StatusBadRequest, &services.ResponseOnly{
		Success: false,
		Message: "Internal Server Error",
	})
}
