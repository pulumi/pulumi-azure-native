PROJECT_NAME := Pulumi AzureRM Resource Provider

PACK            := azurerm
PACKDIR         := sdk
PROJECT         := github.com/pulumi/pulumi-azurerm
PROVIDER        := pulumi-resource-${PACK}
CODEGEN         := pulumi-gen-${PACK}
VERSION         := 0.1.0

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

generate_schema::
	$(GO) install $(VERSION_FLAGS) $(PROJECT)/cmd/$(CODEGEN)
	echo "Generating Pulumi schema..."
	$(CODEGEN) schema
	echo "Finished generating schema."

generate::
	rm -rf sdk/nodejs
	$(GO) install $(VERSION_FLAGS) $(PROJECT)/cmd/$(CODEGEN)
	echo "Generating Pulumi Schema & SDK..."
	$(CODEGEN) schema,nodejs
	echo "Finished generating Schema & SDK."
	cd ${PACKDIR}/nodejs/ && \
		sed -i.bak "s/\$${VERSION}/$(VERSION)/g" ./package.json && \
		yarn link @pulumi/pulumi
