package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type headers struct{
	UserId string
	Browser string
}

type personData struct{
	FirstName string
	LastName string
}
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


func (h *TestHandler) HeaderBinder1(c *gin.Context){

	userID := c.GetHeader("UserId")

	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"userId": userID,
	})

}

func (h *TestHandler) HeaderBinder2(c *gin.Context){
	header := headers{}
	c.BindHeader(&header)
	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder2",
		"header": header,
	})

}
func (h *TestHandler) QueryBinder1(c *gin.Context){
	id := c.Query("id")
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"result": "QueryBinder 1",
		"id": id,
		"name": name,
	})

}
func (h *TestHandler) QueryBinder2(c *gin.Context){
	ids := c.QueryArray("id")
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"result": "QueryBinder 1",
		"id": ids,
		"name": name,
	})

}

func (h *TestHandler) UriBinder(c *gin.Context){

	id := c.Param("id")
	name := c.Param("name")

	c.JSON(http.StatusOK, gin.H{
		"result": "UriBinder",
		"id": id,
		"name": name,
	})

}
func (h *TestHandler) BinderBody(c *gin.Context){

	p := personData{}
	c.ShouldBindJSON(&p)
	c.JSON(http.StatusOK, gin.H{
		"result": "BinderBody",
		"data" : p,
	})

}
func (h *TestHandler) BinderBody1(c *gin.Context){

	p := personData{}
	c.ShouldBindJSON(&p)
	c.JSON(http.StatusOK, gin.H{
		"result": "BinderBody",
		"data" : p,
	})

}


