package getter_async

import (
	"github.com/m3hm3t/account-client-api/internal/config"
	"strconv"
)

func ProvideRestGetterAsyncClient(bufferSize int) *RestGetterAsyncClient {
	timeOutInMilliseconds, err := strconv.ParseInt(config.AccountAPIGetTimeoutInMilliseconds, 10, 64)
	if err != nil {
		panic(err)
	}

	return NewRestGetterAsyncClient(timeOutInMilliseconds, bufferSize)
}
