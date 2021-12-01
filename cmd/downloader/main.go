package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"malshare-downloader/internal/malshare"
	"malshare-downloader/utils"
)

func main() {
	var keys_file, hashFilesDir, output, _type, yara string
	flag.StringVar(&keys_file, "keys_file", "", "MalShare API key file")
	flag.StringVar(&hashFilesDir, "source", "hash_files", "directory of hash files")
	flag.StringVar(&output, "o", "mal_files", "output directory")
	flag.StringVar(&_type, "type", "", "Download file type")
	flag.StringVar(&yara, "yara", "", "Yarp")
	flag.Parse()
	files, err := ioutil.ReadDir(hashFilesDir)
	if err != nil {
		log.Fatalf("open source directory failed: %s", err)
	}
	if keys_file == "" {
		log.Fatalf("need keys file")
	}
	// read keys_file
	file, err := os.Open(keys_file)
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
	malshare.InitApiKeyPool(apiKeys...)
	for _, file := range files {
		hashFilePath := path.Join(hashFilesDir, file.Name())
		fs, err := os.Open(hashFilePath)
		if err != nil {
			log.Printf("open hash file %s error: %s", hashFilePath, err)
			continue
		}
		log.SetPrefix(fmt.Sprintf("[%s]", file.Name()))
		defer fs.Close()
		sc := bufio.NewScanner(fs)
		for sc.Scan() {
			hash := sc.Text()
			if hash == "" {
				continue
			}
			filepath := path.Join(output, utils.GetFileNameWithoutExt(hashFilePath))
			if !utils.IsExist(filepath) {
				err := os.MkdirAll(filepath, os.ModePerm)
				if err != nil {
					log.Fatalf("create directory %s failed: %s", filepath, err)
					continue
				}
			}
			filepath = path.Join(filepath, hash)
			log.Printf("searching details with hash %s", hash)
			searchs, err := malshare.GetSearchResult(hash)
			if err != nil {
				log.Fatalf("get stored file details file with hash %s failed: %s", hash, err)
			}
			if len(*searchs) == 0 {
				continue
			}
			details := (*searchs)[0]
			if _type != "" && details.TypeSample != _type {
				continue
			}
			if yara != "" {
				matched := false
				for _, v := range details.YaraHits.Yara {
					if strings.Contains(v, yara) {
						matched = true
						break
					}
				}
				if !matched {
					continue
				}
			}
			log.Printf("downloading file with hash %s", hash)
			file, err := malshare.DownloadFileFromHash(hash)
			if err != nil {
				log.Fatalf("download file with hash %s failed: %s", hash, err)
			}
			fs, err := os.Create(filepath)
			if err != nil {
				log.Fatalf("create file %s failed: %s", filepath, err)
			}
			defer fs.Close()
			fs.Write(file)
		}
	}
}
