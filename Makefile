lint:
	gofmt -w .

build:
	go build -o ./bin/retorrent ./main.go
