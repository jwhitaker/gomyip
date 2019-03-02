.PHONY: default
defualt: help ;

BINARY_NAME=gomyip

## build: Build the application
build:
	go build -o $(BINARY_NAME) -v

## clean: Clean build files
clean:
	go clean


.PHONY: help
help: Makefile
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'