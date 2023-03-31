build:
	go build -o struct-to-json *.go

build_linux:
	GOOS=linux GOARCH=amd64 go build -o struct-to-json *.go

run:
	make build
	./struct-to-json