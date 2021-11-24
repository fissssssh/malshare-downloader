package malshare

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

const baseUrl = "http://www.malshare.com"

// GetSearchResult return details form search sample hashes, sources and file names
func GetSearchResult(apiKey string, str string) (*[]SearchDetails, error) {
	url := fmt.Sprintf("%s/api.php?api_key=%s&action=search&query=%s", baseUrl, apiKey, str)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var searches []SearchDetails
	err = json.Unmarshal(body, &searches)
	if err != nil {
		return nil, err
	}
	return &searches, nil
}

// DownloadFileFromHash return file for specific hash
func DownloadFileFromHash(apiKey string, hash string) ([]byte, error) {
	url := fmt.Sprintf("%s/api.php?api_key=%s&action=getfile&hash=%s", baseUrl, apiKey, hash)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
