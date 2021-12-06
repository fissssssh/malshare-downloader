package malshare

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// HashList struct for unmarshal general hash fields
type HashList struct {
	Md5    string `json:"md5,omitempty"`
	Sha1   string `json:"sha1,omitempty"`
	Sha256 string `json:"sha256,omitempty"`
}

// SearchDetails return searching result
type SearchDetails struct {
	HashList
	TypeSample string `json:"type,omitempty"`
	Added      uint64 `json:"added,omitempty"`
	Source     string `json:"source,omitempty"`
	YaraHits   struct {
		Yara []string `json:"yara,omitempty"`
	} `json:"yarahits,omitempty"`
	Parentfiles []interface{} `json:"parentfiles,omitempty"`
	Subfiles    []interface{} `json:"subfiles,omitempty"`
}

const baseUrl = "https://www.malshare.com"

// GetSearchResult return details form search sample hashes, sources and file names
func GetSearchResult(str string) (*[]SearchDetails, error) {
	urlTmp := fmt.Sprintf("%s/api.php?api_key=%%s&action=search&query=%s", baseUrl, str)
	data, err := request(urlTmp)
	if err != nil {
		return nil, err
	}
	var searches []SearchDetails
	err = json.Unmarshal(data, &searches)
	if err != nil {
		return nil, err
	}
	return &searches, nil
}

// DownloadFileFromHash return file for specific hash
func DownloadFileFromHash(hash string) ([]byte, error) {
	urlTmp := fmt.Sprintf("%s/api.php?api_key=%%s&action=getfile&hash=%s", baseUrl, hash)
	data, err := request(urlTmp)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func request(urlTmp string) ([]byte, error) {
	for {
		apiKey, err := getApiKey()
		if err != nil {
			return nil, err
		}
		url := fmt.Sprintf(urlTmp, apiKey)
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		if isOverLimit(body) {
			continue
		}
		return body, nil
	}
}

func isOverLimit(data []byte) bool {
	if strings.HasPrefix(string(data), "Error: Over Request Limit.") {
		log.Print("apikey is over limit")
		removeApiKey()
		return true
	}
	return false
}
