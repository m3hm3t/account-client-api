package deleter

import deleterclient "github.com/m3hm3t/account-client-api/internal/pkg/rest/client/deleter"

func ProvideAccountDeleter(deleter deleterclient.RestDeleter) AccountDeleter {
	return NewAdapter(deleter)
}

func ProvideAccountDeleter2() AccountDeleter {
	deleter := deleterclient.ProvideRestDeleterClient()
	return NewAdapter(deleter)
}
