package deleter

import deleterclient "github.com/m3hm3t/account-client-api/internal/pkg/account/client/deleter"

func ProvideAccountDeleter() AccountDeleter {
	deleter := deleterclient.ProvideRestDeleterClient()
	return NewAdapter(deleter)
}
