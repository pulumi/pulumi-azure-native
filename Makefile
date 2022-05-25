PROJECT_NAME := Pulumi Native Azure Resource Provider

PACK            := azure-native
PACKDIR         := sdk
PROJECT         := github.com/pulumi/pulumi-azure-native
PROVIDER        := pulumi-resource-${PACK}
CODEGEN         := pulumi-gen-${PACK}
VERSION         := $(shell pulumictl get version)

PROVIDER_PKGS    := $(shell cd ./provider && go list ./pkg/...)
WORKING_DIR     := $(shell pwd)

VERSION_FLAGS   := -ldflags "-X github.com/pulumi/pulumi-azure-native/provider/pkg/version.Version=${VERSION}"

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
		(cd $$submodule && git checkout main && git pull origin main); \
	done
	rm ./azure-provider-versions/provider_list.json
	az provider list >> ./azure-provider-versions/provider_list.json

ensure:: init_submodules
	@echo "GO111MODULE=on go mod download"; cd provider; GO111MODULE=on go mod download

local_generate_code:: clean
	$(WORKING_DIR)/bin/$(CODEGEN) schema,nodejs,dotnet,python,go ${VERSION}
	cd ${PACKDIR}/dotnet/ && \
	sed -i.bak -e "s/<\/Nullable>/<\/Nullable>\n    <UseSharedCompilation>false<\/UseSharedCompilation>/g" Pulumi.AzureNative.csproj && \
	rm Pulumi.AzureNative.csproj.bak && \
	cd ../nodejs/ && \
	sed -i.bak -e "s/sourceMap/inlineSourceMap/g" tsconfig.json && \
	rm tsconfig.json.bak
	echo "Finished generating."

local_generate:: clean
	$(WORKING_DIR)/bin/$(CODEGEN) schema,docs,nodejs,dotnet,python,go ${VERSION}
	cd ${PACKDIR}/dotnet/ && \
	sed -i.bak -e "s/<\/Nullable>/<\/Nullable>\n    <UseSharedCompilation>false<\/UseSharedCompilation>/g" Pulumi.AzureNative.csproj && \
	rm Pulumi.AzureNative.csproj.bak && \
	cd ../nodejs/ && \
	sed -i.bak -e "s/sourceMap/inlineSourceMap/g" tsconfig.json && \
	rm tsconfig.json.bak
	echo "Finished generating."

generate_schema::
	echo "Generating Pulumi schema..."
	$(WORKING_DIR)/bin/$(CODEGEN) schema ${VERSION}
	echo "Finished generating schema."

generate_docs::
	$(WORKING_DIR)/bin/$(CODEGEN) docs ${VERSION}

arm2pulumi::
	(cd provider && go build -o $(WORKING_DIR)/bin/arm2pulumi $(VERSION_FLAGS) $(PROJECT)/provider/cmd/arm2pulumi)

arm2pulumi_coverage_report::
	(cd provider/pkg/arm2pulumi/internal/testdata && if [ ! -d azure-quickstart-templates ]; then git clone https://github.com/Azure/azure-quickstart-templates && cd azure-quickstart-templates && git checkout 3b2757465c2de537e333f5e2d1c3776c349b8483; fi)
	(cd provider && go test -v -tags=coverage -run TestQuickstartTemplateCoverage github.com/pulumi/pulumi-azure-native/provider/pkg/arm2pulumi/internal/test)

codegen::
	(cd provider && go build -o $(WORKING_DIR)/bin/$(CODEGEN) $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(CODEGEN))

provider::
	(cd provider && go build -o $(WORKING_DIR)/bin/$(PROVIDER) $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(PROVIDER))

install_provider::
	(cd provider && go install $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(PROVIDER))

test_provider::
	(cd provider && go test -v $(PROVIDER_PKGS))

lint_provider:: provider # lint the provider code
	cd provider && GOGC=20 golangci-lint run -c ../.golangci.yml

define FAKE_MODULE
module fake_module // Exclude this directory from Go tools

go 1.16
endef

export FAKE_MODULE

$(WORKING_DIR)/sdk/nodejs/go.mod:
	echo "$$FAKE_MODULE" | sed 's/fake_module/fake_nodejs_module/g' > $@

generate_nodejs:: $(WORKING_DIR)/sdk/nodejs/go.mod
	$(WORKING_DIR)/bin/$(CODEGEN) nodejs ${VERSION} && \
	cd ${PACKDIR}/nodejs/ && \
	sed -i.bak -e "s/sourceMap/inlineSourceMap/g" tsconfig.json && \
	rm tsconfig.json.bak

build_nodejs:: VERSION := $(shell pulumictl get version --language javascript)
build_nodejs::
	cd ${PACKDIR}/nodejs/ && \
	yarn install && \
	NODE_OPTIONS=--max-old-space-size=8192 yarn run tsc --diagnostics && \
	cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/ && \
	sed -i.bak -e "s/\$${VERSION}/$(VERSION)/g" ./bin/package.json

$(WORKING_DIR)/sdk/python/go.mod:
	echo "$$FAKE_MODULE" | sed 's/fake_module/fake_python_module/g' > $@

generate_python:: $(WORKING_DIR)/sdk/python/go.mod
	$(WORKING_DIR)/bin/$(CODEGEN) python ${VERSION}

build_python:: PYPI_VERSION := $(shell pulumictl get version --language python)
build_python::
	cd sdk/python/ && \
	cp ../../README.md . && \
	python3 setup.py clean --all 2>/dev/null && \
	rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
	sed -i.bak -e 's/^VERSION = .*/VERSION = "$(PYPI_VERSION)"/g' -e 's/^PLUGIN_VERSION = .*/PLUGIN_VERSION = "$(VERSION)"/g' ./bin/setup.py && \
	rm ./bin/setup.py.bak && \
	rm ./bin/go.mod && \
	cd ./bin && python3 setup.py build sdist

$(WORKING_DIR)/sdk/dotnet/go.mod:
	echo "$$FAKE_MODULE" | sed 's/fake_module/fake_dotnet_module/g' > $@

generate_dotnet:: $(WORKING_DIR)/sdk/dotnet/go.mod
	$(WORKING_DIR)/bin/$(CODEGEN) dotnet ${VERSION} && \
	cd ${PACKDIR}/dotnet/ && \
	sed -i.bak -e "s/<\/Nullable>/<\/Nullable>\n    <UseSharedCompilation>false<\/UseSharedCompilation>/g" Pulumi.AzureNative.csproj && \
	rm Pulumi.AzureNative.csproj.bak

build_dotnet:: DOTNET_VERSION := $(shell pulumictl get version --language dotnet)
build_dotnet::
	cd ${PACKDIR}/dotnet/ && \
	echo "azure-native\n${DOTNET_VERSION}" >version.txt && \
	dotnet build /p:Version=${DOTNET_VERSION}

generate_go::
	$(WORKING_DIR)/bin/$(CODEGEN) go ${VERSION}

build_go::
	# Only building the top level packages and building 1 package at a time to avoid OOMing
	cd sdk/ && \
	GOGC=50 go list github.com/pulumi/pulumi-azure-native/sdk/go/azure/... | grep -v "latest\|\/v.*"$ | xargs -L 1 go build

clean::
	rm -rf $$(find sdk/nodejs -mindepth 1 -maxdepth 1 ! -name "go.mod")
	rm -rf $$(find sdk/python -mindepth 1 -maxdepth 1 ! -name "go.mod" ! -name "README.md")
	rm -rf $$(find sdk/dotnet -mindepth 1 -maxdepth 1 ! -name "go.mod")
	rm -rf sdk/go/azure
	rm -rf sdk/schema

install_dotnet_sdk::
	mkdir -p $(WORKING_DIR)/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;

install_python_sdk::

install_go_sdk::

install_nodejs_sdk::
	yarn link --cwd $(WORKING_DIR)/sdk/nodejs/bin

test:: export PULUMI_LOCAL_NUGET=${WORKING_DIR}/nuget
test::
	cd examples && go test -v -tags=all -timeout 2h

build:: init_submodules clean codegen local_generate provider build_sdks install_sdks
build_sdks: build_nodejs build_dotnet build_python build_go
install_sdks:: install_dotnet_sdk install_python_sdk install_nodejs_sdk

# Required for the codegen action that runs in pulumi/pulumi
only_build:: build

.PHONY: init_submodules update_submodules ensure generate_schema generate build_provider build arm2pulumi_coverage_report
