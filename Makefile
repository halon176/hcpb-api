run:
	@go run .

update:
	@go get -u ./...

build:
	@go build -ldflags "-w -s" -o bin/$(APP_NAME) .