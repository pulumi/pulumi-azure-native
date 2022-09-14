PROJECT_NAME := Pulumi Native Azure Resource Provider

PACK            := azure-native
PROJECT         := github.com/pulumi/pulumi-azure-native
PROVIDER        := pulumi-resource-${PACK}
CODEGEN         := pulumi-gen-${PACK}
VERSIONER       := pulumi-versioner-${PACK}

PROVIDER_PKGS   := $(shell cd ./provider && go list ./pkg/...)
WORKING_DIR     := $(shell pwd)

PROVIDER_PKG	= $(shell find provider/pkg -type f)
SPECS           = $(shell find azure-rest-api-specs/specification/*/resource-manager -type f -name "*.json" ! -path "**/examples/**")

JAVA_GEN 		 := pulumi-java-gen
JAVA_GEN_VERSION := v0.5.4

# These are lazy variables which are only computed when requested
# When using any version variable, depend on bin/pulumictl
VERSION         = $(shell bin/pulumictl get version)
VERSION_JS      = $(shell bin/pulumictl get version --language javascript)
VERSION_DOTNET  = $(shell bin/pulumictl get version --language dotnet)
VERSION_PYTHON  = $(shell bin/pulumictl get version --language python)
VERSION_FLAGS   = -ldflags "-X github.com/pulumi/pulumi-azure-native/provider/pkg/version.Version=${VERSION}"

default: init_submodules provider arm2pulumi local_generate
build: init_submodules codegen local_generate provider build_sdks
build_sdks: build_nodejs build_dotnet build_python build_go build_java
install_sdks: install_dotnet_sdk install_python_sdk install_nodejs_sdk
arm2pulumi: bin/arm2pulumi
codegen: bin/$(CODEGEN)
provider: bin/$(PROVIDER)
versioner: bin/pulumi-versioner-azure-native
versions: versions/spec.json versions/v1.json versions/v2.json versions/deprecated.json versions/pending.json versions/active.json

generate_schema: provider/cmd/$(PROVIDER)/schema-full.json
generate_docs: provider/cmd/$(PROVIDER)/schema.json

generate_java: sdk/java/gen.sentinel
generate_nodejs: sdk/nodejs/gen.sentinel
generate_python: sdk/python/gen.sentinel
generate_dotnet: sdk/dotnet/gen.sentinel
generate_go: sdk/go/gen.sentinel sdk/pulumi-azure-native-sdk/local.sentinel

local_generate_code: generate_java
local_generate_code: generate_nodejs
local_generate_code: generate_python
local_generate_code: generate_dotnet
local_generate_code: generate_go

local_generate: generate_docs
local_generate: generate_schema
local_generate: local_generate_code

build_nodejs: sdk/nodejs/build.sentinel
build_python: sdk/python/build.sentinel
build_dotnet: sdk/dotnet/build.sentinel
build_java: sdk/java/build.sentinel
build_go: sdk/go/build.sentinel sdk/pulumi-azure-native-sdk/local.sentinel

# Required by CI steps - some can be skipped
install_dotnet_sdk: sdk/dotnet/install.sentinel
install_python_sdk:
install_go_sdk:
install_java_sdk:
install_nodejs_sdk: sdk/nodejs/install.sentinel

install_provider: install_provider.sentinel

prepublish_go: sdk/pulumi-azure-native-sdk/prepublish.sentinel

# Required for the codegen action that runs in pulumi/pulumi
only_build: build

.PHONY: ensure build build_sdks install_sdks arm2pulumi codegen provider versioner versions
.PHONY: generate_schema generate_docs generate_java
.PHONY: init_submodules update_submodules local_generate_code local_generate arm2pulumi_coverage_report
.PHONY: test_provider lint_provider generate_nodejs build_nodejs generate_python build_python generate_dotnet

ensure: init_submodules bin/pulumictl provider/.mod_download.sentinel
	@jq --version > /dev/null

init_submodules:
	@for submodule in $$(git submodule status | awk {'print $$2'}); do \
		if [ ! -f "$$submodule/.git" ]; then \
			echo "Initializing submodule $$submodule" ; \
			(cd $$submodule && git submodule update --init); \
		fi; \
	done

update_submodules: init_submodules
	@for submodule in $$(git submodule status | awk {'print $$2'}); do \
		echo "Updating submodule $$submodule" ; \
		(cd $$submodule && git checkout main && git pull origin main); \
	done
	rm ./azure-provider-versions/provider_list.json
	az provider list | jq 'map({ namespace: .namespace, resourceTypes: .resourceTypes | map({ resourceType: .resourceType, apiVersions: .apiVersions }) | sort_by(.resourceType) }) | sort_by(.namespace)' > ./azure-provider-versions/provider_list.json

arm2pulumi_coverage_report:
	(cd provider/pkg/arm2pulumi/internal/testdata && if [ ! -d azure-quickstart-templates ]; then git clone https://github.com/Azure/azure-quickstart-templates && cd azure-quickstart-templates && git checkout 3b2757465c2de537e333f5e2d1c3776c349b8483; fi)
	(cd provider && go test -v -tags=coverage -run TestQuickstartTemplateCoverage github.com/pulumi/pulumi-azure-native/provider/pkg/arm2pulumi/internal/test)

test_provider:
	(cd provider && go test -v $(PROVIDER_PKGS))

lint_provider: provider # lint the provider code
	cd provider && GOGC=20 golangci-lint run -c ../.golangci.yml

clean:
	find bin -maxdepth 1 -type f -delete
	rm -rf nuget
	find . -maxdepth 2 -name "*.sentinel" -delete
	cd provider/cmd/arm2pulumi && rm -f metadata-compact.json schema-full.json
	cd provider/cmd/pulumi-resource-azure-native && rm -f metadata-compact.json schema-full.json
	rm -rf sdk/dotnet/bin
	rm -rf sdk/dotnet/build sdk/dotnet/src sdk/dotnet/.gradle
	rm -rf sdk/nodejs/bin
	rm -rf sdk/pulumi-azure-native-sdk
	rm -rf sdk/python/bin
	if dotnet nuget list source | grep "$(WORKING_DIR)"; then \
		dotnet nuget remove source "$(WORKING_DIR)" \
	; fi

test: export PULUMI_LOCAL_NUGET=$(WORKING_DIR)/nuget
test:
	cd examples && go test -v -tags=all -timeout 2h

# --------- File-based targets --------- #

# Download local copy of pulumictl based on the version in .pulumictl.version
# Anywhere which uses VERSION or VERSION_FLAGS should depend on bin/pulumictl
bin/pulumictl: PULUMICTL_VERSION := $(shell cat .pulumictl.version)
bin/pulumictl: PLAT := $(shell go version | sed -En "s/go version go.* (.*)\/(.*)/\1-\2/p")
bin/pulumictl: PULUMICTL_URL := "https://github.com/pulumi/pulumictl/releases/download/v$(PULUMICTL_VERSION)/pulumictl-v$(PULUMICTL_VERSION)-$(PLAT).tar.gz"
bin/pulumictl: .pulumictl.version
	@echo "Installing pulumictl"
	@mkdir -p bin
	wget -q -O - "$(PULUMICTL_URL)" | tar -xzf - -C $(WORKING_DIR)/bin pulumictl
	@touch bin/pulumictl
	@echo "pulumictl" $$(./bin/pulumictl version)

bin/pulumi-java-gen: .pulumi-java-gen.version
	@mkdir -p bin
	pulumictl download-binary -n pulumi-language-java -v $(shell cat .pulumi-java-gen.version) -r pulumi/pulumi-java

provider/.mod_download.sentinel: provider/go.mod provider/go.sum
	cd provider && GO111MODULE=on go mod download
	@touch provider/.mod_download.sentinel

bin/arm2pulumi: bin/pulumictl provider/.mod_download.sentinel provider/cmd/arm2pulumi/* $(PROVIDER_PKG)
	cd provider && go build -o $(WORKING_DIR)/bin/arm2pulumi $(VERSION_FLAGS) $(PROJECT)/provider/cmd/arm2pulumi

bin/$(CODEGEN): bin/pulumictl provider/.mod_download.sentinel provider/cmd/$(CODEGEN)/* $(PROVIDER_PKG)
	cd provider && go build -o $(WORKING_DIR)/bin/$(CODEGEN) $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(CODEGEN)

bin/$(PROVIDER): bin/pulumictl provider/.mod_download.sentinel provider/cmd/$(PROVIDER)/*.go provider/cmd/$(PROVIDER)/schema-full.json $(PROVIDER_PKG)
	cd provider && \
	go build -o $(WORKING_DIR)/bin/$(PROVIDER) $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(PROVIDER)

bin/$(VERSIONER): bin/pulumictl provider/.mod_download.sentinel provider/cmd/$(VERSIONER)/* $(PROVIDER_PKG)
	cd provider && go build -o $(WORKING_DIR)/bin/pulumi-versioner-azure-native $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(VERSIONER)

provider/cmd/$(PROVIDER)/schema.json: bin/$(CODEGEN) $(SPECS)
	bin/$(CODEGEN) docs $(VERSION)

provider/cmd/$(PROVIDER)/schema-full.json: bin/$(CODEGEN) $(SPECS)
	bin/$(CODEGEN) schema $(VERSION)

versions/spec.json: bin/pulumi-versioner-azure-native .git/modules/azure-rest-api-specs/HEAD
	bin/pulumi-versioner-azure-native spec

versions/spec-resources.json: bin/pulumi-versioner-azure-native .git/modules/azure-rest-api-specs/HEAD
	bin/pulumi-versioner-azure-native spec

versions/active.json: bin/pulumi-versioner-azure-native azure-provider-versions/provider_list.json
	bin/pulumi-versioner-azure-native active

versions/v1.json: bin/pulumi-versioner-azure-native versions/spec.json versions/v1-config.yaml
	bin/pulumi-versioner-azure-native v1

versions/deprecated.json: bin/pulumi-versioner-azure-native versions/spec.json versions/v1.json
	bin/pulumi-versioner-azure-native deprecated -version=v1.json

versions/pending.json: bin/pulumi-versioner-azure-native versions/spec.json versions/v1.json
	bin/pulumi-versioner-azure-native pending -version=v1.json

versions/v2.json: bin/pulumi-versioner-azure-native versions/spec.json versions/deprecated.json versions/v2-config.yaml
	bin/pulumi-versioner-azure-native v2

define FAKE_MODULE
module fake_module // Exclude this directory from Go tools

go 1.17
endef

export FAKE_MODULE

sdk/java/gen.sentinel: bin/pulumi-java-gen provider/cmd/$(PROVIDER)/schema.json
	@mkdir -p sdk/java
	rm -rf $$(find sdk/java -mindepth 1 -maxdepth 1)
	bin/$(JAVA_GEN) generate --schema provider/cmd/$(PROVIDER)/schema.json --out sdk/java --build gradle-nexus
	echo "$$FAKE_MODULE" | sed 's/fake_module/fake_java_module/g' > sdk/java/go.mod
	@touch sdk/java/gen.sentinel

sdk/nodejs/gen.sentinel: bin/pulumictl bin/$(CODEGEN) provider/cmd/$(PROVIDER)/schema-full.json
	mkdir -p sdk/nodejs
	rm -rf $$(find sdk/nodejs -mindepth 1 -maxdepth 1 ! -name "go.mod")
	bin/$(CODEGEN) nodejs $(VERSION)
	echo "$$FAKE_MODULE" | sed 's/fake_module/fake_nodejs_module/g' > sdk/nodejs/go.mod
	sed -i.bak -e "s/sourceMap/inlineSourceMap/g" sdk/nodejs/tsconfig.json
	rm sdk/nodejs/tsconfig.json.bak
	@touch sdk/nodejs/gen.sentinel

sdk/python/gen.sentinel: bin/pulumictl bin/$(CODEGEN) provider/cmd/$(PROVIDER)/schema-full.json
	mkdir -p sdk/python
	rm -rf $$(find sdk/python -mindepth 1 -maxdepth 1 ! -name "go.mod")
	bin/$(CODEGEN) python $(VERSION)
	echo "$$FAKE_MODULE" | sed 's/fake_module/fake_python_module/g' > sdk/python/go.mod
	cp README.md sdk/python
	@touch sdk/python/gen.sentinel

sdk/dotnet/gen.sentinel: bin/pulumictl bin/$(CODEGEN) provider/cmd/$(PROVIDER)/schema-full.json
	mkdir -p sdk/dotnet
	rm -rf $$(find sdk/dotnet -mindepth 1 -maxdepth 1 ! -name "go.mod")
	bin/$(CODEGEN) dotnet $(VERSION)
	echo "$$FAKE_MODULE" | sed 's/fake_module/fake_dotnet_module/g' > sdk/dotnet/go.mod
	sed -i.bak -e "s/<\/Nullable>/<\/Nullable>\n    <UseSharedCompilation>false<\/UseSharedCompilation>/g" sdk/dotnet/Pulumi.AzureNative.csproj
	rm sdk/dotnet/Pulumi.AzureNative.csproj.bak
	@touch sdk/dotnet/gen.sentinel

sdk/go/gen.sentinel: bin/pulumictl bin/$(CODEGEN) provider/cmd/$(PROVIDER)/schema-full.json
	mkdir -p sdk/go
	rm -rf sdk/go/azure
	bin/$(CODEGEN) go $(VERSION)
	@# HACK: Strip all comments to make SDK smaller
	find sdk/go -type f -exec sed -i '' -e '/^\/\/.*/g' {} \;
	@touch sdk/go/gen.sentinel

sdk/pulumi-azure-native-sdk/local.sentinel: bin/pulumictl bin/$(CODEGEN) provider/cmd/$(PROVIDER)/schema-full.json
	@mkdir -p sdk/pulumi-azure-native-sdk
	@# Unmark this is as an up-to-date local build
	rm -f sdk/pulumi-azure-native-sdk/publish.sentinel
	rm -rf $$(find sdk/pulumi-azure-native-sdk -mindepth 1 -maxdepth 1 ! -name ".git")
	bin/$(CODEGEN) go-split $(VERSION)
	@# Tidy up all go.mod files
	find sdk/pulumi-azure-native-sdk -type d -maxdepth 1 -exec sh -c "cd \"{}\" && go mod tidy" \;
	@touch sdk/pulumi-azure-native-sdk/local.sentinel

sdk/pulumi-azure-native-sdk/publish.sentinel:
	@# Unmark this is as an up-to-date local build - fail if not build locally first
	rm sdk/pulumi-azure-native-sdk/local.sentinel
	@# Remove go module replacements which are added for local testing
	@# Note: must use `sed -i -e` to be portable - but leaves go.mod-e behind on macos
	find sdk/pulumi-azure-native-sdk -maxdepth 2 -type f -name go.mod -exec sed -i -e '/replace github\.com\/pulumi\/pulumi-azure-native-sdk /d' {} \;
	@# Remove sed backup files if using older sed versions
	find sdk/pulumi-azure-native-sdk -maxdepth 2 -type f -name go.mod-e -delete
	@# Delete go.sum files as these are not used at the point of publishing.
	@# This is because we depend on the root package which will come from the same release commit, that doesn't yet exist.
	find sdk/pulumi-azure-native-sdk -maxdepth 2 -type f -name go.sum -delete
	cp README.md LICENSE sdk/pulumi-azure-native-sdk/
	touch sdk/pulumi-azure-native-sdk/publish.sentinel

# Used by build* targets

sdk/nodejs/node_modules.sentinel: sdk/nodejs/gen.sentinel sdk/nodejs/package.json
	yarn install --cwd sdk/nodejs
	@touch sdk/nodejs/node_modules.sentinel

sdk/nodejs/build.sentinel: bin/pulumictl sdk/nodejs/node_modules.sentinel
	cd sdk/nodejs/ && \
	NODE_OPTIONS=--max-old-space-size=8192 yarn run tsc --diagnostics --incremental && \
	cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/ && \
	sed -i.bak -e "s/\$${VERSION}/$(VERSION_JS)/g" ./bin/package.json
	@touch sdk/nodejs/build.sentinel

sdk/python/build.sentinel: bin/pulumictl sdk/python/gen.sentinel
	cd sdk/python && \
		python3 setup.py clean --all 2>/dev/null && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		sed -i.bak -e 's/^VERSION = .*/VERSION = "$(VERSION_PYTHON)"/g' -e 's/^PLUGIN_VERSION = .*/PLUGIN_VERSION = "$(VERSION)"/g' ./bin/setup.py && \
		rm ./bin/setup.py.bak && \
		rm ./bin/go.mod && \
		cd ./bin && python3 setup.py build sdist
	@touch sdk/python/build.sentinel

sdk/dotnet/build.sentinel: bin/pulumictl sdk/dotnet/gen.sentinel
	cd sdk/dotnet && \
		echo "azure-native\n$(VERSION_DOTNET)" >version.txt && \
		dotnet build /p:Version=$(VERSION_DOTNET)
	@touch sdk/dotnet/build.sentinel

sdk/java/build.sentinel: bin/pulumictl sdk/java/gen.sentinel
	cd sdk/java/ && \
		gradle --console=plain -Pversion=$(VERSION) build
	@touch sdk/java/build.sentinel

sdk/go/build.sentinel: sdk/go/gen.sentinel
	# Only building the top level packages and building 1 package at a time to avoid OOMing
	cd sdk && \
		GOGC=50 go list github.com/pulumi/pulumi-azure-native/sdk/go/azure/... | grep -v "latest\|\/v.*"$ | xargs -L 1 go build
	@touch sdk/go/build.sentinel

sdk/pulumi-azure-native-sdk/build.sentinel: sdk/pulumi-azure-native-sdk/local.sentinel
	find sdk/pulumi-azure-native-sdk -type d -maxdepth 1 -exec sh -c "cd \"{}\" && go build" \;
	@touch sdk/pulumi-azure-native-sdk/build.sentinel

# Used by install* targets

sdk/nodejs/install.sentinel: sdk/nodejs/build.sentinel
	yarn link --cwd sdk/nodejs/bin
	@touch sdk/nodejs/install.sentinel

sdk/dotnet/install.sentinel: sdk/dotnet/build.sentinel
	mkdir -p nuget
	find sdk/dotnet/bin -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;
	@if ! dotnet nuget list source | grep ${WORKING_DIR}; then \
		dotnet nuget add source ${WORKING_DIR}/nuget --name ${WORKING_DIR} \
	; fi
	@touch sdk/dotnet/install.sentinel

install_provider.sentinel: bin/pulumictl provider/.mod_download.sentinel provider/cmd/$(PROVIDER)/* $(PROVIDER_PKG)
	cd provider && go install $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(PROVIDER)
	@touch install_provider.sentinel
