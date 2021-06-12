package test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/m3hm3t/account-client-api/internal/config"
	"github.com/m3hm3t/account-client-api/internal/pkg/account/adapter/account/delete"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestShouldDeleteAccountWhenAccountIsCreated(t *testing.T) {
	t.Parallel()

	// GIVEN
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accountID := "MockAccountID"
	version := "0"
	deleterURL := config.AccountAPIURL + "/" + accountID + "/?version=" + version

	expectedResponseStatus := http.StatusNoContent

	mockDeleterRestClient := NewMockRestDeleter(ctrl)
	mockDeleterRestClient.EXPECT().MakeDeleteRequest(deleterURL).Return(expectedResponseStatus, nil).Times(1)

	deleteAdapter := delete.NewAdapter(mockDeleterRestClient)

	// WHEN
	actualResponseStatus, err := deleteAdapter.DeleteAccount(accountID, version)

	// THEN
	assert.Nil(t, err)
	assert.Equal(t, expectedResponseStatus, actualResponseStatus)
}

func TestShouldReturnErrorWhenRestClientReturnError(t *testing.T) {
	t.Parallel()

	// GIVEN
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accountID := "MockAccountID"
	version := "0"
	deleterURL := config.AccountAPIURL + "/" + accountID + "/?version=" + version

	expectedResponseStatus := http.StatusInternalServerError

	mockDeleterRestClient := NewMockRestDeleter(ctrl)
	mockDeleterRestClient.EXPECT().MakeDeleteRequest(deleterURL).Return(expectedResponseStatus, fmt.Errorf("mock error")).Times(1)

	deleteAdapter := delete.NewAdapter(mockDeleterRestClient)

	// WHEN
	actualResponseStatus, err := deleteAdapter.DeleteAccount(accountID, version)

	// THEN
	assert.NotNil(t, err)
	assert.Equal(t, expectedResponseStatus, actualResponseStatus)
}
