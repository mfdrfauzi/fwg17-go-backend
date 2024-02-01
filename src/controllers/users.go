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
	Page int `json:"page"`
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
	users, err := models.GetAllUser()

	if err != nil {
		c.JSON(http.StatusInternalServerError, &errResponse{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &responseList{
		Success: true,
		Message: "List All Users",
		PageInfo: pageInfo{
			Page: page,
		},
		Results: users,
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

// func UpdateUser(c *gin.Context) {
// 	defer panicHandle(c)

// 	id, _ := strconv.Atoi(c.Param("id"))
// 	updateUser := Users{}

// 	var getUser *Users
// 	for i, user := range users {
// 		if user.Id == id {
// 			getUser = &users[i]
// 			break
// 		}
// 	}

// 	if getUser == nil {
// 		panic("User not found")
// 	}

// 	c.Bind(&updateUser)

// 	if updateUser.Email != "" {
// 		getUser.Email = updateUser.Email
// 	}

// 	if updateUser.Password != "" {
// 		getUser.Password = updateUser.Password
// 	}

// 	if updateUser.Email == "" && updateUser.Password == "" {
// 		panic("No data has been changed")
// 	}

// 	c.JSON(http.StatusOK, &response{
// 		Success: true,
// 		Message: "User Updated Successfully",
// 		Results: getUser,
// 	})
// }

// func DeleteUser(c *gin.Context) {
// 	defer panicHandle(c)

// 	id, _ := strconv.Atoi(c.Param("id"))

// 	index := -1
// 	for i, user := range users {
// 		if user.Id == id {
// 			index = i
// 			break
// 		}
// 	}

// 	if index != -1 {
// 		deletedUser := users[index]
// 		users = append(users[:index], users[index+1:]...)
// 		c.JSON(http.StatusOK, &response{
// 			Success: true,
// 			Message: "User deleted successfully",
// 			Results: deletedUser,
// 		})
// 	} else {
// 		panic("User not found")
// 	}
// }

// func panicHandle(c *gin.Context) {
// 	if r := recover(); r != nil {
// 		var errMsg string
// 		switch v := r.(type) {
// 		case string:
// 			errMsg = v
// 		default:
// 			errMsg = "Internal Server Error"
// 		}
// 		c.JSON(http.StatusInternalServerError, &errResponse{
// 			Success: false,
// 			Message: errMsg,
// 		})
// 	}
// }
