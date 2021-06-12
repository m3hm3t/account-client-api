package post

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type RestPosterClient struct {
	Client *http.Client
}

func NewRestPosterClient(timeOutInMilliSec int64) RestPoster {
	return &RestPosterClient{
		Client: &http.Client{
			Timeout: time.Duration(timeOutInMilliSec) * time.Millisecond,
		},
	}
}

func (c *RestPosterClient) MakePostRequest(url string, input interface{}) ([]byte, int, error) {
	ctx := context.Background()

	byteArray, err := json.Marshal(input)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, fmt.Errorf("failed to marshal input dto: %w", err)
	}

	request, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(byteArray))
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
