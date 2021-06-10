package creator

import (
	"github.com/m3hm3t/account-client-api/internal/pkg/rest/client/poster"
)

func ProvideAccountCreator(restPoster poster.RestPoster) AccountCreator {
	return NewAdapter(restPoster)
}

func ProvideAccountCreator2() AccountCreator {
	poster := poster.ProvideAPosterClient()
	return NewAdapter(poster)
}
