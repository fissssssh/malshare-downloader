SHELL=/bin/bash
OUTPUTDIR:=build
VERSION:=v0.0.0
.PHONY: clean build_spider build_downloader
rebuild: clean build_spider build_downloader

# clean execute files.
clean:
	@echo "Cleaning"
	rm -rf $(OUTPUTDIR)/$(VERSION) vendor

# buiil windows, linux and macos executed files.
build_spider: cmd/spider/main.go
	@echo "Build for linux"
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o $(OUTPUTDIR)/$(VERSION)/spider_linux_amd64_$(VERSION) $<
	@echo "Build for macos"
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -o $(OUTPUTDIR)/$(VERSION)/spider_macos_amd64_$(VERSION) $<
	@echo "Build for windows"
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -o $(OUTPUTDIR)/$(VERSION)/spider_windows_amd64_$(VERSION).exe $<

build_downloader: cmd/downloader/main.go
	@echo "Build for linux"
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o $(OUTPUTDIR)/$(VERSION)/downloader_linux_amd64_$(VERSION) $<
	@echo "Build for macos"
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -o $(OUTPUTDIR)/$(VERSION)/downloader_macos_amd64_$(VERSION) $<
	@echo "Build for windows"
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -o $(OUTPUTDIR)/$(VERSION)/downloader_windows_amd64_$(VERSION).exe $<