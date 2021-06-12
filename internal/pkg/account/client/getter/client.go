package getter

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type RestGetterClient struct {
	client *http.Client
}

func NewRestGetterClient(timeOutInMilliSec int64) RestGetter {
	return &RestGetterClient{
		client: &http.Client{
			Timeout: time.Duration(timeOutInMilliSec) * time.Millisecond,
		},
	}
}

func (c *RestGetterClient) MakeGetRequest(url string) ([]byte, int, error) {
	ctx := context.Background()

	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("request cannot be fetch: %w", err)
	}

	request.Header.Add("Accept", `application/vnd.api+json`)

	resp, err := c.client.Do(request)
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
