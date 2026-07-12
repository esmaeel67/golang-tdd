test:
	go test -v ./...
help:
	@echo "Available targets:"
	@echo "  build  - Build the application"
	@echo "  test   - Run tests"
	@echo "  run    - Run the application"
	@echo "  clean  - Remove artifacts"
	@echo "  help   - Show this help"