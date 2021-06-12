package fetch

import (
	"github.com/m3hm3t/account-client-api/internal/pkg/account/adapter/account/dto"
)

type AccountFetcher interface {
	FetchAccount(accountID string, accountResponse *dto.ResponseDto) error
}
