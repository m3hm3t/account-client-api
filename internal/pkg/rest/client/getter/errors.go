package getter

import (
	"errors"
	"fmt"
)

var ErrRestRequestNotSuccess = errors.New("the rest request is not successful")

func RestRequestNotSuccessError(statusCode string, responseBody string) error {
	return fmt.Errorf("%w : %s :  %s", ErrRestRequestNotSuccess, statusCode, responseBody)
}
