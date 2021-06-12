package deleter

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type RestDeleterAsyncClient struct {
	client                *http.Client
	ResponseStatusChannel chan int
	ErrorChannel          chan error
}

func NewRestDeleterAsyncClient(timeOutInMilliSec int64, bufferSize int) *RestDeleterAsyncClient {
	responseStatusChannel := make(chan int, bufferSize)
	errorChannel := make(chan error, bufferSize)
	return &RestDeleterAsyncClient{
		client: &http.Client{
			Timeout: time.Duration(timeOutInMilliSec) * time.Millisecond,
		},
		ResponseStatusChannel: responseStatusChannel,
		ErrorChannel:          errorChannel,
	}
}

func (c *RestDeleterAsyncClient) DeleteAll(urls []string) {
	for _, url := range urls {
		go c.makeDeleteRequest(url)
	}
}

func (c *RestDeleterAsyncClient) makeDeleteRequest(url string) {
	ctx := context.Background()

	request, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		c.ErrorChannel <- fmt.Errorf("request cannot be deleted: %w", err)
		return
	}

	request.Header.Add("Accept", `application/vnd.api+json`)

	resp, err := c.client.Do(request)
	if err != nil {
		c.ErrorChannel <- fmt.Errorf("request is failed: %w", err)
		return
	}

	defer resp.Body.Close()

	c.ResponseStatusChannel <- resp.StatusCode
}
