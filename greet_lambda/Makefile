.PHONY: clean build test dist

clean:
	rm -f bootstrap function.zip

build: clean
	go get
	go mod tidy
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -tags lambda.norpc -o bootstrap main.go

dist: build
	zip function.zip bootstrap
	rm bootstrap

test:
	go test -v ./...
