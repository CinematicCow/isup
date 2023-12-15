.PHONY: run build tidy clean build-mac build-windows build-linux

run:
	go run cmd/main.go

build:
	go build -o bin/isup cmd/main.go

tidy:
	go mod tidy

clean:
	rm -rf bin/

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/isup-mac cmd/main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/isup-windows.exe cmd/main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/isup-linux cmd/main.go
