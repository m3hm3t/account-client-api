package creator

import (
	"github.com/m3hm3t/account-client-api/internal/pkg/account/client/poster"
)

func ProvideAccountCreator() AccountCreator {
	posterClient := poster.ProvideAPosterClient()
	return NewAdapter(posterClient)
}
