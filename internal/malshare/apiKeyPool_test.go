package malshare

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestApiKeyPool(t *testing.T) {
	// read keys_file
	file, err := os.Open("apikey_test")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	apiKeySc := bufio.NewScanner(file)
	var apiKeys []string
	for apiKeySc.Scan() {
		line := apiKeySc.Text()
		if line != "" {
			apiKeys = append(apiKeys, line)
		}
	}
	if len(apiKeys) == 0 {
		log.Fatalln("no api keys")
	}
	// init api key pool
	InitApiKeyPool(apiKeys...)
	for i := 0; i < len(apiKeys); i++ {
		key, err := getApiKey()
		if err != nil {
			t.Fail()
		}
		if key != apiKeys[i] {
			t.Fail()
		}
		removeApiKey()
	}
	_, err = getApiKey()
	if err == nil {
		t.Fail()
	}
	removeApiKey()
	removeApiKey()
	removeApiKey()
	removeApiKey()
	removeApiKey()
	removeApiKey()
	removeApiKey()
	_, err = getApiKey()
	if err == nil {
		t.Fail()
	}
}
