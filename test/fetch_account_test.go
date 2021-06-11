package test

import (
	"github.com/google/uuid"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/adapter/account/dto"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/adapter/account/fetcher"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type FetchAccountTestSuite struct {
	suite.Suite
}

func TestFetchAccountTestSuite(t *testing.T) {
	suite.Run(t, new(FetchAccountTestSuite))
}

func (s *FetchAccountTestSuite) TestShouldReturnErrorWhenAccountIDInvalid() {
	// GIVEN
	accountID := "InvalidAccountID"

	actualAccountResponse := dto.ResponseDto{}

	fetcherAdapter := fetcher.ProvideAccountFetcher()

	// WHEN
	err := fetcherAdapter.FetchAccount(accountID, &actualAccountResponse)

	// THEN
	assert.NotNil(s.T(), err)
	assert.Contains(s.T(), err.Error(), "error_message")
	assert.Empty(s.T(), actualAccountResponse)
}

func (s *FetchAccountTestSuite) TestShouldReturnErrorWhenAccountNotCreated() {
	// GIVEN
	accountID := uuid.NewString()

	actualAccountResponse := dto.ResponseDto{}

	fetcherAdapter := fetcher.ProvideAccountFetcher()

	// WHEN
	err := fetcherAdapter.FetchAccount(accountID, &actualAccountResponse)

	// THEN
	assert.NotNil(s.T(), err)
	assert.Contains(s.T(), err.Error(), "error_message")
	assert.Empty(s.T(), actualAccountResponse)
}
