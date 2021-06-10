package fetcher

import "github.com/m3hm3t/account-client-api/internal/pkg/account/client/getter"

func ProvideAccountFetcher() AccountFetcher {
	restGetter := getter.ProvideRestGetterClient()
	return NewAdapter(restGetter)
}
