package main

import (
	"bufio"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"sync"

	"malshare-downloader/internal/malshare"
	"malshare-downloader/utils"
)

func main() {
	var apiKey, hashFilesDir, output, _type, yara string
	flag.StringVar(&apiKey, "api", "", "API key MalShare")
	flag.StringVar(&hashFilesDir, "source", "hash_files", "directory of hash files")
	flag.StringVar(&output, "o", "mal_files", "output directory")
	flag.StringVar(&_type, "type", "", "Download file type")
	flag.StringVar(&yara, "yara", "", "Yarp")
	flag.Parse()
	files, err := ioutil.ReadDir(hashFilesDir)
	if err != nil {
		log.Fatalf("open source directory failed: %s", err)
	}
	wg := sync.WaitGroup{}
	concurrent := make(chan struct{}, 10)
	for _, file := range files {
		hashFilePath := path.Join(hashFilesDir, file.Name())
		fs, err := os.Open(hashFilePath)
		if err != nil {
			log.Printf("open hash file %s error: %s", hashFilePath, err)
			continue
		}
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
					log.Printf("create directory %s failed: %s", filepath, err)
					continue
				}
			}
			filepath = path.Join(filepath, hash)
			wg.Add(1)
			concurrent <- struct{}{}
			go func(hash string, filepath string) {
				searchs, err := malshare.GetSearchResult(apiKey, hash)
				if err != nil {
					log.Printf("get stored file details file with hash %s failed: %s", hash, err)
				}
				if len(*searchs) == 0 {
					<-concurrent
					wg.Done()
					return
				}
				details := (*searchs)[0]
				if _type != "" && details.TypeSample != _type {
					<-concurrent
					wg.Done()
					return
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
						<-concurrent
						wg.Done()
						return
					}
				}
				file, err := malshare.DownloadFileFromHash(apiKey, hash)
				if err != nil {
					log.Printf("download file with hash %s failed: %s", hash, err)
					<-concurrent
					wg.Done()
				}
				fs, err := os.Create(filepath)
				if err != nil {
					log.Printf("create file %s failed: %s", filepath, err)
					<-concurrent
					wg.Done()
				}
				defer fs.Close()
				fs.Write(file)
				<-concurrent
				wg.Done()
			}(hash, filepath)
		}
	}
	wg.Wait()
}
