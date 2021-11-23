package main

import (
	"bufio"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sync"

	"malshare-downloader/utils"

	"github.com/MonaxGT/gomalshare"
)

func main() {
	apiKeyPtr := flag.String("api", "", "API key MalShare")
	var hashFilesDir, output, t string
	flag.StringVar(&hashFilesDir, "source", "hash_files", "directory of hash files")
	flag.StringVar(&output, "o", "mal_files", "output directory")
	flag.StringVar(&t, "t", "", "Download file type")
	flag.Parse()
	files, err := ioutil.ReadDir(hashFilesDir)
	if err != nil {
		log.Fatalf("open source directory failed: %s", err)
	}
	conf, err := gomalshare.New(*apiKeyPtr, "https://www.malshare.com/")
	if err != nil {
		log.Fatalf("create gomalshare client failed: %s", err)
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
				details, err := conf.GetStoredFileDetails(hash)
				if err != nil {
					log.Printf("get stored file details file with hash %s failed: %s", hash, err)
				}
				if t != "" && details.FType != t {
					<-concurrent
					wg.Done()
					return
				}
				file, err := conf.DownloadFileFromHash(hash)
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
