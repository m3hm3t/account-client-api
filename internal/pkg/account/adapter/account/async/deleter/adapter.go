package deleter

import (
	"github.com/m3hm3t/account-client-api/internal/config"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/client/async/deleter"
)

type Adapter struct {
	url string
}

func NewAdapter() AccountDeleterAsync {
	return &Adapter{
		url: config.AccountAPIURL,
	}
}

func (a *Adapter) DeleteAll(accountIDs []string, versions []string) ([]int, error) {
	deleteURLs := make([]string, len(accountIDs))
	for index, accountID := range accountIDs {
		deleteURL := a.url + "/" + accountID + "/?version=" + versions[index]
		deleteURLs = append(deleteURLs, deleteURL)
	}

	deleteAsyncClient := deleter.ProvideRestDeleterAsyncClient(len(accountIDs))

	deleteAsyncClient.DeleteAll(deleteURLs)

	responseStatuses := make([]int, len(accountIDs))

	for i := 0; i < len(accountIDs); i++ {
		select {
		case responseStatus := <- deleteAsyncClient.ResponseStatusChannel:
			responseStatuses = append(responseStatuses, responseStatus)
		case err := <- deleteAsyncClient.ErrorChannel:
			return responseStatuses, err
		}
	}

	return responseStatuses, nil
}
