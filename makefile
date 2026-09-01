BOOK_SWAP_PORT ?= 9090  # Sets default port to 9090 if not set elsewhere

test:
	LONG=true go test -v ./handlers/...
help:
	@echo "Available targets:"
	@echo "  build  - Build the application"
	@echo "  test   - Run tests"
	@echo "  run    - Run the application"
	@echo "  clean  - Remove artifacts"
	@echo "  help   - Show this help"
serve:
	BOOK_SWAP_PORT=$(BOOK_SWAP_PORT) go run ./cmd/main.go