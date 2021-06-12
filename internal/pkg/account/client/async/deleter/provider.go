package deleter

import (
	"github.com/m3hm3t/account-client-api/internal/config"
	"strconv"
)

func ProvideRestDeleterAsyncClient(bufferSize int) *RestDeleterAsyncClient {
	timeOutInMilliseconds, err := strconv.ParseInt(config.AccountAPIGetTimeoutInMilliseconds, 10, 64)
	if err != nil {
		panic(err)
	}

	return NewRestDeleterAsyncClient(timeOutInMilliseconds, bufferSize)
}
