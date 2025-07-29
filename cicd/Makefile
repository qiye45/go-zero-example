#OS := $(shell uname -s 2>/dev/null)
OS := Windows

ifeq ($(OS),Linux)
	TIME_CMD = time
else
	TIME_CMD = powershell -Command "Measure-Command {"
	TIME_END = "}"
endif

.PHONY: lint
lint:
	@echo "==> Running golangci-lint..."
ifeq ($(OS),Linux)
	@cd zero-admin && $(TIME_CMD) golangci-lint run -v --timeout=5m ./...
else
	@cd zero-admin && $(TIME_CMD) golangci-lint run -v --timeout=5m ./... $(TIME_END)
endif

.PHONY: sec
sec:
	@echo "==> Running gosec..."
ifeq ($(OS),Linux)
	@cd zero-admin && $(TIME_CMD) gosec ./...
else
	@cd zero-admin && $(TIME_CMD) gosec ./... $(TIME_END)
endif
