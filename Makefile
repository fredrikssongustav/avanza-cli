GOCMD=go
GOBUILD=$(GOCMD) build -o
BINARY_NAME=avanza-cli

build:
	@CGO_ENABLED=0 $(GOBUILD) $(BINARY_NAME) -v github.com/fredrikssongustav/avanza-cli

## Clean output files.
clean:
	@rm -f $(BINARY_NAME)

fmt:
	@go fmt ./...

