build: clean
	go build -o kasa-stocks *.go

run: build
	./kasa-stocks

clean:
	rm -f ./kasa-stocks