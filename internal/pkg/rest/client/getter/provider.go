package getter

import (
	"github.com/m3hm3t/account-client-api/internal/config"
	"strconv"
)

func ProvideRestGetterClient() RestGetter {
	timeOutInMilliseconds, err := strconv.ParseInt(config.AccountAPIGetTimeoutInMilliseconds, 10, 64)
	if err != nil {
		panic(err)
	}

	return NewRestGetterClient(timeOutInMilliseconds)
}
