package malshare

import (
	"errors"
)

var apiKeys []string

func InitApiKeyPool(keys ...string) {
	apiKeys = make([]string, len(keys))
	copy(apiKeys, keys)
}

func removeApiKey(key string) {
	apiKeys = apiKeys[1:]
}

func getApiKey() (string, error) {
	if len(apiKeys) > 0 {
		return apiKeys[0], nil
	}
	return "", errors.New("all apikeys are over limit")
}
