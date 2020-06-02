PROJECT_NAME := Pulumi AzureRM Resource Provider

PACK            := azurerm
PROJECT         := github.com/pulumi/pulumi-azurerm
PROVIDER        := pulumi-resource-${PACK}
VERSION         := 0.1

VERSION_FLAGS   := -ldflags "-X github.com/pulumi/pulumi-azurerm/provider/v2/pkg/version.Version=${VERSION}"

GO              ?= go
CURL            ?= curl

ensure::
	@echo "GO111MODULE=on go mod tidy"; GO111MODULE=on go mod tidy
	@echo "GO111MODULE=on go mod download"; GO111MODULE=on go mod download

build::
	$(GO) install $(VERSION_FLAGS) $(PROJECT)/cmd/$(PROVIDER)

builddebug:
	$(GO) install $(VERSION_FLAGS) -gcflags="all=-N -l" $(PROJECT)/cmd/$(PROVIDER)
