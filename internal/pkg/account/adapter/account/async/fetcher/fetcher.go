package fetcher

import (
	"github.com/m3hm3t/account-client-api/internal/pkg/account/adapter/account/dto"
)

type AccountFetcherAsync interface {
	FetchAll(accountIDs []string, accountResponses []dto.ResponseDto) error
}
