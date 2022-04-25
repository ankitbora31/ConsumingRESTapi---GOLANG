package requests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"micro/handler"
	"micro/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserWeb struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

func GetUsers(c *gin.Context) {
	info, err := http.Get("http://localhost:8080/users")
	handler.ErrorHandle(err)

	responseData, err := ioutil.ReadAll(info.Body)
	handler.ErrorHandle(err)

	var responseObject []responses.UserResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"users": responseObject})

}

func CreateUser(c *gin.Context) {
	var input UserWeb

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	user := responses.UserResponse{
		Name:   input.Name,
		Gender: input.Gender,
		Age:    input.Age,
	}

	postBody, _ := json.Marshal(user)
	info, err := http.Post("http://localhost:8080/users", "application/json", bytes.NewBuffer(postBody))
	handler.ErrorHandle(err)

	ioutil.ReadAll(info.Body)

	c.JSON(http.StatusOK, gin.H{"message": "created"})
}

func UpdateUser(c *gin.Context) {
	client := &http.Client{}
	url := "http://localhost:8080/users/"
	uid, _ := c.Params.Get("userid")
	url = url + uid

	id, _ := handler.UserResolve(url)

	if id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No User"})
	} else {

		var input UserWeb

		userid, _ := strconv.ParseUint(uid, 10, 64)

		user := responses.UserResponse{
			ID:     userid,
			Name:   input.Name,
			Gender: input.Gender,
			Age:    input.Age,
		}
		c.BindJSON(&user)
		data, err := json.Marshal(user)
		if err != nil {
			handler.ErrorHandle(err)
		}
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
		if err != nil {
			handler.ErrorHandle(err)
		}
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		}

		c.JSON(http.StatusOK, gin.H{"message": resp.Status})
	}

}

func DeleteUser(c *gin.Context) {
	url := "http://localhost:8080/users/"
	uid, _ := c.Params.Get("userid")
	url = url + uid

	id, _ := handler.UserResolve(url)
	if id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No User"})
	} else {
		client := &http.Client{}
		req, err := http.NewRequest("DELETE", url, nil)
		handler.ErrorHandle(err)
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "no user"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": resp.Status})
		}
	}

}

func GetUserById(c *gin.Context) {
	url := "http://localhost:8080/users/"
	uid, _ := c.Params.Get("userid")
	url = url + uid

	id, response := handler.UserResolve(url)

	if id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No User"})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": response})
	}
}
