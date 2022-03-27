version := 0.1

build: 
	rm -f ./kasa-stocks
	go build -o kasa-stocks *.go

build_all: clean build_linux build_windows build_mac

build_linux:
	GOOS=linux GOARCH=amd64 go build -o bin/linux_amd64/kasa-stocks_$(version) *.go
	sha256sum bin/linux_amd64/kasa-stocks_$(version) > bin/linux_amd64/kasa-stocks_$(version)_sha256.txt
	
	GOOS=linux GOARCH=arm64 go build -o bin/linux_arm64/kasa-stocks_$(version) *.go
	sha256sum bin/linux_arm64/kasa-stocks_$(version) > bin/linux_arm64/kasa-stocks_$(version)_sha256.txt

build_windows:
	GOOS=windows GOARCH=amd64 go build -o bin/windows_amd64/kasa-stocks_$(version) *.go
	sha256sum bin/windows_amd64/kasa-stocks_$(version) > bin/windows_amd64/kasa-stocks_$(version)_sha256.txt

build_mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin_amd64/kasa-stocks_$(version) *.go
	sha256sum bin/darwin_amd64/kasa-stocks_$(version) > bin/darwin_amd64/kasa-stocks_$(version)_sha256.txt

	GOOS=darwin GOARCH=arm64 go build -o bin/darwin_arm64/kasa-stocks_$(version) *.go
	sha256sum bin/darwin_arm64/kasa-stocks_$(version) > bin/darwin_arm64/kasa-stocks_$(version)_sha256.txt

run: build
	./kasa-stocks

clean:
	rm -f bin/linux_amd64/kasa-stocks_*
	rm -f bin/linux_arm64/kasa-stocks_*
	rm -f bin/windows_amd64/kasa-stocks_*
	rm -f bin/darwin_amd64/kasa-stocks_*
	rm -f bin/darwin_arm64/kasa-stocks_*