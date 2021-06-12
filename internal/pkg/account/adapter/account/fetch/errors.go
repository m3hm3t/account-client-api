package fetch

import (
	"fmt"
)

func generateError(status int, body []byte) error {
	return fmt.Errorf("response status: %d, body: %s", status, body)
}
