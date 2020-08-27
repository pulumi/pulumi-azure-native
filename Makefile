PROJECT_NAME := Pulumi AzureRM Resource Provider

PACK            := azurerm
PACKDIR         := sdk
PROJECT         := github.com/pulumi/pulumi-azurerm
PROVIDER        := pulumi-resource-${PACK}
CODEGEN         := pulumi-gen-${PACK}
VERSION         := 0.1.0

WORKING_DIR     := $(shell pwd)

VERSION_FLAGS   := -ldflags "-X github.com/pulumi/pulumi-azurerm/provider/pkg/version.Version=${VERSION}"

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

local_generate::
	echo "Generating Pulumi schema..."
	$(WORKING_DIR)/bin/$(CODEGEN) schema,nodejs,go,dotnet,python
	echo "Finished generating schema."

generate_schema::
	echo "Generating Pulumi schema..."
	$(WORKING_DIR)/bin/$(CODEGEN) schema
	echo "Finished generating schema."

codegen::
	(cd provider && go build -a -o $(WORKING_DIR)/bin/$(CODEGEN) $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(CODEGEN))

provider::
	(cd provider && go build -a -o $(WORKING_DIR)/bin/$(PROVIDER) $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(PROVIDER))

generate_nodejs::
	$(WORKING_DIR)/bin/$(CODEGEN) nodejs

build_nodejs::
	cd ${PACKDIR}/nodejs/ && \
		yarn install && \
		yarn run tsc && \
		cp ../../README.md package.json yarn.lock ./bin/ && \
		sed -i.bak -e "s/\$${VERSION}/$(VERSION)/g" ./bin/package.json

generate_python::
	$(WORKING_DIR)/bin/$(CODEGEN) python

build_python::
	#TODO: remove -e "40d" below when plugin installation works again
	cd ${PACKDIR}/python/ && \
		cp ../../README.md . && \
		sed -i.bak -e "s/\$${VERSION}/$(VERSION)/g" -e "s/\$${PLUGIN_VERSION}/$(VERSION)/g" -e "40d" ./setup.py && \
		rm ./setup.py.bak

generate_dotnet::
	$(WORKING_DIR)/bin/$(CODEGEN) dotnet

build_dotnet::
	cd ${PACKDIR}/dotnet/ && \
		echo "${VERSION:v%=%}" >version.txt && \
		dotnet build

generate_go::
	$(WORKING_DIR)/bin/$(CODEGEN) go

build_go::

clean::
	rm -rf sdk/nodejs
	rm -rf sdk/python
	rm -rf sdk/dotnet
	rm -rf sdk/go/azurerm
	rm -rf sdk/schema

install_dotnet_sdk::
	mkdir -p $(WORKING_DIR)/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;

install_python_sdk::

install_go_sdk::

install_nodejs_sdk::
	yarn link --cwd $(WORKING_DIR)/sdk/nodejs/bin

test::
	cd examples && go test -v -tags=all -timeout 2h

build:: init_submodules clean codegen local_generate provider build_sdks install_sdks
build_sdks: build_nodejs build_dotnet build_python build_go
install_sdks:: install_dotnet_sdk install_python_sdk install_nodejs_sdk

.PHONY: init_submodules update_submodules ensure generate_schema generate build_provider build
