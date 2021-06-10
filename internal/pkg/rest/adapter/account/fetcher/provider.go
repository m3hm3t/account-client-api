package fetcher

import "github.com/m3hm3t/account-client-api/internal/pkg/rest/client/getter"

func ProvideAccountFetcher(restGetter getter.RestGetter) AccountFetcher {
	return NewAdapter(restGetter)
}
func ProvideAccountFetcher2() AccountFetcher {
	restGetter := getter.ProvideRestGetterClient()
	return NewAdapter(restGetter)
}
