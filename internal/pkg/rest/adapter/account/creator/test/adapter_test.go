package test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/m3hm3t/account-client-api/internal/config"
	"github.com/m3hm3t/account-client-api/internal/pkg/rest/adapter/account/creator"
	"github.com/m3hm3t/account-client-api/internal/pkg/rest/adapter/account/dto"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestShouldFetchAccountWhenAccountIsCreated(t *testing.T) {
	t.Parallel()

	// GIVEN
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	createResponse := `{
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
			"created_on": "2021-06-10T10:07:36.136Z",
			"id": "396a3a28-2247-4474-bfa0-ce214a66c07a",
			"modified_on": "2021-06-10T10:07:36.136Z",
			"organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			"type": "accounts",
			"version": 0
		},
		"links": {
			"self": "/v1/organisation/accounts/396a3a28-2247-4474-bfa0-ce214a66c07a"
		}
	}`

	accountRequest := dto.RequestDto{
		Data: dto.DataRequestDto{
			ID:             "396a3a28-2247-4474-bfa0-ce214a66c07a",
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

	expectedURL := config.AccountAPIURL

	mockPosterRestClient := NewMockRestPoster(ctrl)
	mockPosterRestClient.EXPECT().MakePostRequest(expectedURL, accountRequest).
		Return([]byte(createResponse), http.StatusCreated, nil).Times(1)

	actualAccountResponse := dto.ResponseDto{}
	expectedAccountResponse := dto.ResponseDto{
		Data: dto.DataResponseDto{
			ID:             "396a3a28-2247-4474-bfa0-ce214a66c07a",
			Type:           "accounts",
			OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Version:        0,
			ModifiedOn:     "2021-06-10T10:07:36.136Z",
			CreatedOn:      "2021-06-10T10:07:36.136Z",
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

	creatorAdapter := creator.NewAdapter(mockPosterRestClient)

	// WHEN
	err := creatorAdapter.CreateAccount(accountRequest, &actualAccountResponse)

	// THEN
	assert.Nil(t, err)
	assert.Equal(t, expectedAccountResponse, actualAccountResponse)
}

func TestShouldReturnErrorWhenAccountIsCreated(t *testing.T) {
	t.Parallel()

	// GIVEN
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	createResponse := `{
		"error_message": "validation failure list:\nvalidation failure list:\nid in body must be of type uuid: \"496a3a28-2247-4474-bfa0-ce214a66c07aa\""
	}`

	accountRequest := dto.RequestDto{
		Data: dto.DataRequestDto{
			ID:             "396a3a28-2247-4474-bfa0-ce214a66c07a",
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

	expectedURL := config.AccountAPIURL

	mockPosterRestClient := NewMockRestPoster(ctrl)
	mockPosterRestClient.EXPECT().MakePostRequest(expectedURL, accountRequest).
		Return([]byte(createResponse), http.StatusBadRequest, nil).Times(1)

	actualAccountResponse := dto.ResponseDto{}

	creatorAdapter := creator.NewAdapter(mockPosterRestClient)

	// WHEN
	err := creatorAdapter.CreateAccount(accountRequest, &actualAccountResponse)

	// THEN
	assert.NotNil(t, err)
	assert.Empty(t, actualAccountResponse)
}

func TestShouldReturnErrorWhenResponseJSONNotValid(t *testing.T) {
	t.Parallel()

	// GIVEN
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	createResponse := `{
	  "data": {
		"type": "accounts",
		"id": "496a3a28-2247-4474-bfa0-ce214a66c07a",
		"organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
	  }
	}`

	accountRequest := dto.RequestDto{
		Data: dto.DataRequestDto{
			ID:             "396a3a28-2247-4474-bfa0-ce214a66c07a",
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

	expectedURL := config.AccountAPIURL

	mockPosterRestClient := NewMockRestPoster(ctrl)
	mockPosterRestClient.EXPECT().MakePostRequest(expectedURL, accountRequest).
		Return([]byte(createResponse), http.StatusCreated, nil).Times(1)

	actualAccountResponse := dto.ResponseDto{}

	creatorAdapter := creator.NewAdapter(mockPosterRestClient)

	// WHEN
	err := creatorAdapter.CreateAccount(accountRequest, &actualAccountResponse)

	// THEN
	assert.NotNil(t, err)
	assert.Empty(t, actualAccountResponse)
}

func TestShouldReturnErrorWhenRestClientReturnError(t *testing.T) {
	t.Parallel()

	// GIVEN
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accountRequest := dto.RequestDto{
		Data: dto.DataRequestDto{
			ID:             "396a3a28-2247-4474-bfa0-ce214a66c07a",
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

	expectedURL := config.AccountAPIURL

	mockPosterRestClient := NewMockRestPoster(ctrl)
	mockPosterRestClient.EXPECT().MakePostRequest(expectedURL, accountRequest).
		Return([]byte(""), http.StatusInternalServerError, fmt.Errorf("mock error")).Times(1)

	actualAccountResponse := dto.ResponseDto{}

	creatorAdapter := creator.NewAdapter(mockPosterRestClient)

	// WHEN
	err := creatorAdapter.CreateAccount(accountRequest, &actualAccountResponse)

	// THEN
	assert.NotNil(t, err)
	assert.Empty(t, actualAccountResponse)
}
