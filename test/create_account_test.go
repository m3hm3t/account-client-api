package test

import (
	"github.com/google/uuid"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/adapter/account/creator"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/adapter/account/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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

	// When
	err := creatorAdapter.CreateAccount(accountRequest, &actualAccountResponse)

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), accountRequest.Data.ID, actualAccountResponse.Data.ID)
	assert.Equal(s.T(), accountRequest.Data.Type, actualAccountResponse.Data.Type)
	assert.Equal(s.T(), accountRequest.Data.OrganisationID, actualAccountResponse.Data.OrganisationID)
	assert.Equal(s.T(), accountRequest.Data.Attributes.Country, actualAccountResponse.Data.Attributes.Country)
	assert.Equal(s.T(), accountRequest.Data.Attributes.BaseCurrency, actualAccountResponse.Data.Attributes.BaseCurrency)
	assert.Equal(s.T(), accountRequest.Data.Attributes.BankID, actualAccountResponse.Data.Attributes.BankID)
	assert.Equal(s.T(), accountRequest.Data.Attributes.BankIDCode, actualAccountResponse.Data.Attributes.BankIDCode)
	assert.Equal(s.T(), accountRequest.Data.Attributes.BIC, actualAccountResponse.Data.Attributes.BIC)
	assert.Equal(s.T(), accountRequest.Data.Attributes.Name, actualAccountResponse.Data.Attributes.Name)
	assert.Equal(s.T(), accountRequest.Data.Attributes.AlternativeName,
		actualAccountResponse.Data.Attributes.AlternativeName)
	assert.Equal(s.T(), accountRequest.Data.Attributes.AccountClassification,
		actualAccountResponse.Data.Attributes.AccountClassification)
	assert.Equal(s.T(), accountRequest.Data.Attributes.SecondaryIdentification,
		actualAccountResponse.Data.Attributes.SecondaryIdentification)
	assert.Equal(s.T(), accountRequest.Data.Attributes.JointAccount,
		actualAccountResponse.Data.Attributes.JointAccount)
}

func (s *CreateAccountTestSuite) TestShouldReturnErrorWhenAccountIDInvalid() {
	// Given
	accountRequest := dto.RequestDto{
		Data: dto.DataRequestDto{
			ID:             "InvalidAccountID",
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

	// When
	err := creatorAdapter.CreateAccount(accountRequest, &actualAccountResponse)

	// Then
	assert.NotNil(s.T(), err)
	assert.Contains(s.T(), err.Error(), "error_message")
	assert.Empty(s.T(), actualAccountResponse)
}

func (s *CreateAccountTestSuite) TestShouldReturnErrorWhenCountryMissing() {
	// Given
	accountRequest := dto.RequestDto{
		Data: dto.DataRequestDto{
			ID:             uuid.NewString(),
			Type:           "accounts",
			OrganisationID: "496a3a28-2247-4474-bfa0-ce214a66c07b",
			Attributes: dto.AttributesRequestDto{
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

	// When
	err := creatorAdapter.CreateAccount(accountRequest, &actualAccountResponse)

	// Then
	assert.NotNil(s.T(), err)
	assert.Contains(s.T(), err.Error(), "error_message")
	assert.Empty(s.T(), actualAccountResponse)
}
