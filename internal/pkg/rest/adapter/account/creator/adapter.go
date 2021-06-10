package creator

import (
	"encoding/json"
	"github.com/m3hm3t/account-client-api/internal/config"
	"github.com/m3hm3t/account-client-api/internal/pkg/rest/adapter/account/dto"
	"github.com/m3hm3t/account-client-api/internal/pkg/rest/client/poster"
	"net/http"
)

type Adapter struct {
	restPoster poster.RestPoster
	url        string
}

func NewAdapter(restPoster poster.RestPoster) AccountCreator {
	return &Adapter{
		restPoster: restPoster,
		url:        config.AccountAPIURL,
	}
}

func (a *Adapter) CreateAccount(accountRequest dto.RequestDto, accountResponse *dto.ResponseDto) error {

	responseBody, responseStatus, err := a.restPoster.MakePostRequest(config.AccountAPIURL, accountRequest)
	if err != nil {
		return err
	}

	if responseStatus != http.StatusCreated {
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
		return err
	}

	return nil
}
