BINARY_NAME := os-check
LOCAL_BIN_PATH := ${HOME}/.local/bin
.PHONY: build
build:
	@echo "Building..."
	@go build -o bin/$(BINARY_NAME) -v

.PHONY: run
run: build
	@echo "Running..."
	@bin/$(BINARY_NAME)

.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -rf bin/

.PHONY: mv
mv: build
	@echo "Moving..."
	@mv bin/$(BINARY_NAME) $(LOCAL_BIN_PATH)/$(BINARY_NAME)