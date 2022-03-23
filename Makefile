version := 0.1

build: clean build_linux build_windows build_mac
	go build -o kasa-stocks *.go

build_linux:
	GOOS=linux GOARCH=amd64 go build -o bin/kasa-stocks_$(version)_linux_amd64 *.go
	GOOS=linux GOARCH=arm64 go build -o bin/kasa-stocks_$(version)_linux_arm64 *.go

build_windows:
	GOOS=windows GOARCH=amd64 go build -o bin/kasa-stocks_$(version)_windows_arm64 *.go

build_mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/kasa-stocks_$(version)_darwin_amd64 *.go
	GOOS=darwin GOARCH=arm64 go build -o bin/kasa-stocks_$(version)_darwin_arm64 *.go

run: build
	./kasa-stocks

clean:
	rm -f bin/kasa-stocks_*