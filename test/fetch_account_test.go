package test

import (
	"bytes"
	"github.com/m3hm3t/account-client-api/internal/config"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/adapter/account/dto"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/adapter/account/fetcher"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type FetchAccountTestSuite struct {
	suite.Suite
}

func TestFetchAccountTestSuite(t *testing.T) {
	suite.Run(t, new(FetchAccountTestSuite))
}

func (s *FetchAccountTestSuite) TestShouldFetchAccountWhenAccountIsCreated() {
	// GIVEN
	if mockErr := s.setMockAccountAPIReturnSuccess(); mockErr != nil {
		s.Error(mockErr)
	}

	config.AccountAPIURL = MockServerURL
	accountID := "MockAccountID"

	actualAccountResponse := dto.ResponseDto{}
	expectedAccountResponse := dto.ResponseDto{
		Data: dto.DataResponseDto{
			ID:             "496a3a28-2247-4474-bfa0-ce214a66c07b",
			Type:           "accounts",
			OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Version:        0,
			ModifiedOn:     "2021-06-10T12:38:04.627Z",
			CreatedOn:      "2021-06-10T12:38:04.627Z",
			Attributes: dto.AttributesResponseDto{
				Country:                 "GB",
				BaseCurrency:            "GBP",
				BankID:                  "400300",
				BankIDCode:              "GBDSC",
				Name:                    []string{"Samantha Holder"},
				AlternativeName:         []string{"Sam Holder"},
				AccountClassification:   "Personal",
				JointAccount:            false,
				AccountMatchingOptOut:   false,
				SecondaryIdentification: "A1B2C3D4",
			},
		},
	}

	fetcherAdapter := fetcher.ProvideAccountFetcher()

	// WHEN
	err := fetcherAdapter.FetchAccount(accountID, &actualAccountResponse)

	// THEN
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), expectedAccountResponse, actualAccountResponse)
}

func (s *FetchAccountTestSuite) TestShouldReturnErrorWhenMockAPIReturnError() {
	// GIVEN
	if mockErr := s.setMockAccountAPIReturnError(); mockErr != nil {
		s.Error(mockErr)
	}

	config.AccountAPIURL = MockServerURL
	accountID := "MockAccountID"

	actualAccountResponse := dto.ResponseDto{}

	fetcherAdapter := fetcher.ProvideAccountFetcher()

	// WHEN
	err := fetcherAdapter.FetchAccount(accountID, &actualAccountResponse)

	// THEN
	assert.NotNil(s.T(), err)
	assert.Empty(s.T(), actualAccountResponse)
}

func (s *FetchAccountTestSuite) setMockAccountAPIReturnSuccess() error {
	client := &http.Client{}

	body := `{
				"httpRequest": {
					"method": "GET"
				},
				"httpResponse": {
					"statusCode": 200,
					"body": {
						"data": {
							"attributes": {
								"account_classification": "Personal",
								"account_matching_opt_out": false,
								"alternative_names": [
									"Sam Holder"
								],
								"bank_id": "400300",
								"bank_id_code": "GBDSC",
								"base_currency": "GBP",
								"bic": "NWBKGB22",
								"country": "GB",
								"joint_account": false,
								"name": [
									"Samantha Holder"
								],
								"secondary_identification": "A1B2C3D4"
							},
							"created_on": "2021-06-10T12:38:04.627Z",
							"id": "496a3a28-2247-4474-bfa0-ce214a66c07b",
							"modified_on": "2021-06-10T12:38:04.627Z",
							"organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
							"type": "accounts",
							"version": 0
						},
						"links": {
							"self": "/v1/organisation/accounts/496a3a28-2247-4474-bfa0-ce214a66c07b"
						}
					}
				},
				"times": {
					"unlimited": false,
					"remainingTimes" : 1
					}
				}`

	req, err := http.NewRequest(http.MethodPut, MockServerURL, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return err
	}

	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func (s *FetchAccountTestSuite) setMockAccountAPIReturnError() error {
	client := &http.Client{}

	body := `{
				"httpRequest": {
					"method": "GET"
				},
				"httpResponse": {
					"statusCode": 400,
					"body": {
						"error_message": "record 396a3a28-2247-4474-bfa0-ce214a66c07b does not exist"
					}
				},
				"times": {
					"unlimited": false,
					"remainingTimes" : 1
					}
				}`

	req, err := http.NewRequest(http.MethodPut, MockServerURL, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return err
	}

	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}
