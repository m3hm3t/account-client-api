package test

import (
	"github.com/google/uuid"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/adapter/account/creator"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/adapter/account/deleter"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/adapter/account/dto"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/adapter/account/fetcher"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"strconv"
	"testing"
)

type CreateFetchDeleteAccountTestSuite struct {
	suite.Suite
}

func TestCreateFetchDeleteAccountTestSuite(t *testing.T) {
	suite.Run(t, new(CreateFetchDeleteAccountTestSuite))
}

func (s *CreateFetchDeleteAccountTestSuite) TestShouldCreateFetchDeleteAccountWhenRequestDataValid() {
	// Given
	accountRequest := dto.RequestDto{
		Data: dto.DataRequestDto{
			ID:             uuid.NewString(),
			Type:           "accounts",
			OrganisationID: "496a3a28-2247-4474-bfa0-ce214a66c07b",
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

	actualAccountResponse := dto.ResponseDto{}

	creatorAdapter := creator.ProvideAccountCreator()
	fetcherAdapter := fetcher.ProvideAccountFetcher()
	deleteAdapter := deleter.ProvideAccountDeleter()

	// When
	err := creatorAdapter.CreateAccount(accountRequest, &actualAccountResponse)

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), accountRequest.Data.ID, actualAccountResponse.Data.ID)

	// WHEN
	err = fetcherAdapter.FetchAccount(accountRequest.Data.ID, &actualAccountResponse)

	// THEN
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), accountRequest.Data.ID, actualAccountResponse.Data.ID)

	// WHEN
	responseStatus, deleteErr := deleteAdapter.DeleteAccount(accountRequest.Data.ID,
		strconv.Itoa(actualAccountResponse.Data.Version))

	// THEN
	assert.Nil(s.T(), deleteErr)
	assert.Equal(s.T(), http.StatusNoContent, responseStatus)
}
