default:
	@mkdir -p bin
	@go build -o bin/scrape-flickr-albums cmd/scan-collections/main.go
