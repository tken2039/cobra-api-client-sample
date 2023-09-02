BIN := bin/cobrapi
MAIN := main.go

.PHONY: build
build:
	@echo "Building..."
	@go build -o $(BIN) $(MAIN)
