package admin_controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/KEINOS/go-argonize"
	"github.com/gin-gonic/gin"
	"github.com/mfdrfauzi/fwg17-go-backend/src/models"
	"github.com/mfdrfauzi/fwg17-go-backend/src/services"
)

type Users struct {
	Id       int    `json:"id"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ListAllUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	search := c.Query("search")
	sortBy := c.Query("sortBy")
	orderBy := c.Query("orderBy")

	if page <= 0 {
		page = 1
	}

	users, totalPage, err := models.GetAllUser(search, sortBy, orderBy, page)

	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	var result interface{}
	if len(users) == 0 {
		result = "User Not Found"
	} else {
		result = users
	}

	nextPage := page + 1
	prevPage := page - 1
	if nextPage > totalPage {
		nextPage = 0
	}

	if prevPage < 1 {
		prevPage = 0
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "List All Users",
		PageInfo: services.PageInfo{
			Page:      page,
			TotalPage: totalPage,
			NextPage:  nextPage,
			PrevPage:  prevPage,
		},
		Results: result,
	})
}

func DetailUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := models.GetOneUser(id)

	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
				Success: false,
				Message: "User Not Found",
			})
			return
		}

		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Detail Users",
		Results: users,
	})
}

func CreateUser(c *gin.Context) {
	data := models.User{}

	c.Bind(&data)

	plainPass := []byte(data.Password)
	hashedPass, err := argonize.Hash(plainPass)
	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Failed to generate Hash",
		})
	}

	data.Password = hashedPass.String()

	newUser, err := models.CreateUser(data)

	if err != nil {
		if strings.HasPrefix(err.Error(), "pq: duplicate key value") {
			c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
				Success: false,
				Message: "Email already exists",
			})
			return
		}

		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "User Created Successfully",
		Results: newUser,
	})
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.User{}

	c.Bind(&data)
	data.Id = id

	plainPass := []byte(data.Password)
	hashedPass, err := argonize.Hash(plainPass)
	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Failed to generate Hash",
		})
	}

	data.Password = hashedPass.String()

	user, err := models.UpdateUser(data)

	if err != nil {
		if strings.HasPrefix(err.Error(), "pq: duplicate key value") {
			c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
				Success: false,
				Message: "Email already exists",
			})
			return
		}

		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "User Updated Successfully",
		Results: user,
	})
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := models.DeleteUser(id)

	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
				Success: false,
				Message: "User Not Found",
			})
			return
		}

		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Deleted Users",
		Results: users,
	})
}
