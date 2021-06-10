package test

import (
	"bytes"
	"github.com/m3hm3t/account-client-api/internal/config"
	"github.com/m3hm3t/account-client-api/internal/pkg/rest/adapter/account/deleter"
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

func (s *DeleteAccountTestSuite) TestShouldDeleteAccountWhenAccountIsCreated() {
	// GIVEN
	if mockErr := s.setMockAccountAPIReturnSuccess(); mockErr != nil {
		s.Error(mockErr)
	}

	config.AccountAPIURL = MockServerURL


	accountID := "MockAccountID"
	version := "0"

	deleteAdapter := deleter.ProvideAccountDeleter2()

	expectedResponseStatus := http.StatusNoContent

	// WHEN

	actualResponseStatus, err := deleteAdapter.DeleteAccount(accountID, version)

	// THEN
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), expectedResponseStatus, actualResponseStatus)
}

func (s *DeleteAccountTestSuite) setMockAccountAPIReturnSuccess() error {
	client := &http.Client{}

	body := `{
				"httpRequest": {
					"method": "DELETE"
				},
				"httpResponse": {
					"statusCode": 204,
					"body": {
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
