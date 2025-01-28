MAKEFLAGS    := --warn-undefined-variables

PROJECT         := github.com/pulumi/pulumi-azure-native
PROVIDER        := pulumi-resource-azure-native
CODEGEN         := pulumi-gen-azure-native

WORKING_DIR     := $(shell pwd)

PROVIDER_PKG    := $(shell find provider/pkg -type f)
SPECS           := .git/modules/azure-rest-api-specs/HEAD

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
# Check version doesn't start with a "v" - this is a common mistake
ifeq ($(shell echo $(PROVIDER_VERSION) | cut -c1),v)
$(error PROVIDER_VERSION should not start with a "v")
endif

VERSION_FLAGS   = -ldflags "-X github.com/pulumi/pulumi-azure-native/v2/provider/pkg/version.Version=${PROVIDER_VERSION}"
MAJOR_VERSION   = $(shell echo $(PROVIDER_VERSION) | cut -d. -f1)
PREVIOUS_MAJOR_VERSION = $(shell echo $(MAJOR_VERSION)-1 | bc)
NEXT_MAJOR_VERSION = $(shell echo $(MAJOR_VERSION)+1 | bc)

# Set these variables to enable signing of the windows binary
AZURE_SIGNING_CLIENT_ID ?=
AZURE_SIGNING_CLIENT_SECRET ?=
AZURE_SIGNING_TENANT_ID ?=
AZURE_SIGNING_KEY_VAULT_URI ?=

# Ensure make directory exists
# For targets which either don't generate a single file output, or the file is committed, we use a "sentinel"
# file within `.make/` to track the staleness of the target and only rebuild when needed. At the end of each
# relevant target we run `@touch $@` to update the file which is the name of the target.
_ := $(shell mkdir -p .make)

.PHONY: default ensure dist
default: provider build_sdks
ensure: bin/pulumictl .make/provider_mod_download
dist: dist/pulumi-azure-native_$(PROVIDER_VERSION)_checksums.txt dist/docs-schema.json

# Binaries
.PHONY: codegen provider
codegen: bin/$(CODEGEN)
provider: bin/$(LOCAL_PROVIDER_FILENAME)

.PHONY: install_provider
install_provider: .make/install_provider

.PHONY: provider_prebuild
provider_prebuild: .make/provider_prebuild

.PHONY: prebuild
prebuild: .make/prebuild

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
	rm versions/az-provider-list.json
	# Use query option to remove all except specific fields and sort to reduce diff noise
	az provider list --query 'sort_by([*].{ namespace: namespace, resourceTypes: sort_by(resourceTypes[*].{ resourceType: resourceType, defaultApiVersion: defaultApiVersion, apiVersions: apiVersions, locations: locations }, &resourceType) }, &namespace)' > versions/az-provider-list.json

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
	rm -rf .make/*
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
schema_squeeze: bin/$(CODEGEN)
	bin/$(CODEGEN) squeeze

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

bin/$(CODEGEN): .make/prebuild .make/provider_mod_download provider/cmd/$(CODEGEN)/* $(PROVIDER_PKG)
	cd provider && go build -o $(WORKING_DIR)/bin/$(CODEGEN) $(VERSION_FLAGS) $(PROJECT)/v2/provider/cmd/$(CODEGEN)

# Writes schema-full.json and metadata-compact.json to bin/
# Also re-calculates files in versions/ at same time
bin/schema-full.json bin/metadata-compact.json &: bin/$(CODEGEN) $(SPECS) versions/az-provider-list.json versions/v${PREVIOUS_MAJOR_VERSION}.yaml versions/v${MAJOR_VERSION}-config.yaml versions/v${MAJOR_VERSION}-spec.yaml versions/v${MAJOR_VERSION}-removed.json versions/v${MAJOR_VERSION}-removed-resources.json versions/v${NEXT_MAJOR_VERSION}-removed-resources.json
	bin/$(CODEGEN) schema

# Docs schema - treat as phony becasuse it's committed so we always need to rebuild it.
.PHONY: provider/cmd/pulumi-resource-azure-native/schema.json
provider/cmd/pulumi-resource-azure-native/schema.json: bin/$(CODEGEN) $(SPECS) versions/v${PREVIOUS_MAJOR_VERSION}.yaml versions/v${MAJOR_VERSION}-config.yaml versions/v${MAJOR_VERSION}-removed-resources.json
	bin/$(CODEGEN) docs

bin/$(LOCAL_PROVIDER_FILENAME): .make/prebuild .make/provider_mod_download provider/cmd/$(PROVIDER)/*.go .make/provider_prebuild $(PROVIDER_PKG)
	cd provider && \
		CGO_ENABLED=0 go build -o $(WORKING_DIR)/bin/$(LOCAL_PROVIDER_FILENAME) $(VERSION_FLAGS) $(PROJECT)/v2/provider/cmd/$(PROVIDER)

bin/linux-amd64/$(PROVIDER): TARGET := linux-amd64
bin/linux-arm64/$(PROVIDER): TARGET := linux-arm64
bin/darwin-amd64/$(PROVIDER): TARGET := darwin-amd64
bin/darwin-arm64/$(PROVIDER): TARGET := darwin-arm64
bin/windows-amd64/$(PROVIDER).exe: TARGET := windows-amd64
bin/%/$(PROVIDER) bin/%/$(PROVIDER).exe: .make/provider_mod_download .make/prebuild provider/cmd/$(PROVIDER)/*.go .make/provider_prebuild $(PROVIDER_PKG)
	@# check the TARGET is set
	test $(TARGET)
	cd provider && \
		export GOOS=$$(echo "$(TARGET)" | cut -d "-" -f 1) && \
		export GOARCH=$$(echo "$(TARGET)" | cut -d "-" -f 2) && \
		CGO_ENABLED=0 go build -o ${WORKING_DIR}/$@ $(VERSION_FLAGS) $(PROJECT)/v2/provider/cmd/$(PROVIDER)
	@# Only sign if configured. Test variables set by joining with | between and looking for || showing at least one variable is empty
	@# Move the binary to a temporary location and sign it there to avoid the target being up-to-date if signing fails
	set -e; \
	if [[ "${TARGET}" = "windows-amd64" ]]; then \
		if [[ "|${AZURE_SIGNING_CLIENT_ID}|${AZURE_SIGNING_CLIENT_SECRET}|${AZURE_SIGNING_TENANT_ID}|${AZURE_SIGNING_KEY_VAULT_URI}|" == *"||"* ]]; then \
			echo "Skipping signing as required configuration not set: AZURE_SIGNING_CLIENT_ID, AZURE_SIGNING_CLIENT_SECRET, AZURE_SIGNING_TENANT_ID, AZURE_SIGNING_KEY_VAULT_URI"; \
			echo "To rebuild with signing delete the unsigned $@ and re-run with the fixed configuration"; \
		else \
			mv $@ $@.unsigned; \
			az login --service-principal \
				--username "${AZURE_SIGNING_CLIENT_ID}" \
				--password "${AZURE_SIGNING_CLIENT_SECRET}" \
				--tenant "${AZURE_SIGNING_TENANT_ID}" \
				--output none; \
			ACCESS_TOKEN=$$(az account get-access-token --resource "https://vault.azure.net" | jq -r .accessToken); \
			wget https://github.com/ebourg/jsign/releases/download/6.0/jsign-6.0.jar --output-document=bin/jsign-6.0.jar; \
			java -jar bin/jsign-6.0.jar \
				--storetype AZUREKEYVAULT \
				--keystore "PulumiCodeSigning" \
				--url "${AZURE_SIGNING_KEY_VAULT_URI}" \
				--storepass "$${ACCESS_TOKEN}" \
				$@.unsigned; \
			mv $@.unsigned $@; \
			az logout; \
		fi; \
	fi

dist/$(PROVIDER)-v$(PROVIDER_VERSION)-linux-amd64.tar.gz: bin/linux-amd64/$(PROVIDER)
dist/$(PROVIDER)-v$(PROVIDER_VERSION)-linux-arm64.tar.gz: bin/linux-arm64/$(PROVIDER)
dist/$(PROVIDER)-v$(PROVIDER_VERSION)-darwin-amd64.tar.gz: bin/darwin-amd64/$(PROVIDER)
dist/$(PROVIDER)-v$(PROVIDER_VERSION)-darwin-arm64.tar.gz: bin/darwin-arm64/$(PROVIDER)
dist/$(PROVIDER)-v$(PROVIDER_VERSION)-windows-amd64.tar.gz: bin/windows-amd64/$(PROVIDER).exe
dist/$(PROVIDER)-v$(PROVIDER_VERSION)-%.tar.gz:
	@mkdir -p dist
	@# $< is the last dependency (the binary path from above)
	tar --gzip -cf $@ README.md LICENSE -C $$(dirname $<) .

dist/pulumi-azure-native_$(PROVIDER_VERSION)_checksums.txt: dist/$(PROVIDER)-v$(PROVIDER_VERSION)-linux-amd64.tar.gz
dist/pulumi-azure-native_$(PROVIDER_VERSION)_checksums.txt: dist/$(PROVIDER)-v$(PROVIDER_VERSION)-linux-arm64.tar.gz
dist/pulumi-azure-native_$(PROVIDER_VERSION)_checksums.txt: dist/$(PROVIDER)-v$(PROVIDER_VERSION)-darwin-amd64.tar.gz
dist/pulumi-azure-native_$(PROVIDER_VERSION)_checksums.txt: dist/$(PROVIDER)-v$(PROVIDER_VERSION)-darwin-arm64.tar.gz
dist/pulumi-azure-native_$(PROVIDER_VERSION)_checksums.txt: dist/$(PROVIDER)-v$(PROVIDER_VERSION)-windows-amd64.tar.gz
	cd dist && shasum *.tar.gz > pulumi-azure-native_$(PROVIDER_VERSION)_checksums.txt

# --------- Sentinel targets --------- #

.make/provider_mod_download: provider/go.mod provider/go.sum
	cd provider && go mod download
	@touch $@

.make/provider_prebuild: .make/prebuild bin/schema-full.json bin/metadata-compact.json versions/v${MAJOR_VERSION}.yaml
	cp bin/schema-full.json provider/cmd/$(PROVIDER)
	cp bin/metadata-compact.json provider/cmd/$(PROVIDER)
	cp -v versions/v${MAJOR_VERSION}.yaml provider/pkg/versionLookup/default-versions.yaml
	@touch $@

.make/prebuild: .pulumi/bin/pulumi
	cp -v versions/v${MAJOR_VERSION}.yaml provider/pkg/versionLookup/default-versions.yaml
	@touch $@

define FAKE_MODULE
module fake_module // Exclude this directory from Go tools

go 1.17
endef

export FAKE_MODULE

# We use the docs schema for java but don't depend on it because it changes on every generation
.make/generate_java: bin/pulumi-java-gen
	@mkdir -p sdk/java
	rm -rf $$(find sdk/java -mindepth 1 -maxdepth 1)
	bin/$(JAVA_GEN) generate --schema provider/cmd/$(PROVIDER)/schema.json --out sdk/java --build gradle-nexus
	echo "$$FAKE_MODULE" | sed 's/fake_module/fake_java_module/g' > sdk/java/go.mod
	@touch $@

.make/generate_nodejs: .pulumi/bin/pulumi bin/schema-full.json
	mkdir -p sdk/nodejs
	rm -rf $$(find sdk/nodejs -mindepth 1 -maxdepth 1 ! -name "go.mod")
	.pulumi/bin/pulumi package gen-sdk bin/schema-full.json --language nodejs
	echo "$$FAKE_MODULE" | sed 's/fake_module/fake_nodejs_module/g' > sdk/nodejs/go.mod
	sed -i.bak -e "s/sourceMap/inlineSourceMap/g" sdk/nodejs/tsconfig.json
	rm sdk/nodejs/tsconfig.json.bak
	@touch $@

.make/generate_python: .pulumi/bin/pulumi bin/schema-full.json
	mkdir -p sdk/python
	rm -rf $$(find sdk/python -mindepth 1 -maxdepth 1 ! -name "go.mod")
	.pulumi/bin/pulumi package gen-sdk bin/schema-full.json --language python
	echo "$$FAKE_MODULE" | sed 's/fake_module/fake_python_module/g' > sdk/python/go.mod
	cp README.md sdk/python
	@touch $@

.make/generate_dotnet: .pulumi/bin/pulumi bin/schema-full.json
	mkdir -p sdk/dotnet
	rm -rf $$(find sdk/dotnet -mindepth 1 -maxdepth 1 ! -name "go.mod")
	.pulumi/bin/pulumi package gen-sdk bin/schema-full.json --language dotnet
	echo "$$FAKE_MODULE" | sed 's/fake_module/fake_dotnet_module/g' > sdk/dotnet/go.mod
	sed -i.bak -e "s/<\/Nullable>/<\/Nullable>\n    <UseSharedCompilation>false<\/UseSharedCompilation>/g" sdk/dotnet/Pulumi.AzureNative.csproj
	rm sdk/dotnet/Pulumi.AzureNative.csproj.bak
	echo "azure-native\n$(PROVIDER_VERSION)" > sdk/dotnet/version.txt
	@touch $@

.make/generate_go_local: bin/$(CODEGEN) bin/schema-full.json
	@mkdir -p sdk/pulumi-azure-native-sdk
	@# Unmark this is as an up-to-date local build
	rm -f .make/prepublish_go
	rm -rf $$(find sdk/pulumi-azure-native-sdk -mindepth 1 -maxdepth 1 ! -name ".git")
	bin/$(CODEGEN) go
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

.make/build_nodejs: .make/nodejs_yarn_install
	cd sdk/nodejs/ && \
		NODE_OPTIONS=--max-old-space-size=12288 yarn run tsc --diagnostics --incremental && \
		cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/
	@touch $@

.make/build_python: .make/generate_python
	cd sdk/python && \
		git clean -fxd && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		rm ./bin/go.mod && \
		python3 -m venv venv && \
		./venv/bin/python -m pip install build && \
		cd ./bin && \
		../venv/bin/python -m build .
	@touch $@

.make/build_dotnet: .make/generate_dotnet
	cd sdk/dotnet && dotnet build
	@touch $@

.make/build_java: .make/generate_java
	cd sdk/java/ && \
		gradle --console=plain -Pversion=$(PROVIDER_VERSION) build
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

.make/install_provider: .make/provider_mod_download provider/cmd/$(PROVIDER)/* $(PROVIDER_PKG)
	cd provider && go install $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(PROVIDER)
	@touch $@
