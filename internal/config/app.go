package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	AccountAPIURL                         string
	AccountAPIGetTimeoutInMilliseconds    string
	AccountAPIPostTimeoutInMilliseconds   string
	AccountAPIDeleteTimeoutInMilliseconds string
)

func init() {

	_ = godotenv.Load()

	AccountAPIURL = GetEnvWithDefault(
		"ACCOUNT_API_URL",
		"http://localhost:8080/v1/organisation/accounts",
	)

	AccountAPIGetTimeoutInMilliseconds = GetEnvWithDefault(
		"ACCOUNT_API_GET_TIMEOUT_IN_MILLISECONDS",
		"3000",
	)

	AccountAPIPostTimeoutInMilliseconds = GetEnvWithDefault(
		"ACCOUNT_API_POST_TIMEOUT_IN_MILLISECONDS",
		"3000",
	)

	AccountAPIDeleteTimeoutInMilliseconds = GetEnvWithDefault(
		"ACCOUNT_API_DELETE_TIMEOUT_IN_MILLISECONDS",
		"3000",
	)
}

func GetEnvWithDefault(key string, defaultValue string) string {
	var env string
	if value, ok := os.LookupEnv(key); !ok {
		env = defaultValue
	} else {
		env = value
	}
	return env
}
