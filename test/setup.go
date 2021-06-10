package test

import (
	"net/http"
)

var MockServerURL = "http://localhost:1080/mockserver/expectation"


func Setup() {

	client := &http.Client{}

	//Clear MockServer Expectations
	req, err := http.NewRequest(http.MethodPut, MockServerURL, nil)
	if err != nil {
		panic(err)
	}

	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}

func Shutdown() {

	client := &http.Client{}

	//Clear MockServer Expectations
	req, err := http.NewRequest(http.MethodPut, MockServerURL, nil)
	if err != nil {
		panic(err)
	}

	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}
