package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

const hashTmpl = "https://www.malshare.com/daily/%04d-%02d-%02d/malshare_fileList.%02d-%02d-%02d.txt"

func main() {
	var start, end int64
	var output string
	flag.Int64Var(&start, "start", -1, "start")
	flag.Int64Var(&end, "end", -1, "end")
	flag.StringVar(&output, "o", "hash_files", "output directory")
	flag.Parse()
	if start == -1 {
		log.Fatalf("start cannot be empty")
	}
	var st, et time.Time
	st = time.UnixMilli(start)
	if end == -1 {
		et = time.Now()
	} else {
		et = time.UnixMilli(end)
	}
	if info, err := os.Stat(output); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(output, os.ModePerm)
		}
	} else if !info.IsDir() {
		log.Fatalf("exist a file with the same name: %s", output)
	}
	wg := sync.WaitGroup{}
	concurrent := make(chan struct{}, 10)
	for st.Before(et) {
		year, month, day := st.Date()
		url := fmt.Sprintf(hashTmpl, year, month, day, year, month, day)
		wg.Add(1)
		concurrent <- struct{}{}
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				log.Printf("download error from %s: %s", url, err)
				<-concurrent
				wg.Done()
				return
			}
			if resp.StatusCode != 200 {
				log.Printf("download error from %s with statuscode: %d", url, resp.StatusCode)
				<-concurrent
				wg.Done()
				return
			}
			filename := fmt.Sprintf("%02d-%02d-%02d.txt", year, month, day)
			fs, err := os.Create(path.Join(output, filename))
			if err != nil {
				log.Printf("create file named %s error: %s", filename, err)
				<-concurrent
				wg.Done()
				return
			}
			defer fs.Close()
			io.Copy(fs, resp.Body)
			<-concurrent
			wg.Done()
		}(url)
		st = st.Add(24 * time.Hour)
	}
	wg.Wait()
}
