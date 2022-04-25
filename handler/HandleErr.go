package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"micro/responses"
	"net/http"
)

func UserResolve(url string) (int, responses.UserResponse) {
	info, err := http.Get(url)
	ErrorHandle(err)

	responseData, _ := ioutil.ReadAll(info.Body)

	var responseObject responses.UserResponse
	json.Unmarshal(responseData, &responseObject)

	return int(responseObject.ID), responseObject
}

func AddressResolve(url string) (int, responses.AddressResponse) {
	info, err := http.Get(url)
	ErrorHandle(err)

	responseData, _ := ioutil.ReadAll(info.Body)

	var responseObject responses.AddressResponse
	json.Unmarshal(responseData, &responseObject)

	return int(responseObject.ID), responseObject
}
func ErrorHandle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
