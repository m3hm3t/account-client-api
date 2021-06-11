package test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/m3hm3t/account-client-api/internal/config"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/adapter/account/dto"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/adapter/account/fetcher"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestShouldFetchAccountWhenAccountIsCreated(t *testing.T) {
	t.Parallel()

	// GIVEN
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fetchResponse := `{
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

	accountID := "MockAccountID"
	fetcherURL := config.AccountAPIURL + "/" + accountID

	mockGetterRestClient := NewMockRestGetter(ctrl)
	mockGetterRestClient.EXPECT().MakeGetRequest(fetcherURL).Return([]byte(fetchResponse), http.StatusOK, nil).Times(1)

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
				BIC: "NWBKGB22",
				Name:                    []string{"Samantha Holder"},
				AlternativeName:         []string{"Sam Holder"},
				AccountClassification:   "Personal",
				JointAccount:            false,
				AccountMatchingOptOut:   false,
				SecondaryIdentification: "A1B2C3D4",
			},
		},
	}

	fetchAdapter := fetcher.NewAdapter(mockGetterRestClient)

	// WHEN
	err := fetchAdapter.FetchAccount(accountID, &actualAccountResponse)

	// THEN
	assert.Nil(t, err)
	assert.Equal(t, expectedAccountResponse, actualAccountResponse)
}

func TestShouldReturnErrorWhenAccountIsNotCreated(t *testing.T) {
	t.Parallel()

	// GIVEN
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fetchResponse := `{
		"error_message": "record 396a3a28-2247-4474-bfa0-ce214a66c07b does not exist"
	}`

	accountID := "MockAccountID"
	fetcherURL := config.AccountAPIURL + "/" + accountID

	mockGetterRestClient := NewMockRestGetter(ctrl)
	mockGetterRestClient.EXPECT().MakeGetRequest(fetcherURL).Return([]byte(fetchResponse), http.StatusNotFound, nil).Times(1)

	actualAccountResponse := dto.ResponseDto{}

	fetchAdapter := fetcher.NewAdapter(mockGetterRestClient)

	// WHEN
	err := fetchAdapter.FetchAccount(accountID, &actualAccountResponse)

	// THEN
	assert.NotNil(t, err)
	assert.Empty(t, actualAccountResponse)
}

func TestShouldReturnErrorWhenAccountIDIsInvalid(t *testing.T) {
	t.Parallel()

	// GIVEN
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fetchResponse := `{
		"error_message": "id is not a valid uuid"
	}`

	accountID := "MockAccountID"
	fetcherURL := config.AccountAPIURL + "/" + accountID

	mockGetterRestClient := NewMockRestGetter(ctrl)
	mockGetterRestClient.EXPECT().MakeGetRequest(fetcherURL).Return([]byte(fetchResponse), http.StatusBadRequest, nil).Times(1)

	actualAccountResponse := dto.ResponseDto{}

	fetchAdapter := fetcher.NewAdapter(mockGetterRestClient)

	// WHEN
	err := fetchAdapter.FetchAccount(accountID, &actualAccountResponse)

	// THEN
	assert.NotNil(t, err)
	assert.Empty(t, actualAccountResponse)
}

func TestShouldReturnErrorWhenAccountAPIReturnsInvalidJSON(t *testing.T) {
	t.Parallel()

	// GIVEN
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fetchResponse := `{
	  "data": {
		"type": "accounts",
		"id": "496a3a28-2247-4474-bfa0-ce214a66c07a",
		"organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
	  }
	}`

	accountID := "MockAccountID"
	fetcherURL := config.AccountAPIURL + "/" + accountID

	mockGetterRestClient := NewMockRestGetter(ctrl)
	mockGetterRestClient.EXPECT().MakeGetRequest(fetcherURL).Return([]byte(fetchResponse), http.StatusOK, nil).Times(1)

	actualAccountResponse := dto.ResponseDto{}

	fetchAdapter := fetcher.NewAdapter(mockGetterRestClient)

	// WHEN
	err := fetchAdapter.FetchAccount(accountID, &actualAccountResponse)

	// THEN
	assert.NotNil(t, err)
	assert.Empty(t, actualAccountResponse)
}

func TestShouldReturnErrorWhenRestClientReturnError(t *testing.T) {
	t.Parallel()

	// GIVEN
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accountID := "MockAccountID"
	fetcherURL := config.AccountAPIURL + "/" + accountID

	mockGetterRestClient := NewMockRestGetter(ctrl)
	mockGetterRestClient.EXPECT().MakeGetRequest(fetcherURL).Return(nil, http.StatusBadRequest, fmt.Errorf("mock error")).Times(1)

	actualAccountResponse := dto.ResponseDto{}

	fetchAdapter := fetcher.NewAdapter(mockGetterRestClient)

	// WHEN
	err := fetchAdapter.FetchAccount(accountID, &actualAccountResponse)

	// THEN
	assert.NotNil(t, err)
	assert.Empty(t, actualAccountResponse)
}
