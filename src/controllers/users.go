package controller

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mfdrfauzi/fwg17-go-backend/src/models"
)

type Users struct {
	Id       int    `json:"id"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type pageInfo struct {
	Page      int `json:"page"`
	TotalPage int `json:"totalPage"`
	NextPage  int `json:"nextPage,omitempty"`
	PrevPage  int `json:"prevPage,omitempty"`
}

type responseList struct {
	Success  bool        `json:"success"`
	Message  string      `json:"message"`
	PageInfo pageInfo    `json:"pageInfo"`
	Results  interface{} `json:"results"`
}
type response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Results interface{} `json:"results"`
}

type errResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
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
		c.JSON(http.StatusInternalServerError, &errResponse{
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

	c.JSON(http.StatusOK, &responseList{
		Success: true,
		Message: "List All Users",
		PageInfo: pageInfo{
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
	users, err := models.FindOneUser(id)

	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusInternalServerError, &errResponse{
				Success: false,
				Message: "User Not Found",
			})
			return
		}

		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, &errResponse{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "Detail Users",
		Results: users,
	})
}

func CreateUser(c *gin.Context) {
	data := models.User{}

	c.Bind(&data)

	newUser, err := models.CreateUser(data)

	if err != nil {
		if strings.HasPrefix(err.Error(), "pq: duplicate key value") {
			c.JSON(http.StatusInternalServerError, &errResponse{
				Success: false,
				Message: "Email already exists",
			})
			return
		}

		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, &errResponse{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &response{
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

	user, err := models.UpdateUser(data)

	if err != nil {
		if strings.HasPrefix(err.Error(), "pq: duplicate key value") {
			c.JSON(http.StatusInternalServerError, &errResponse{
				Success: false,
				Message: "Email already exists",
			})
			return
		}

		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, &errResponse{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &response{
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
			c.JSON(http.StatusInternalServerError, &errResponse{
				Success: false,
				Message: "User Not Found",
			})
			return
		}

		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, &errResponse{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "Deleted Users",
		Results: users,
	})
}
