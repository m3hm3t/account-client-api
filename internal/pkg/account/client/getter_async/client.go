package getter_async

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type RestGetterAsyncClient struct {
	HTTPClient      *http.Client
	ResponseChannel chan []byte
	ErrorChannel    chan error
}

func NewRestGetterAsyncClient(timeOutInMilliSec int64, bufferSize int) *RestGetterAsyncClient {
	responseChannel := make(chan []byte, bufferSize)
	errorChannel := make(chan error, bufferSize)
	return &RestGetterAsyncClient{
		HTTPClient:      &http.Client{
			Timeout: time.Duration(timeOutInMilliSec) * time.Millisecond,
		},
		ResponseChannel: responseChannel,
		ErrorChannel:    errorChannel,
	}
}

func (c *RestGetterAsyncClient) FetchAll(urls []string) {
	for _, url := range urls {
		go c.makeGetRequest(url)
	}
}

func (c *RestGetterAsyncClient) makeGetRequest(url string) {
	ctx := context.Background()

	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		c.ErrorChannel <- fmt.Errorf("request cannot be created: %w", err)
		return
	}

	request.Header.Add("Accept", `application/vnd.api+json`)

	resp, err := c.HTTPClient.Do(request)
	if err != nil {
		c.ErrorChannel <- fmt.Errorf("request is failed: %w", err)
		return
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorChannel <- fmt.Errorf("response responseBody cannot be read: %w", err)
		return
	}

	c.ResponseChannel <- responseBody
}
