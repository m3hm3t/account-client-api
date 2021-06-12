package fetch

import (
	"encoding/json"
	"fmt"
	"github.com/m3hm3t/account-client-api/internal/config"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/adapter/account/dto"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/client/get"
	"net/http"
)

type Adapter struct {
	restGetter get.RestGetter
	url        string
}

func NewAdapter(restGetter get.RestGetter) AccountFetcher {
	return &Adapter{
		restGetter: restGetter,
		url:        config.AccountAPIURL,
	}
}

func (a *Adapter) FetchAccount(accountID string, accountResponse *dto.ResponseDto) error {
	fetcherURL := a.url + "/" + accountID

	responseBody, responseStatus, err := a.restGetter.MakeGetRequest(fetcherURL)
	if err != nil {
		return fmt.Errorf("fetch account error: %w", err)
	}

	if responseStatus != http.StatusOK {
		return generateError(responseStatus, responseBody)
	}

	if err := a.BuildResponse(responseBody, accountResponse); err != nil {
		return err
	}

	return nil
}

func (a *Adapter) BuildResponse(responseBody []byte, accountResponse *dto.ResponseDto) error {
	err := json.Unmarshal(responseBody, accountResponse)
	if err != nil {
		return fmt.Errorf("account response unmarshal error: %w", err)
	}

	return nil
}
