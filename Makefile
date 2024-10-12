# Define variables
APP_NAME := todoapp
SRC_DIR := ./cmd/todoapp
BIN_DIR := ./bin
GO := go
GO_BUILD := $(GO) build -o $(BIN_DIR)/$(APP_NAME)

# Default target to build the executable
all: build

# Target to build the Go application
build:
	@mkdir -p $(BIN_DIR) # Create the bin directory if it doesn't exist
	$(GO_BUILD) $(SRC_DIR)/main.go

# Clean target to remove the built executable
clean:
	rm -f $(BIN_DIR)/$(APP_NAME)

# Run the application
run: build
	$(BIN_DIR)/$(APP_NAME)

# Help target to display available commands
help:
	@echo "Makefile commands:"
	@echo "  make         Build the application"
	@echo "  make clean   Remove the built executable"
	@echo "  make run     Build and run the application"
	@echo "  make help    Show this help message"

