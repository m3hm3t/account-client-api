// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/m3hm3t/account-client-api/internal/pkg/rest/adapter/account/creator"
	"github.com/m3hm3t/account-client-api/internal/pkg/rest/adapter/account/deleter"
	"github.com/m3hm3t/account-client-api/internal/pkg/rest/adapter/account/dto"
	"github.com/m3hm3t/account-client-api/internal/pkg/rest/adapter/account/fetcher"
	"log"
	"strconv"
)

func main() {

	creatorAdapter := creator.ProvideAccountCreator2()
	fetcherAdapter := fetcher.ProvideAccountFetcher2()
	deleterAdapter := deleter.ProvideAccountDeleter2()

	accountResponse := dto.ResponseDto{}

	// fetch invalid account id example
	accountID := "invalid_id"
	err := fetcherAdapter.FetchAccount(accountID, &accountResponse)
	if err != nil {
		log.Println("ERROR: ", err)
	} else {
		printJSON("account create response: ", accountResponse)
	}

	accountRequest := dto.RequestDto{
		Data: dto.DataRequestDto{
			ID:             uuid.NewString(),
			Type:           "accounts",
			OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Attributes: dto.AttributesRequestDto{
				Country:                 "GB",
				BaseCurrency:            "GBP",
				BankID:                  "400300",
				BankIDCode:              "GBDSC",
				BIC:                     "NWBKGB22",
				Name:                    []string{"Samantha Holder"},
				AlternativeName:         []string{"Sam Holder"},
				AccountClassification:   "Personal",
				JointAccount:            false,
				AccountMatchingOptOut:   false,
				SecondaryIdentification: "A1B2C3D4",
			},
		},
	}

	err = creatorAdapter.CreateAccount(accountRequest, &accountResponse)
	if err != nil {
		log.Println("ERROR: ", err)
	} else {
		printJSON("account fetch response: ", accountResponse)
	}

	err = fetcherAdapter.FetchAccount(accountRequest.Data.ID, &accountResponse)
	if err != nil {
		log.Println("ERROR: ", err)
	} else {
		printJSON("account fetch response: ", accountResponse)
	}

	var deleteResponseStatus int
	version := strconv.FormatInt(int64(accountResponse.Data.Version), 10)
	deleteResponseStatus, err = deleterAdapter.DeleteAccount(accountRequest.Data.ID, version)
	if err != nil {
		log.Println("ERROR: ", err)
	} else {
		log.Println("account delete response: ", deleteResponseStatus)
	}
}

func printJSON(msg string, data interface{}) {
	b, _ := json.Marshal(data)
	fmt.Printf("%s: %s\n\n", msg, b)
}
