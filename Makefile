build: clean
	go build -o kasa-stocks *.go
	ls -l

run: build
	./kasa-stocks

clean:
	rm -f ./kasa-stocks