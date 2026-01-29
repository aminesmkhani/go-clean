package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestHandler struct{

}

func NewTestHandler() *TestHandler{
	return &TestHandler{}
}

func(h *TestHandler) Test(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"result" : "Test",
	})
}

func(h *TestHandler) Users(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"result" : "Users",
	})
}
func(h *TestHandler) UserById(c *gin.Context){
	id := c.Param("id")
	c.JSON(http.StatusOK,gin.H{
		"result" : "User By Id",
		"id" : id,
	})

}
func(h *TestHandler) UserByUserName(c *gin.Context){
	userName := c.Param("username")
	c.JSON(http.StatusOK,gin.H{
		"result" : "User By User Name",
		"username": userName,
	})
}

func(h *TestHandler) Accounts(c *gin.Context){
	id := c.Param("id")
	c.JSON(http.StatusOK,gin.H{
		"result" : "Accounts",
		"id": id,
	})
}


func(h *TestHandler) AddUser(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"result" : "User Created",
	})
}



