package test

import (
	"github.com/google/uuid"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/adapter/account/delete"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type DeleteAccountTestSuite struct {
	suite.Suite
}

func TestDeleteAccountTestSuite(t *testing.T) {
	suite.Run(t, new(DeleteAccountTestSuite))
}

func (s *DeleteAccountTestSuite) TestShouldReturnErrorWhenAccountIDInvalid() {
	// GIVEN
	accountID := "InvalidAccountID"
	version := "0"

	deleteAdapter := delete.ProvideAccountDeleter()

	expectedResponseStatus := http.StatusBadRequest

	// WHEN

	actualResponseStatus, err := deleteAdapter.DeleteAccount(accountID, version)

	// THEN
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), expectedResponseStatus, actualResponseStatus)
}

func (s *DeleteAccountTestSuite) TestShouldReturnErrorWhenAccountNotExisting() {
	// GIVEN
	accountID := uuid.NewString()
	version := "0"

	deleteAdapter := delete.ProvideAccountDeleter()

	expectedResponseStatus := http.StatusNotFound

	// WHEN

	actualResponseStatus, err := deleteAdapter.DeleteAccount(accountID, version)

	// THEN
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), expectedResponseStatus, actualResponseStatus)
}

func (s *DeleteAccountTestSuite) TestShouldReturnErrorWhenVersionInvalid() {
	// GIVEN
	accountID := "InvalidAccountID"
	version := "a"

	deleteAdapter := delete.ProvideAccountDeleter()

	expectedResponseStatus := http.StatusBadRequest

	// WHEN

	actualResponseStatus, err := deleteAdapter.DeleteAccount(accountID, version)

	// THEN
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), expectedResponseStatus, actualResponseStatus)
}
