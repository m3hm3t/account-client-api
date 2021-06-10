package poster

import (
	"github.com/m3hm3t/account-client-api/internal/config"
	"strconv"
)

func ProvideAPosterClient() RestPoster {
	timeOutInMilliseconds, err := strconv.ParseInt(config.AccountAPIPostTimeoutInMilliseconds, 10, 64)
	if err != nil {
		panic(err)
	}

	return NewRestPosterClient(timeOutInMilliseconds)
}
