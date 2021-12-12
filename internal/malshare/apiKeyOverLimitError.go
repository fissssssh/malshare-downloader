package malshare

import "fmt"

type apiKeyOverLimitError struct {
	key string
}

func (err apiKeyOverLimitError) Error() string {
	return fmt.Sprintf("ApiKey %s is over limit", err.key)
}
