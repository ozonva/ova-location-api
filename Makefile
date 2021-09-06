GO_VERSION_SHORT:=$(shell echo `go version` | sed -E 's/.* go(.*) .*/\1/g')
ifneq ("1.16","$(shell printf "$(GO_VERSION_SHORT)\n1.16" | sort -V | head -1)")
$(error NEED GO VERSION >= 1.16. Found: $(GO_VERSION_SHORT))
endif

export GO111MODULE=on
export GOPROXY=https://proxy.golang.org|direct

PGV_VERSION:="v0.6.1"
GOOGLEAPIS_VERSION="master"
BUF_VERSION:="v0.51.0"

all: generate build

.PHONY: vendor-proto
vendor-proto:
	$(eval VENDOR_PROTOGEN:=$(CURDIR)/vendor.protogen)
	@[ -f $(VENDOR_PROTOGEN)/validate/validate.proto ] || (mkdir -p $(VENDOR_PROTOGEN)/validate/ && curl -sSL0 https://raw.githubusercontent.com/envoyproxy/protoc-gen-validate/$(PGV_VERSION)/validate/validate.proto -o $(VENDOR_PROTOGEN)/validate/validate.proto)
	@[ -f $(VENDOR_PROTOGEN)/google/api/http.proto ] || (mkdir -p $(VENDOR_PROTOGEN)/google/api/ && curl -sSL0 https://raw.githubusercontent.com/googleapis/googleapis/$(GOOGLEAPIS_VERSION)/google/api/http.proto -o $(VENDOR_PROTOGEN)/google/api/http.proto)
	@[ -f $(VENDOR_PROTOGEN)/google/api/annotations.proto ] || (mkdir -p $(VENDOR_PROTOGEN)/google/api/ && curl -sSL0 https://raw.githubusercontent.com/googleapis/googleapis/$(GOOGLEAPIS_VERSION)/google/api/annotations.proto -o $(VENDOR_PROTOGEN)/google/api/annotations.proto)

GOBIN?=$(GOPATH)/bin
buf.work.yaml:
	@echo "version: v1\ndirectories:\n  - api\n  - vendor.protogen\n" > $(CURDIR)/buf.work.yaml
buf.gen.yaml:
	@echo "version: v1\nplugins:\n  - name: go\n    out: .\n    opt: module=github.com/ozonva/ova-location-api\n  - name: go\n    out: .\n    opt: module=github.com/ozonva/ova-location-api\n  - name: go-grpc\n    out: .\n    opt: module=github.com/ozonva/ova-location-api\n  - name: grpc-gateway\n    out: .\n    opt: logtostderr=true,module=github.com/ozonva/ova-location-api\n  - name: openapiv2\n    out: swagger\n    opt: allow_merge=true,merge_file_name=api" > $(CURDIR)/buf.gen.yaml

.PHONY: generate
generate: vendor-proto buf.work.yaml buf.gen.yaml
	@command -v buf 2>&1 > /dev/null || (mkdir -p $(GOBIN) && curl -sSL0 https://github.com/bufbuild/buf/releases/download/$(BUF_VERSION)/buf-$(shell uname -s)-$(shell uname -m) -o $(GOBIN)/buf && chmod +x $(GOBIN)/buf)
	$(eval PROTOS:=$(CURDIR)/api)
	@[ -f $(PROTOS)/buf.yaml ] || PATH=$(GOBIN):$(PATH) buf mod init --doc -o $(PROTOS)
	PATH=$(GOBIN):$(PATH) buf generate $(PROTOS)

.PHONY: deps
deps:
	@[ -f go.mod ] || go mod init github.com/ozonva/ova-location-api
	find . -name go.mod -exec bash -c 'pushd "$${1%go.mod}" && go mod tidy && popd' _ {} \;

.PHONY: build
build: deps
	go build -o $(CURDIR)/bin/project $(CURDIR)/main.go

bin-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.5.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.5.0
	go install github.com/envoyproxy/protoc-gen-validate@$(PGV_VERSION)