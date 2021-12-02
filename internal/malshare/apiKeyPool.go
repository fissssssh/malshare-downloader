package malshare

import (
	"errors"
)

var apiKeys []string

func InitApiKeyPool(keys ...string) {
	apiKeys = make([]string, len(keys))
	copy(apiKeys, keys)
}

func removeApiKey() {
	if len(apiKeys) > 0 {
		apiKeys = apiKeys[1:]
	}
}

func getApiKey() (string, error) {
	if len(apiKeys) > 0 {
		return apiKeys[0], nil
	}
	return "", errors.New("all apikeys are over limit")
}
