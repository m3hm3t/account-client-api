package deleter

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type RestDeleterClient struct {
	Client *http.Client
}

func NewRestDeleter(timeOutInMilliSec int64) RestDeleter {
	return &RestDeleterClient{
		Client: &http.Client{
			Timeout: time.Duration(timeOutInMilliSec) * time.Millisecond,
		},
	}
}

func (c *RestDeleterClient) MakeDeleteRequest(url string) (int, error) {

	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("request cannot be created: %w", err)
	}

	request.Header.Add("Accept", `application/vnd.api+json`)

	resp, err := c.Client.Do(request)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("request is failed: %w", err)
	}

	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("response body cannot be read: %w", err)
	}

	return resp.StatusCode, nil
}
