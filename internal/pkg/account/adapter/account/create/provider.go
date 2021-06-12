package create

import (
	"github.com/m3hm3t/account-client-api/internal/pkg/account/client/post"
)

func ProvideAccountCreator() AccountCreator {
	posterClient := post.ProvideAPosterClient()
	return NewAdapter(posterClient)
}
