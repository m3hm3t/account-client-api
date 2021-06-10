// +build ignore

package main

import (
	"github.com/google/uuid"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/adapter/account/creator"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/adapter/account/dto"
	"log"
)

func main() {

	creatorAdapter := creator.ProvideAccountCreator()

	accountResponse := dto.ResponseDto{}

	accountRequest := dto.RequestDto{
		Data: dto.DataRequestDto{
			ID:             uuid.NewString(),
			Type:           "accounts",
			OrganisationID: "invalid_id",
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

	err := creatorAdapter.CreateAccount(accountRequest, &accountResponse)
	if err != nil {
		log.Println("ERROR: ", err)
	}
}
