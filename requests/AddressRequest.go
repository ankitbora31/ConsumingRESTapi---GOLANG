package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"micro/handler"
	"micro/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AddressWeb struct {
	State string `json:"state"`
	Pin   string `json:"pin"`
	City  string `json:"city"`
}

func GetAddress(c *gin.Context) {
	userid, _ := c.Params.Get("userid")
	urlUser := fmt.Sprintf("http://localhost:8080/users/%s", userid)

	id, _ := handler.UserResolve(urlUser)

	if id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No User"})
	} else {
		urlAddress := fmt.Sprintf("http://localhost:8080/users/%s/addresses", userid)
		info, err := http.Get(urlAddress)
		handler.ErrorHandle(err)

		responseData, _ := ioutil.ReadAll(info.Body)

		var responseObject []responses.AddressResponse
		err = json.Unmarshal([]byte(responseData), &responseObject)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "NO addresss"})
		} else {
			c.JSON(http.StatusOK, gin.H{"users": responseObject})
		}
	}
}

func CreateAddress(c *gin.Context) {
	userid, _ := c.Params.Get("userid")
	urlUser := fmt.Sprintf("http://localhost:8080/users/%s", userid)
	urlAddress := fmt.Sprintf("http://localhost:8080/users/%s/addresses", userid)

	id, _ := handler.UserResolve(urlUser)

	if id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No User"})
	} else {
		var input AddressWeb
		err := c.ShouldBindJSON(&input)
		handler.ErrorHandle(err)

		address := responses.AddressResponse{
			State:  input.State,
			Pin:    input.Pin,
			City:   input.City,
			UserID: userid,
		}
		postBody, _ := json.Marshal(address)
		info, err := http.Post(urlAddress, "application/json", bytes.NewBuffer(postBody))
		handler.ErrorHandle(err)

		ioutil.ReadAll(info.Body)
		c.JSON(http.StatusOK, gin.H{"message": "address created"})
	}
}

func UpdateAddress(c *gin.Context) {
	client := &http.Client{}
	userid, _ := c.Params.Get("userid")
	addressid, _ := c.Params.Get("addressid")

	urlUser := fmt.Sprintf("http://localhost:8080/users/%s", userid)
	id, _ := handler.UserResolve(urlUser)
	if id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No User"})
	} else {
		url := fmt.Sprintf("http://localhost:8080/users/%s/addresses/%s", userid, addressid)
		aid, _ := handler.AddressResolve(url)
		if aid == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		} else {
			var input AddressWeb

			uid, _ := strconv.ParseUint(userid, 10, 64)

			address := responses.AddressResponse{
				ID:     uid,
				State:  input.State,
				Pin:    input.Pin,
				City:   input.City,
				UserID: addressid,
			}
			c.BindJSON(&address)
			data, err := json.Marshal(address)
			handler.ErrorHandle(err)
			req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
			if err != nil {
				handler.ErrorHandle(err)
			}
			req.Header.Set("Content-Type", "application/json; charset=utf-8")
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}

			c.JSON(http.StatusOK, gin.H{"message": resp.Status})
		}
	}
}

func DeleteAddress(c *gin.Context) {
	userid, _ := c.Params.Get("userid")
	addressid, _ := c.Params.Get("addressid")

	urlUser := fmt.Sprintf("http://localhost:8080/users/%s", userid)
	id, _ := handler.UserResolve(urlUser)
	if id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No User"})
	} else {
		url := fmt.Sprintf("http://localhost:8080/users/%s/addresses/%s", userid, addressid)

		aid, _ := handler.AddressResolve(url)
		if aid == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "No Address"})
		} else {
			client := &http.Client{}
			req, err := http.NewRequest("DELETE", url, nil)
			handler.ErrorHandle(err)
			resp, err := client.Do(req)
			handler.ErrorHandle(err)
			c.JSON(http.StatusOK, gin.H{"message": resp.Status})
		}
	}
}

func GetAddressById(c *gin.Context) {
	userid, _ := c.Params.Get("userid")
	addressid, _ := c.Params.Get("addressid")

	urlUser := fmt.Sprintf("http://localhost:8080/users/%s", userid)

	id, _ := handler.UserResolve(urlUser)
	if id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No User"})
	} else {
		urlAddress := fmt.Sprintf("http://localhost:8080/users/%s/addresses/%s", userid, addressid)
		aid, response := handler.AddressResolve(urlAddress)
		if aid == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		} else {
			c.JSON(http.StatusOK, gin.H{"address": response})
		}
	}
}
