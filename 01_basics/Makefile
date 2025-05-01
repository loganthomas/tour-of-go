.DEFAULT_GOAL := build

.PHONY: clean fmt vet build

# Set default directory to `00_hello` if not specified
# DIR ?= 00_hello

# Check if the specified directory exists and if it is a valid Go module
check-dir:
	@if [ ! -d "$(DIR)" ]; then \
        echo "Directory $(DIR) does not exist."; \
        exit 1; \
	elif [ ! -f "$(DIR)/go.mod" ]; then \
        echo "No Go module found in $(DIR). Please run 'go mod init' in this directory."; \
        exit 1; \
    fi

# Clean command
clean: check-dir
	@echo "Cleaning project in $(DIR)..."
	@cd $(DIR) && go clean

# Format command
fmt: clean
	@echo "Formatting Go files in $(DIR)..."
	@cd $(DIR) && go fmt ./...

# Vet command
vet: fmt
	@echo "Vetting Go files in $(DIR)..."
	@cd $(DIR) && go vet ./...

# Build command
build: vet
	@echo "Building the project in $(DIR)..."
	@cd $(DIR) && go build -o ./bin/ ./...
	@echo "Build completed: ./$(DIR)/bin/"
