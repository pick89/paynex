# Define variables
BINARY_NAME=paynex
DEBUG_BINARY_NAME=paynex-debug

# Build your program
build:
	@mkdir -p bin
	@go build -o bin/paynex

# Clean the project
clean:
	@rm -rf bin/*

# Run your program
run: build
	@./bin/$(BINARY_NAME)

# Debug your program with Delve
debug:
	@dlv debug bin/paynex
# Run tests
test:
	@go test -v ./...
