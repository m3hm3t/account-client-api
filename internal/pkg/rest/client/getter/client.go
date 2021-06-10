package getter

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type RestGetterClient struct {
	Client *http.Client
}

func NewRestGetterClient(timeOutInMilliSec int64) RestGetter {
	return &RestGetterClient{
		Client: &http.Client{
			Timeout: time.Duration(timeOutInMilliSec) * time.Millisecond,
		},
	}
}

func (c *RestGetterClient) MakeGetRequest(url string) ([]byte, int, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("request cannot be created: %w", err)
	}

	request.Header.Add("Accept", `application/vnd.api+json`)

	resp, err := c.Client.Do(request)
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("request is failed: %w", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("response body cannot be read: %w", err)
	}

	return body, resp.StatusCode, nil
}
