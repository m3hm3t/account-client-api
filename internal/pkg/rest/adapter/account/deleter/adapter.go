package deleter

import (
	"github.com/m3hm3t/account-client-api/internal/config"
	deleterclient "github.com/m3hm3t/account-client-api/internal/pkg/rest/client/deleter"
	"net/http"
)

type Adapter struct {
	restDeleter deleterclient.RestDeleter
	url         string
}

func NewAdapter(restDeleter deleterclient.RestDeleter) AccountDeleter {
	return &Adapter{
		restDeleter: restDeleter,
		url:         config.AccountAPIURL,
	}
}

func (a *Adapter) DeleteAccount(accountID string, version string) (int, error) {
	deleterURL := a.url + "/" + accountID + "/?version=" + version

	responseStatus, err := a.restDeleter.MakeDeleteRequest(deleterURL)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return responseStatus, nil
}
