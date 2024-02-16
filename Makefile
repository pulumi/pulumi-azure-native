MAKEFLAGS    := --warn-undefined-variables

PROJECT         := github.com/pulumi/pulumi-azure-native
PROVIDER        := pulumi-resource-azure-native
CODEGEN         := pulumi-gen-azure-native

WORKING_DIR     := $(shell pwd)

PROVIDER_PKG    := $(shell find provider/pkg -type f)
SPECS           := $(shell find azure-rest-api-specs/specification/*/resource-manager -type f -name "*.json" ! -path "**/examples/**")

# Fail fast if the specs submodule doesn't exist
ifeq (,$(SPECS))
$(error Specs missing! Checkout submodules before building: git submodule update --init --recursive)
endif

JAVA_GEN        := pulumi-java-gen

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOEXE ?= $(shell go env GOEXE)
LOCAL_PROVIDER_FILENAME := $(PROVIDER)$(GOEXE)
export GOWORK := off
# Only use explicitly installed plugins - this is to avoid using any ambient plugins from the PATH
export PULUMI_IGNORE_AMBIENT_PLUGINS = true

ifeq ($(CI)$(PROVIDER_VERSION),true)
$(error PROVIDER_VERSION must be set in CI environments)
endif

# Input during CI using `make [TARGET] PROVIDER_VERSION=""` or by setting a PROVIDER_VERSION environment variable
# Local builds will just used this fixed default version unless specified
PROVIDER_VERSION ?= 2.0.0-alpha.0+dev
# Ensure the leading "v" is removed - use this normalised version everywhere rather than the raw input to ensure consistency.
# These variables are lazy (no `:`) so they're not calculated until after the dependency is installed
VERSION_GENERIC = $(shell bin/pulumictl convert-version -l generic -v "$(PROVIDER_VERSION)")
VERSION_FLAGS   = -ldflags "-X github.com/pulumi/pulumi-azure-native/v2/provider/pkg/version.Version=${VERSION_GENERIC}"

# Ensure make directory exists
# For targets which either don't generate a single file output, or the file is committed, we use a "sentinel"
# file within `.make/` to track the staleness of the target and only rebuild when needed. At the end of each
# relevant target we run `@touch $@` to update the file which is the name of the target.
_ := $(shell mkdir -p .make)

.PHONY: default ensure dist
default: provider build_sdks
ensure: bin/pulumictl .make/provider_mod_download
dist: dist/pulumi-azure-native_$(VERSION_GENERIC)_checksums.txt dist/docs-schema.json

# Binaries
.PHONY: codegen provider
codegen: bin/$(CODEGEN)
provider: bin/$(LOCAL_PROVIDER_FILENAME)

.PHONY: install_provider
install_provider: .make/install_provider

.PHONY: provider_prebuild
provider_prebuild: .make/provider_prebuild

# We don't include v2 here yet as this is executed on the nightly updates
.PHONY: schema generate_schema generate_docs
schema: bin/schema-full.json
generate_schema: bin/schema-full.json
generate_docs: provider/cmd/pulumi-resource-azure-native/schema.json

.PHONY: generate generate_java generate_nodejs generate_python generate_dotnet generate_go
generate: generate_java generate_nodejs generate_python generate_dotnet generate_go
generate_java: .make/generate_java
generate_nodejs: .make/generate_nodejs
generate_python: .make/generate_python
generate_dotnet: .make/generate_dotnet
generate_go: .make/generate_go_local

.PHONY: local_generate_code
local_generate_code: generate_java
local_generate_code: generate_nodejs
local_generate_code: generate_python
local_generate_code: generate_dotnet
local_generate_code: generate_go
local_generate: generate_schema local_generate_code generate_docs

.PHONY: build only_build build_sdks build_nodejs build_python build_dotnet build_java build_go
build: codegen local_generate provider build_sdks
# Required for the codegen action that runs in pulumi/pulumi
only_build: build
build_sdks: build_nodejs build_dotnet build_python build_go build_java
build_nodejs: .make/build_nodejs
build_python: .make/build_python
build_dotnet: .make/build_dotnet
build_java: .make/build_java
build_go: .make/build_go

.PHONY: install_dotnet_sdk install_python_sdk install_go_sdk install_java_sdk install_nodejs_sdk install_sdks
# Required by CI steps - some can be skipped
install_dotnet_sdk: .make/install_dotnet_sdk
install_python_sdk:
install_go_sdk:
install_java_sdk:
install_nodejs_sdk: .make/install_nodejs_sdk
install_sdks: install_dotnet_sdk install_nodejs_sdk

.PHONY: prepublish_go
prepublish_go: .make/prepublish_go

.PHONY: update_submodules
update_submodules:
	@for submodule in $$(git submodule status | awk {'print $$2'}); do \
		echo "Updating submodule $$submodule" ; \
		(cd $$submodule && git checkout main && git pull origin main); \
	done
	rm ./azure-provider-versions/provider_list.json
	az provider list | jq 'map({ namespace: .namespace, resourceTypes: .resourceTypes | map({ resourceType: .resourceType, apiVersions: .apiVersions }) | sort_by(.resourceType) }) | sort_by(.namespace)' > ./azure-provider-versions/provider_list.json

# Use PROVIDER_TEST_TAGS=all to run all tests including examples integrate tests
PROVIDER_TEST_TAGS ?= unit # Default to unit tests only for quick local feedback
.PHONY: test_provider
test_provider: .make/provider_mod_download .make/provider_prebuild provider/cmd/$(PROVIDER)/*.go $(PROVIDER_PKG)
	cd provider && go test -v --tags=$(PROVIDER_TEST_TAGS) -coverprofile="coverage.txt" -coverpkg=./... ./...

.PHONY: lint_provider
lint_provider: .make/provider_mod_download provider/cmd/$(PROVIDER)/*.go $(PROVIDER_PKG)
	cd provider && GOGC=20 golangci-lint run -c ../.golangci.yml

.PHONY: clean
clean:
	rm -rf nuget
	rm -rf .make
	rm -rf bin
	rm -rf dist
	cd provider/cmd/pulumi-resource-azure-native && rm -f metadata-compact.json schema-full.json
	rm -rf sdk/dotnet/bin
	rm -rf sdk/dotnet/build sdk/dotnet/src
	rm -rf sdk/nodejs/bin
	rm -rf sdk/pulumi-azure-native-sdk
	rm -rf sdk/python/bin
	rm -rf sdk/java/.gradle
	if dotnet nuget list source | grep "$(WORKING_DIR)"; then \
		dotnet nuget remove source "$(WORKING_DIR)" \
	; fi

.PHONY: test test_dotnet test_python test_go test_nodejs
# Set TEST_TAGS to override -tags for tests
TEST_TAGS ?= all
# Set TEST_NAME to filter tests by name
TEST_NAME ?=
TEST_RUN =
ifneq ($(TEST_NAME),)
TEST_RUN = -run ^$(TEST_NAME)$$
endif
export PULUMI_LOCAL_NUGET=$(WORKING_DIR)/nuget
test: provider install_sdks
	cd examples && go test -v -tags=$(TEST_TAGS) -timeout 2h $(TEST_RUN)
test_dotnet: provider build_dotnet install_dotnet_sdk
	cd examples && go test -v -tags=dotnet -timeout 2h $(TEST_RUN)
test_python: provider build_python
	cd examples && go test -v -tags=python -timeout 2h $(TEST_RUN)
test_go: provider generate_go
	cd examples && go test -v -tags=go -timeout 2h $(TEST_RUN)
test_nodejs: provider install_nodejs_sdk
	cd examples && go test -v -tags=nodejs -timeout 2h $(TEST_RUN)

.PHONY: schema_squeeze
schema_squeeze: bin/$(CODEGEN) bin/schema-tools bin/schema-full.json
	bin/$(CODEGEN) raw-schema $(VERSION_GENERIC)
	./bin/schema-tools squeeze -s bin/raw-schema.json --out versions/v2-removed-resources.json

.PHONY: explode_schema
explode_schema: dist/docs-schema.json

.PHONY: upgrade_tools upgrade_java upgrade_pulumi upgrade_pulumictl upgrade_schematools
upgrade_tools: upgrade_java upgrade_pulumi upgrade_pulumictl upgrade_schematools
upgrade_java:
	gh release list --repo pulumi/pulumi-java --exclude-drafts --exclude-pre-releases --limit 1 | cut -f1 > .pulumi-java-gen.version
upgrade_pulumi:
	gh release list --repo pulumi/pulumi --exclude-drafts --exclude-pre-releases --limit 1 | cut -f1 | sed 's/^v//' > .pulumi.version
upgrade_pulumictl:
	gh release list --repo pulumi/pulumictl --exclude-drafts --exclude-pre-releases --limit 1 | cut -f1 | sed 's/^v//' > .pulumictl.version
upgrade_schematools:
	gh release list --repo pulumi/schema-tools --exclude-drafts --exclude-pre-releases --limit 1 | cut -f1 | sed 's/^v//' > .schema-tools.version

# --------- File-based targets --------- #

.pulumi/bin/pulumi: PULUMI_VERSION := $(shell cat .pulumi.version)
.pulumi/bin/pulumi: HOME := $(WORKING_DIR)
.pulumi/bin/pulumi: .pulumi.version
	curl -fsSL https://get.pulumi.com | sh -s -- --version "$(PULUMI_VERSION)"

# Download local copy of pulumictl based on the version in .pulumictl.version
# Anywhere which uses VERSION_GENERIC or VERSION_FLAGS should depend on bin/pulumictl
bin/pulumictl: PULUMICTL_VERSION := $(shell cat .pulumictl.version)
bin/pulumictl: PLAT := $(shell go version | sed -En "s/go version go.* (.*)\/(.*)/\1-\2/p")
bin/pulumictl: PULUMICTL_URL := "https://github.com/pulumi/pulumictl/releases/download/v$(PULUMICTL_VERSION)/pulumictl-v$(PULUMICTL_VERSION)-$(PLAT).tar.gz"
bin/pulumictl: .pulumictl.version
	@echo "Installing pulumictl"
	@mkdir -p bin
	wget -q -O - "$(PULUMICTL_URL)" | tar -xzf - -C $(WORKING_DIR)/bin pulumictl
	@touch bin/pulumictl
	@echo "pulumictl" $$(./bin/pulumictl version)

bin/pulumi-java-gen: .pulumi-java-gen.version bin/pulumictl
	@mkdir -p bin
	bin/pulumictl download-binary -n pulumi-language-java -v $(shell cat .pulumi-java-gen.version) -r pulumi/pulumi-java

# Download local copy of schema-tools based on the version in .schema-tools.version
bin/schema-tools: SCHEMA_TOOLS_VERSION := $(shell cat .schema-tools.version)
bin/schema-tools: PLAT := $(shell go version | sed -En "s/go version go.* (.*)\/(.*)/\1-\2/p")
bin/schema-tools: SCHEMA_TOOLS_URL := "https://github.com/pulumi/schema-tools/releases/download/v$(SCHEMA_TOOLS_VERSION)/schema-tools-v$(SCHEMA_TOOLS_VERSION)-$(PLAT).tar.gz"
bin/schema-tools: .schema-tools.version
	@echo "Installing schema-tools"
	@mkdir -p bin
	wget -q -O - "$(SCHEMA_TOOLS_URL)" | tar -xzf - -C $(WORKING_DIR)/bin schema-tools
	@touch bin/schema-tools
	@echo "schema-tools" $$(./bin/schema-tools version)

dist/docs-schema.json: bin/schema-full.json
	rm -rf bin/schema
	mkdir -p bin/schema
	yarn install
	yarn schema explode --schema bin/schema-full.json --outDir bin/schema
	# Write docs schema over the top so we include examples
	yarn schema explode --schema provider/cmd/pulumi-resource-azure-native/schema.json --outDir bin/schema
	# Combine all the schemas into one
	mkdir -p dist
	yarn schema implode --cwd bin/schema --outFile dist/docs-schema.json

bin/$(CODEGEN): bin/pulumictl .make/provider_mod_download provider/cmd/$(CODEGEN)/* $(PROVIDER_PKG)
	cd provider && go build -o $(WORKING_DIR)/bin/$(CODEGEN) $(VERSION_FLAGS) $(PROJECT)/v2/provider/cmd/$(CODEGEN)

# Writes schema-full.json and metadata-compact.json to bin/
# Also re-calculates files in versions/ at same time
bin/schema-full.json bin/metadata-compact.json &: bin/$(CODEGEN) $(SPECS) azure-provider-versions/provider_list.json versions/v1-lock.json versions/v2-config.yaml versions/v2-spec.yaml versions/v2-removed-resources.json
	bin/$(CODEGEN) schema $(VERSION_GENERIC)

# Docs schema
provider/cmd/pulumi-resource-azure-native/schema.json: bin/$(CODEGEN) $(SPECS) versions/v1-lock.json versions/v2-config.yaml versions/v2-removed-resources.json
	bin/$(CODEGEN) docs $(VERSION_GENERIC)

bin/$(LOCAL_PROVIDER_FILENAME): bin/pulumictl .make/provider_mod_download provider/cmd/$(PROVIDER)/*.go .make/provider_prebuild $(PROVIDER_PKG)
	cd provider && \
		CGO_ENABLED=0 go build -o $(WORKING_DIR)/bin/$(LOCAL_PROVIDER_FILENAME) $(VERSION_FLAGS) $(PROJECT)/v2/provider/cmd/$(PROVIDER)

bin/linux-amd64/$(PROVIDER): TARGET := linux-amd64
bin/linux-arm64/$(PROVIDER): TARGET := linux-arm64
bin/darwin-amd64/$(PROVIDER): TARGET := darwin-amd64
bin/darwin-arm64/$(PROVIDER): TARGET := darwin-arm64
bin/windows-amd64/$(PROVIDER).exe: TARGET := windows-amd64
bin/%/$(PROVIDER) bin/%/$(PROVIDER).exe: bin/pulumictl .make/provider_mod_download provider/cmd/$(PROVIDER)/*.go .make/provider_prebuild $(PROVIDER_PKG)
	@# check the TARGET is set
	test $(TARGET)
	cd provider && \
		export GOOS=$$(echo "$(TARGET)" | cut -d "-" -f 1) && \
		export GOARCH=$$(echo "$(TARGET)" | cut -d "-" -f 2) && \
		CGO_ENABLED=0 go build -o ${WORKING_DIR}/$@ $(VERSION_FLAGS) $(PROJECT)/v2/provider/cmd/$(PROVIDER)

dist/$(PROVIDER)-v$(PROVIDER_VERSION)-linux-amd64.tar.gz: bin/linux-amd64/$(PROVIDER)
dist/$(PROVIDER)-v$(PROVIDER_VERSION)-linux-arm64.tar.gz: bin/linux-arm64/$(PROVIDER)
dist/$(PROVIDER)-v$(PROVIDER_VERSION)-darwin-amd64.tar.gz: bin/darwin-amd64/$(PROVIDER)
dist/$(PROVIDER)-v$(PROVIDER_VERSION)-darwin-arm64.tar.gz: bin/darwin-arm64/$(PROVIDER)
dist/$(PROVIDER)-v$(PROVIDER_VERSION)-windows-amd64.tar.gz: bin/windows-amd64/$(PROVIDER).exe
dist/$(PROVIDER)-v$(PROVIDER_VERSION)-%.tar.gz:
	@mkdir -p dist
	@# $< is the last dependency (the binary path from above)
	tar --gzip -cf $@ README.md LICENSE -C $$(dirname $<) .

dist/pulumi-azure-native_$(VERSION_GENERIC)_checksums.txt: dist/$(PROVIDER)-v$(PROVIDER_VERSION)-linux-amd64.tar.gz
dist/pulumi-azure-native_$(VERSION_GENERIC)_checksums.txt: dist/$(PROVIDER)-v$(PROVIDER_VERSION)-linux-arm64.tar.gz
dist/pulumi-azure-native_$(VERSION_GENERIC)_checksums.txt: dist/$(PROVIDER)-v$(PROVIDER_VERSION)-darwin-amd64.tar.gz
dist/pulumi-azure-native_$(VERSION_GENERIC)_checksums.txt: dist/$(PROVIDER)-v$(PROVIDER_VERSION)-darwin-arm64.tar.gz
dist/pulumi-azure-native_$(VERSION_GENERIC)_checksums.txt: dist/$(PROVIDER)-v$(PROVIDER_VERSION)-windows-amd64.tar.gz
	cd dist && shasum *.tar.gz > pulumi-azure-native_$(VERSION_GENERIC)_checksums.txt

# --------- Sentinel targets --------- #

.make/provider_mod_download: provider/go.mod provider/go.sum
	cd provider && go mod download
	@touch $@

.make/provider_prebuild: bin/schema-full.json bin/metadata-compact.json versions/v2-lock.json
	cp bin/schema-full.json provider/cmd/$(PROVIDER)
	cp bin/metadata-compact.json provider/cmd/$(PROVIDER)
	@# For API version lookups at run time
	cp versions/v2-lock.json provider/pkg/versionLookup/
	@touch $@

define FAKE_MODULE
module fake_module // Exclude this directory from Go tools

go 1.17
endef

export FAKE_MODULE

# We use the docs schema for java but don't depend on it because it changes on every generation
.make/generate_java: bin/pulumictl bin/pulumi-java-gen
	@mkdir -p sdk/java
	rm -rf $$(find sdk/java -mindepth 1 -maxdepth 1)
	bin/$(JAVA_GEN) generate --schema provider/cmd/$(PROVIDER)/schema.json --out sdk/java --build gradle-nexus
	echo "$$FAKE_MODULE" | sed 's/fake_module/fake_java_module/g' > sdk/java/go.mod
	@touch $@

.make/generate_nodejs: bin/pulumictl .pulumi/bin/pulumi bin/schema-full.json
	mkdir -p sdk/nodejs
	rm -rf $$(find sdk/nodejs -mindepth 1 -maxdepth 1 ! -name "go.mod")
	.pulumi/bin/pulumi package gen-sdk bin/schema-full.json --language nodejs
	echo "$$FAKE_MODULE" | sed 's/fake_module/fake_nodejs_module/g' > sdk/nodejs/go.mod
	sed -i.bak -e "s/sourceMap/inlineSourceMap/g" sdk/nodejs/tsconfig.json
	rm sdk/nodejs/tsconfig.json.bak
	@touch $@

.make/generate_python: bin/pulumictl .pulumi/bin/pulumi bin/schema-full.json
	mkdir -p sdk/python
	rm -rf $$(find sdk/python -mindepth 1 -maxdepth 1 ! -name "go.mod")
	.pulumi/bin/pulumi package gen-sdk bin/schema-full.json --language python
	echo "$$FAKE_MODULE" | sed 's/fake_module/fake_python_module/g' > sdk/python/go.mod
	cp README.md sdk/python
	@touch $@

.make/generate_dotnet: bin/pulumictl .pulumi/bin/pulumi bin/schema-full.json
	mkdir -p sdk/dotnet
	rm -rf $$(find sdk/dotnet -mindepth 1 -maxdepth 1 ! -name "go.mod")
	.pulumi/bin/pulumi package gen-sdk bin/schema-full.json --language dotnet
	echo "$$FAKE_MODULE" | sed 's/fake_module/fake_dotnet_module/g' > sdk/dotnet/go.mod
	sed -i.bak -e "s/<\/Nullable>/<\/Nullable>\n    <UseSharedCompilation>false<\/UseSharedCompilation>/g" sdk/dotnet/Pulumi.AzureNative.csproj
	rm sdk/dotnet/Pulumi.AzureNative.csproj.bak
	@touch $@

.make/generate_go_local: bin/pulumictl bin/$(CODEGEN) bin/schema-full.json
	@mkdir -p sdk/pulumi-azure-native-sdk
	@# Unmark this is as an up-to-date local build
	rm -f .make/prepublish_go
	rm -rf $$(find sdk/pulumi-azure-native-sdk -mindepth 1 -maxdepth 1 ! -name ".git")
	bin/$(CODEGEN) go $(VERSION_GENERIC)
	@# Tidy up all go.mod files
	find sdk/pulumi-azure-native-sdk -type d -maxdepth 1 -exec sh -c "cd \"{}\" && go mod tidy" \;
	@touch $@

.make/prepublish_go:
	@# Unmark this is as an up-to-date local build
	rm -f .make/generate_go_local
	@# Remove go module replacements which are added for local testing
	@# Note: must use `sed -i -e` to be portable - but leaves go.mod-e behind on macos
	find sdk/pulumi-azure-native-sdk -maxdepth 2 -type f -name go.mod -exec sed -i -e '/replace github\.com\/pulumi\/pulumi-azure-native-sdk/d' {} \;
	@# Remove sed backup files if using older sed versions
	find sdk/pulumi-azure-native-sdk -maxdepth 2 -type f -name go.mod-e -delete
	@# Delete go.sum files as these are not used at the point of publishing.
	@# This is because we depend on the root package which will come from the same release commit, that doesn't yet exist.
	find sdk/pulumi-azure-native-sdk -maxdepth 2 -type f -name go.sum -delete
	cp README.md LICENSE sdk/pulumi-azure-native-sdk/
	@touch $@

# Used by build* targets

.make/nodejs_yarn_install: .make/generate_nodejs sdk/nodejs/package.json
	yarn install --cwd sdk/nodejs
	@touch $@

.make/build_nodejs: VERSION_JS = $(shell bin/pulumictl convert-version -l javascript -v "$(VERSION_GENERIC)")
.make/build_nodejs: bin/pulumictl .make/nodejs_yarn_install
	cd sdk/nodejs/ && \
		NODE_OPTIONS=--max-old-space-size=12288 yarn run tsc --diagnostics --incremental && \
		cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/ && \
		mkdir -p bin/scripts && \
		sed -i.bak -e "s/\$${VERSION}/$(VERSION_JS)/g" ./bin/package.json
	@touch $@

.make/build_python: VERSION_PYTHON = $(shell bin/pulumictl convert-version -l python -v "$(VERSION_GENERIC)")
.make/build_python: bin/pulumictl .make/generate_python
	cd sdk/python && \
		git clean -fxd && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		sed -i.bak -e 's/^  version = .*/  version = "$(VERSION_PYTHON)"/g' ./bin/pyproject.toml && \
		rm ./bin/pyproject.toml.bak && \
		rm ./bin/go.mod && \
		python3 -m venv venv && \
		./venv/bin/python -m pip install build && \
		cd ./bin && \
		../venv/bin/python -m build .
	@touch $@

.make/build_dotnet: VERSION_DOTNET = $(shell bin/pulumictl convert-version -l dotnet -v "$(PROVIDER_VERSION)")
.make/build_dotnet: bin/pulumictl .make/generate_dotnet
	cd sdk/dotnet && \
		echo "azure-native\n$(VERSION_DOTNET)" >version.txt && \
		dotnet build /p:Version=$(VERSION_DOTNET)
	@touch $@

.make/build_java: bin/pulumictl .make/generate_java
	cd sdk/java/ && \
		gradle --console=plain -Pversion=$(VERSION_GENERIC) build
	@touch $@

.make/build_go: .make/generate_go_local
	find sdk/pulumi-azure-native-sdk -type d -maxdepth 1 -exec sh -c "cd \"{}\" && go build" \;
	@touch $@

# Used by install* targets

.make/install_nodejs_sdk: .make/build_nodejs
	yarn link --cwd sdk/nodejs/bin
	@touch $@

.make/install_dotnet_sdk: .make/build_dotnet
	mkdir -p nuget
	find sdk/dotnet/bin -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;
	if ! dotnet nuget list source | grep ${WORKING_DIR}; then \
		dotnet nuget add source ${WORKING_DIR}/nuget --name ${WORKING_DIR} \
	; fi
	@touch $@

.make/install_provider: bin/pulumictl .make/provider_mod_download provider/cmd/$(PROVIDER)/* $(PROVIDER_PKG)
	cd provider && go install $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(PROVIDER)
	@touch $@
