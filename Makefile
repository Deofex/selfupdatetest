# Define the VERSION variable
VERSION ?=

# Define the output directory
OUTPUT_DIR := docs

# Replace dots in the version with dashes for the output file
VERSION_DASHED := $(shell echo $(VERSION) | tr '.' '-')

# Define the target name format
OUTPUT_FILE_PREFIX := $(OUTPUT_DIR)/selfupdatetest-$(VERSION_DASHED)

# Main target
build: check-version
	@mkdir -p $(OUTPUT_DIR)
	@echo "Building version $(VERSION)..."
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w -X main.version=$(VERSION)" -o $(OUTPUT_FILE_PREFIX)-darwin-amd64
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w -X main.version=$(VERSION)" -o $(OUTPUT_FILE_PREFIX)-darwin-arm64
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -X main.version=$(VERSION)" -o $(OUTPUT_FILE_PREFIX)-windows-amd64.exe
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X main.version=$(VERSION)" -o $(OUTPUT_FILE_PREFIX)-linux-amd64

# Check if VERSION is set
check-version:
ifndef VERSION
	$(error VERSION is not set. Please set it using 'make build VERSION=<version>')
endif

