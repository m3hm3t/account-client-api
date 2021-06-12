package deleter

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type RestDeleterClient struct {
	client *http.Client
}

func NewRestDeleter(timeOutInMilliSec int64) RestDeleter {
	return &RestDeleterClient{
		client: &http.Client{
			Timeout: time.Duration(timeOutInMilliSec) * time.Millisecond,
		},
	}
}

func (c *RestDeleterClient) MakeDeleteRequest(url string) (int, error) {
	ctx := context.Background()
	request, err := http.NewRequestWithContext(ctx,"DELETE", url, nil)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("request cannot be deleted: %w", err)
	}

	request.Header.Add("Accept", `application/vnd.api+json`)

	resp, err := c.client.Do(request)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("request is failed: %w", err)
	}

	defer resp.Body.Close()

	return resp.StatusCode, nil
}
