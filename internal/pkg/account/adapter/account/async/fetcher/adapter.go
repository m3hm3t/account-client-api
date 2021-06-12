package fetcher

import (
	"encoding/json"
	"github.com/m3hm3t/account-client-api/internal/config"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/adapter/account/dto"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/client/async/getter"
)

type Adapter struct {
	url string
}

func NewAdapter() AccountFetcherAsync {
	return &Adapter{
		url: config.AccountAPIURL,
	}
}

func (a *Adapter) FetchAll(accountIDs []string, accountResponses []dto.ResponseDto) error {
	fetcherURLs := make([]string, len(accountIDs))
	for index, accountID := range accountIDs {
		fetcherURL := a.url + "/" + accountID
		fetcherURLs[index] = fetcherURL
	}

	getterAsyncClient := getter.ProvideRestGetterAsyncClient(len(accountIDs))

	getterAsyncClient.FetchAll(fetcherURLs)

	for i := 0; i < len(accountIDs); i++ {
		select {
		case responseBody := <-getterAsyncClient.ResponseChannel:
			var accountResponse dto.ResponseDto
			if err := a.BuildResponse(responseBody, &accountResponse); err != nil {
				return err
			}
			accountResponses[i] = accountResponse
		case err := <-getterAsyncClient.ErrorChannel:
			return err
		}
	}

	return nil
}

func (a *Adapter) BuildResponse(responseBody []byte, accountResponse *dto.ResponseDto) error {
	err := json.Unmarshal(responseBody, accountResponse)
	if err != nil {
		return err
	}

	return nil
}
