VERSION ?= 0.0.1
CONTAINER_MANAGER ?= podman

# Image URL to use all building/pushing image targets
IMG ?= quay.io/ariobolo/seminario-pffp:v${VERSION}

# Go and compilation related variables
GOPATH ?= $(shell go env GOPATH)
BUILD_DIR ?= out
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

# Add default target
.PHONY: default
default: install

# Create and update the vendor directory
.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

.PHONY: check
check: build test lint

# Start of the actual build targets
$(BUILD_DIR)/seminario: $(SOURCES)
	GOOS="$(GOOS)" GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/seminario .
 
.PHONY: build
build: $(BUILD_DIR)/seminario

.PHONY: e2e
e2e:
	curl localhost:8080/seminario

.PHONY: test
test:
	go test .


.PHONY: clean ## Remove all build artifacts
clean:
	rm -rf $(BUILD_DIR)

.PHONY: fmt
fmt:
	@gofmt -l -w $(SOURCE_DIRS)

$(GOPATH)/bin/golangci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.0

# Run golangci-lint against code
.PHONY: lint
lint: $(GOPATH)/bin/golangci-lint
	$(GOPATH)/bin/golangci-lint run

# Build the container image
.PHONY: oci-build
oci-build: clean oci-build-amd64 oci-build-arm64

# Build for amd64 architecture only
.PHONY: oci-build-amd64
oci-build-amd64: clean
	# Build the container image for amd64
	${CONTAINER_MANAGER} build --platform linux/amd64 --manifest $(IMG)-amd64 -f oci/Containerfile .

# Build for arm64 architecture only
.PHONY: oci-build-arm64
oci-build-arm64: clean
	# Build the container image for arm64
	${CONTAINER_MANAGER} build --platform linux/arm64 --manifest $(IMG)-arm64 -f oci/Containerfile .

# Save images for amd64 architecture only
.PHONY: oci-save-amd64
oci-save-amd64:
	${CONTAINER_MANAGER} save -m -o $(PFFP_SAVE)-amd64.tar $(IMG)-amd64

# Save images for arm64 architecture only
.PHONY: oci-save-arm64
oci-save-arm64:
	${CONTAINER_MANAGER} save -m -o $(PFFP_SAVE)-arm64.tar $(IMG)-arm64


PFFP_SAVE ?= pffp
.PHONY: oci-save 
oci-save: oci-save-amd64 oci-save-arm64

oci-load:
	${CONTAINER_MANAGER} load -i $(PFFP_SAVE)-arm64/$(PFFP_SAVE)-arm64.tar 
	${CONTAINER_MANAGER} load -i $(PFFP_SAVE)-amd64/$(PFFP_SAVE)-amd64.tar 

# Push the docker image
.PHONY: oci-push
oci-push:
	${CONTAINER_MANAGER} push $(IMG)-arm64
	${CONTAINER_MANAGER} push $(IMG)-amd64
	${CONTAINER_MANAGER} manifest create $(IMG)
	${CONTAINER_MANAGER} manifest add $(IMG) docker://$(IMG)-arm64
	${CONTAINER_MANAGER} manifest add $(IMG) docker://$(IMG)-amd64
	${CONTAINER_MANAGER} manifest push --all $(IMG)

# Build for amd64 architecture only
.PHONY: oci-run-amd64
oci-run-amd64: 
	${CONTAINER_MANAGER} run -d --rm --name seminario \
		-p 8081:8080 \
		--name seminario $(IMG)-amd64 
	
	${CONTAINER_MANAGER} logs -f seminario 