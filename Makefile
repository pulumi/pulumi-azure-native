PROJECT_NAME := Pulumi AzureRM Resource Provider

PACK            := azurerm
PACKDIR         := sdk
PROJECT         := github.com/pulumi/pulumi-azurerm
PROVIDER        := pulumi-resource-${PACK}
CODEGEN         := pulumi-gen-${PACK}
VERSION         := 0.1.0

PROVIDER_PKGS    := $(shell cd ./provider && go list ./...)
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

# Need to generate provider for example generation since program-gen causes the provider to be loaded
local_generate:: codegen generate_schema provider generate_sdk_with_examples
	echo "Finished generating schema."

generate_schema:
	echo "Generating Pulumi schema..."
	$(WORKING_DIR)/bin/$(CODEGEN) -version=$(VERSION) -schema=true
	echo "Finished generating schema."

generate_sdk_with_examples::
	echo "Generating Pulumi schema with examples..."
	# not generating golang examples currently due to https://github.com/pulumi/pulumi-azurerm/issues/156
	$(WORKING_DIR)/bin/$(CODEGEN) -version=$(VERSION) -sdk=true "go"
	$(WORKING_DIR)/bin/$(CODEGEN) -version=$(VERSION) -examples -sdk -debug "nodejs,dotnet,python"
	echo "Finished generating schema with examples."

codegen::
	(cd provider && go build -a -o $(WORKING_DIR)/bin/$(CODEGEN) $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(CODEGEN))

provider::
	(cd provider && go build -a -o $(WORKING_DIR)/bin/$(PROVIDER) $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(PROVIDER) && cp $(WORKING_DIR)/bin/$(PROVIDER) $(GOPATH)/bin/)

test_provider::
	(cd provider && go test -v $(PROVIDER_PKGS))

lint_provider:: provider # lint the provider code
	cd provider && GOGC=20 golangci-lint run -c ../.golangci.yml

generate_nodejs::
	$(WORKING_DIR)/bin/$(CODEGEN) -version=$(VERSION) -examples=true -sdk=true nodejs

build_nodejs::
	cd ${PACKDIR}/nodejs/ && \
		yarn install && \
		tsc && \
		cp ../../README.md package.json yarn.lock ./bin/ && \
		sed -i.bak -e "s/\$${VERSION}/$(VERSION)/g" ./bin/package.json

generate_python::
	$(WORKING_DIR)/bin/$(CODEGEN) -version=$(VERSION) -examples=true -sdk=true python

build_python::
	cd sdk/python/ && \
        cp ../../README.md . && \
        python3 setup.py clean --all 2>/dev/null && \
        rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
        sed -i.bak -e "s/\$${VERSION}/$(VERSION)/g" -e "s/\$${PLUGIN_VERSION}/$(VERSION)/g" ./bin/setup.py && \
        rm ./bin/setup.py.bak && \
        cd ./bin && python3 setup.py build sdist

generate_dotnet::
	$(WORKING_DIR)/bin/$(CODEGEN) -version=$(VERSION) -examples=true -sdk=true dotnet

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

.PHONY: init_submodules update_submodules ensure generate_schema local_generate generate_sdk_with_examples provider build
