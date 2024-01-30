package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

var users = []Users{
	{
		Id:       1,
		Email:    "tony.stark@gmail.com",
		Password: "tony1234",
	},
	{
		Id:       2,
		Email:    "thor.thunder@gmail.com",
		Password: "kingofthunder",
	},
	{
		Id:       3,
		Email:    "peter.spider@gmail.com",
		Password: "parker123",
	},
	{
		Id:       4,
		Email:    "stephen.strange@gmail.com",
		Password: "strange1212",
	},
}

func ListAllUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))

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

	var getUser *Users
	for _, user := range users {
		if user.Id == id {
			getUser = &user
			break
		}
	}

	if getUser == nil {
		c.JSON(http.StatusNotFound, &errResponse{
			Success: false,
			Message: "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "Detail Users",
		Results: getUser,
	})
}

func CreateUser(c *gin.Context) {
	defer panicHandle(c)

	newUser := Users{}

	c.Bind(&newUser)

	if newUser.Password == "" {
		panic("Password cannot be empty")
	}

	newUser.Id = len(users) + 1

	users = append(users, newUser)

	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "User Created Successfully",
		Results: newUser,
	})
}

func UpdateUser(c *gin.Context) {
	defer panicHandle(c)

	id, _ := strconv.Atoi(c.Param("id"))
	updateUser := Users{}

	var getUser *Users
	for i, user := range users {
		if user.Id == id {
			getUser = &users[i]
			break
		}
	}

	if getUser == nil {
		panic("User not found")
	}

	c.Bind(&updateUser)

	if updateUser.Email != "" {
		getUser.Email = updateUser.Email
	}

	if updateUser.Password != "" {
		getUser.Password = updateUser.Password
	}

	if updateUser.Email == "" && updateUser.Password == "" {
		panic("No data has been changed")
	}

	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "User Updated Successfully",
		Results: getUser,
	})
}

func DeleteUser(c *gin.Context) {
	defer panicHandle(c)

	id, _ := strconv.Atoi(c.Param("id"))

	index := -1
	for i, user := range users {
		if user.Id == id {
			index = i
			break
		}
	}

	if index != -1 {
		deletedUser := users[index]
		users = append(users[:index], users[index+1:]...)
		c.JSON(http.StatusOK, &response{
			Success: true,
			Message: "User deleted successfully",
			Results: deletedUser,
		})
	} else {
		panic("User not found")
	}
}

func panicHandle(c *gin.Context) {
	if r := recover(); r != nil {
		var errMsg string
		switch v := r.(type) {
		case string:
			errMsg = v
		default:
			errMsg = "Internal Server Error"
		}
		c.JSON(http.StatusInternalServerError, &errResponse{
			Success: false,
			Message: errMsg,
		})
	}
}
