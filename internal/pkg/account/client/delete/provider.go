package delete

import (
	"github.com/m3hm3t/account-client-api/internal/config"
	"strconv"
)

func ProvideRestDeleterClient() RestDeleter {
	timeOutInMilliseconds, err := strconv.ParseInt(config.AccountAPIDeleteTimeoutInMilliseconds, 10, 64)
	if err != nil {
		panic(err)
	}

	return NewRestDeleter(timeOutInMilliseconds)
}
