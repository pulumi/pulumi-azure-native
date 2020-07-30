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

init_submodules::
	@for submodule in $$(git submodule status | awk {'print $$2'}); do \
		if [ ! -f "$$submodule/.git" ]; then \
			echo "Initializing submodule $$submodule" ; \
			(cd $$submodule && git submodule update --init); \
		fi; \
	done

update_submodules:: init_submodules
	@for submodule in $$(git submodule status | awk {'print $$2'}); do \
		echo "Updating submodule $$submodule" ; \
		(cd $$submodule && git pull origin master); \
	done

ensure:: init_submodules
	@echo "GO111MODULE=on go mod tidy"; cd provider; GO111MODULE=on go mod tidy
	@echo "GO111MODULE=on go mod download"; cd provider; GO111MODULE=on go mod download

generate_schema::
	cd provider; $(GO) install $(VERSION_FLAGS) $(PROJECT)/cmd/$(CODEGEN)
	echo "Generating Pulumi schema..."
	$(CODEGEN) schema $(VERSION)
	echo "Finished generating schema."

generate::
	rm -rf sdk/nodejs
	rm -rf sdk/python
	rm -rf sdk/dotnet
	rm -rf sdk/go/azurerm
	cd provider; $(GO) install $(VERSION_FLAGS) $(PROJECT)/cmd/$(CODEGEN)
	echo "Generating Pulumi Schema & SDK..."
	$(CODEGEN) schema,nodejs,python,go,dotnet
	echo "Finished generating Schema & SDK."
	cd ${PACKDIR}/nodejs/ && \
		sed -i.bak "s/\$${VERSION}/$(VERSION)/g" ./package.json && \
		yarn install
	cd ${PACKDIR}/python/ && \
			cp ../../README.md . && \
			sed -i.bak -e "s/\$${VERSION}/$(PYPI_VERSION)/g" -e "s/\$${PLUGIN_VERSION}/$(VERSION)/g" ./setup.py && \
			rm ./setup.py.bak
	cd ${PACKDIR}/dotnet/ && \
		echo "${VERSION:v%=%}" >version.txt && \
		dotnet build

build_provider::
	cd provider; $(GO) install $(VERSION_FLAGS) $(PROJECT)/cmd/$(PROVIDER)

build:: generate build_provider

.PHONY: init_submodules update_submodules ensure generate_schema generate build_provider build
