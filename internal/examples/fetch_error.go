// +build ignore

package main

import (
	"github.com/google/uuid"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/adapter/account/dto"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/adapter/account/fetcher"
	"log"
)

func main() {

	fetcherAdapter := fetcher.ProvideAccountFetcher()

	accountResponse := dto.ResponseDto{}

	// try fetch with invalid account id
	accountID := "invalid_id"
	err := fetcherAdapter.FetchAccount(accountID, &accountResponse)
	if err != nil {
		log.Println("ERROR: ", err)
	}

	// try fetch with no exist account id
	accountID = uuid.NewString()
	err = fetcherAdapter.FetchAccount(accountID, &accountResponse)
	if err != nil {
		log.Println("ERROR: ", err)
	}
}
