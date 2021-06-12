package fetch

import "github.com/m3hm3t/account-client-api/internal/pkg/account/client/get"

func ProvideAccountFetcher() AccountFetcher {
	restGetter := get.ProvideRestGetterClient()
	return NewAdapter(restGetter)
}
