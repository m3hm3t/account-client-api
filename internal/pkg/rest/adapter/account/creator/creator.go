package creator

import (
	"github.com/m3hm3t/account-client-api/internal/pkg/rest/adapter/account/dto"
)

type AccountCreator interface {
	CreateAccount(accountRequest dto.RequestDto, accountResponse *dto.ResponseDto) error
}
