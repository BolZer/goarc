test:
	go test -v ./...

run:
	go run main.go

download_deps:
	go mod vendor

build:
	go build -o goarc.exe main.go