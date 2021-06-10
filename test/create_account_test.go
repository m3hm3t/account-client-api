package test

import (
	"bytes"
	"github.com/m3hm3t/account-client-api/internal/config"
	"github.com/m3hm3t/account-client-api/internal/pkg/rest/adapter/account/creator"
	"github.com/m3hm3t/account-client-api/internal/pkg/rest/adapter/account/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type CreateAccountTestSuite struct {
	suite.Suite
}

func TestCreateAccountTestSuite(t *testing.T) {
	suite.Run(t, new(CreateAccountTestSuite))
}

func (s *CreateAccountTestSuite) TestShouldCreateAccountWhenMockAPIReturnSuccess() {

	// Given
	if mockErr := s.setMockAccountAPIReturnSuccess(); mockErr != nil {
		s.Error(mockErr)
	}

	config.AccountAPIURL = MockServerURL

	accountRequest := dto.RequestDto{
		Data: dto.DataRequestDto{
			ID:             "396a3a28-2247-4474-bfa0-ce214a66c07a",
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
	expectedAccountResponse := dto.ResponseDto{
		Data: dto.DataResponseDto{
			ID:             "496a3a28-2247-4474-bfa0-ce214a66c07b",
			Type:           "accounts",
			OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Version:        0,
			ModifiedOn: "2021-06-10T12:38:04.627Z",
			CreatedOn: "2021-06-10T12:38:04.627Z",
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

	creatorAdapter := creator.ProvideAccountCreator2()

	// When
	err := creatorAdapter.CreateAccount(accountRequest, &actualAccountResponse)

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), expectedAccountResponse, actualAccountResponse)
}

func (s *CreateAccountTestSuite) TestShouldReturnErrorWhenMockAPIReturnErrorMessage() {

	// Given
	if mockErr := s.setMockAccountAPIReturnError(); mockErr != nil {
		s.Error(mockErr)
	}

	config.AccountAPIURL = MockServerURL

	accountRequest := dto.RequestDto{
		Data: dto.DataRequestDto{
			ID:             "396a3a28-2247-4474-bfa0-ce214a66c07a",
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

	creatorAdapter := creator.ProvideAccountCreator2()

	// When
	err := creatorAdapter.CreateAccount(accountRequest, &actualAccountResponse)

	// Then
	assert.NotNil(s.T(), err)
	assert.Empty(s.T(), actualAccountResponse)
}

func (s *CreateAccountTestSuite) setMockAccountAPIReturnSuccess() error {
	client := &http.Client{}

	body := `{
				"httpRequest": {
					"method": "POST"
				},
				"httpResponse": {
					"statusCode": 201,
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

func (s *CreateAccountTestSuite) setMockAccountAPIReturnError() error {
	client := &http.Client{}

	body := `{
				"httpRequest": {
					"method": "POST"
				},
				"httpResponse": {
					"statusCode": 400,
					"body": {
						"error_message": "validation failure list:\nvalidation failure list:\nid in body must be of type uuid: \"496a3a28-2247-4474-bfa0-ce214a66c07aa\""
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