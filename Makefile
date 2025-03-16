run:
	go run cmd/server/main.go

build:
	CGO_ENABLED=0 go build -o bin/recipio-server -ldflags "-s -w" cmd/server/main.go

linux:
	GOOS=linux $(MAKE) build

clean:
	rm -rf bin/
