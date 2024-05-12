# Define variables
BINARY_NAME=paynex
DEBUG_BINARY_NAME=paynex-debug
BIN_DIR=bin

# Set Go module
GO111MODULE=on

# Build your program
build:
	@mkdir -p $(BIN_DIR)
	@echo "Building $(BINARY_NAME)..."
	@go build -v -o $(BIN_DIR)/$(BINARY_NAME)
	@echo "Build completed: $(BIN_DIR)/$(BINARY_NAME)"

# Clean the project
clean:
	@echo "Cleaning $(BIN_DIR)..."
	@rm -rf $(BIN_DIR)/*
	@echo "Cleaned $(BIN_DIR)"

# Run your program
run: build
	@echo "Running $(BINARY_NAME)..."
	@./$(BIN_DIR)/$(BINARY_NAME)

# Debug your program with Delve
debug: build
	@if ! command -v dlv &> /dev/null; then \
		echo "Delve (dlv) not found. Please install it first."; \
		exit 1; \
	fi
	@echo "Debugging $(BINARY_NAME)..."
	@dlv exec $(BIN_DIR)/$(BINARY_NAME)

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...
	@echo "Tests completed."

# Display available commands
help:
	@echo "Usage: make [command]"
	@echo ""
	@echo "Available commands:"
	@echo "  build    - Build the binary"
	@echo "  clean    - Clean the binary files"
	@echo "  run      - Build and run the program"
	@echo "  debug    - Build and debug the program with Delve"
	@echo "  test     - Run tests"
	@echo "  help     - Display this help message"
