package create

import (
	"encoding/json"
	"fmt"
	"github.com/m3hm3t/account-client-api/internal/config"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/adapter/account/dto"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/client/post"
	"net/http"
)

type Adapter struct {
	restPoster post.RestPoster
	url        string
}

func NewAdapter(restPoster post.RestPoster) AccountCreator {
	return &Adapter{
		restPoster: restPoster,
		url:        config.AccountAPIURL,
	}
}

func (a *Adapter) CreateAccount(accountRequest dto.RequestDto, accountResponse *dto.ResponseDto) error {
	responseBody, responseStatus, err := a.restPoster.MakePostRequest(config.AccountAPIURL, accountRequest)
	if err != nil {
		return fmt.Errorf("create account error: %w", err)
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
		return fmt.Errorf("account response unmarshal error: %w", err)
	}

	return nil
}
