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

# Calls to pulumictl are slow, so we just do the call once if required and cache
# Write to a new file then compare and swap to avoid updating the timestamp, which causes all dependant targets to rebuild
VERSION         = $(shell bin/pulumictl get version)
VERSION_FLAGS   = -ldflags "-X github.com/pulumi/pulumi-azure-native/provider/pkg/version.Version=${VERSION}"

ensure: init_submodules bin/pulumictl provider/.mod_download.sentinel
	@jq --version > /dev/null

build: init_submodules clean codegen local_generate provider build_sdks install_sdks
build_sdks: build_nodejs build_dotnet build_python build_go build_java
install_sdks: install_dotnet_sdk install_python_sdk install_nodejs_sdk
arm2pulumi: bin/arm2pulumi
codegen: bin/$(CODEGEN)
provider: bin/$(PROVIDER)
versioner: bin/pulumi-versioner-azure-native
versions: versions/spec.json versions/v1.json versions/v2.json versions/deprecated.json versions/pending.json versions/active.json
generate_schema: provider/cmd/$(PROVIDER)/schema-full.json
generate_docs: provider/cmd/$(PROVIDER)/schema.json
generate_java: sdk/java
generate_nodejs: sdk/nodejs
generate_python: sdk/python
generate_dotnet: sdk/dotnet
generate_go: sdk/go

# Required for the codegen action that runs in pulumi/pulumi
only_build: build

.PHONY: ensure build build_sdks install_sdks arm2pulumi codegen provider versioner versions
.PHONY: generate_schema generate_docs generate_java
.PHONY: init_submodules update_submodules local_generate_code local_generate arm2pulumi_coverage_report
.PHONY: test_provider lint_provider generate_nodejs build_nodejs generate_python build_python generate_dotnet

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

local_generate_code: clean bin/pulumi-java-gen bin/$(CODEGEN)
	bin/$(CODEGEN) schema,nodejs,dotnet,python,go,go-split $(VERSION)
	bin/$(JAVA_GEN) generate --schema provider/cmd/$(PROVIDER)/schema.json --out sdk/java --build gradle-nexus
	cd sdk/go/ && find . -type f -exec sed -i '' -e '/^\/\/.*/g' {} \;
	cd sdk/dotnet/ && \
	sed -i.bak -e "s/<\/Nullable>/<\/Nullable>\n    <UseSharedCompilation>false<\/UseSharedCompilation>/g" Pulumi.AzureNative.csproj && \
	rm Pulumi.AzureNative.csproj.bak && \
	cd ../nodejs/ && \
	sed -i.bak -e "s/sourceMap/inlineSourceMap/g" tsconfig.json && \
	rm tsconfig.json.bak
	echo "Finished generating."

local_generate: clean bin/pulumi-java-gen bin/pulumictl bin/$(CODEGEN)
	bin/$(CODEGEN) schema,docs,nodejs,dotnet,python,go,go-split $(VERSION)
	bin/$(JAVA_GEN) generate --schema provider/cmd/$(PROVIDER)/schema.json --out sdk/java --build gradle-nexus
	cd sdk/go/ && find . -type f -exec sed -i '' -e '/^\/\/.*/g' {} \;
	cd sdk/dotnet/ && \
	sed -i.bak -e "s/<\/Nullable>/<\/Nullable>\n    <UseSharedCompilation>false<\/UseSharedCompilation>/g" Pulumi.AzureNative.csproj && \
	rm Pulumi.AzureNative.csproj.bak && \
	cd ../nodejs/ && \
	sed -i.bak -e "s/sourceMap/inlineSourceMap/g" tsconfig.json && \
	rm tsconfig.json.bak
	echo "Finished generating."

arm2pulumi_coverage_report:
	(cd provider/pkg/arm2pulumi/internal/testdata && if [ ! -d azure-quickstart-templates ]; then git clone https://github.com/Azure/azure-quickstart-templates && cd azure-quickstart-templates && git checkout 3b2757465c2de537e333f5e2d1c3776c349b8483; fi)
	(cd provider && go test -v -tags=coverage -run TestQuickstartTemplateCoverage github.com/pulumi/pulumi-azure-native/provider/pkg/arm2pulumi/internal/test)

test_provider:
	(cd provider && go test -v $(PROVIDER_PKGS))

lint_provider: provider # lint the provider code
	cd provider && GOGC=20 golangci-lint run -c ../.golangci.yml

build_nodejs:
	cd sdk/nodejs/ && \
	yarn install && \
	NODE_OPTIONS=--max-old-space-size=8192 yarn run tsc --diagnostics && \
	cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/ && \
	sed -i.bak -e "s/\$${VERSION}/$(shell pulumictl get version --language javascript)/g" ./bin/package.json

build_python: bin/pulumictl
	cd sdk/python/ && \
	cp ../../README.md . && \
	python3 setup.py clean --all 2>/dev/null && \
	rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
	sed -i.bak -e 's/^VERSION = .*/VERSION = "$(shell pulumictl get version --language python)"/g' -e 's/^PLUGIN_VERSION = .*/PLUGIN_VERSION = "$(VERSION)"/g' ./bin/setup.py && \
	rm ./bin/setup.py.bak && \
	rm ./bin/go.mod && \
	cd ./bin && python3 setup.py build sdist

build_dotnet:
	cd sdk/dotnet/ && \
	echo "azure-native\n$(shell pulumictl get version --language dotnet)" >version.txt && \
	dotnet build /p:Version=$(shell pulumictl get version --language dotnet)

build_java: bin/pulumictl
	cd sdk/java/ && \
		gradle --console=plain -Pversion=$(VERSION) build

build_go:
	# Only building the top level packages and building 1 package at a time to avoid OOMing
	cd sdk/ && \
	GOGC=50 go list github.com/pulumi/pulumi-azure-native/sdk/go/azure/... | grep -v "latest\|\/v.*"$ | xargs -L 1 go build
	find sdk/pulumi-azure-native-sdk -type d -maxdepth 1 -exec sh -c "cd \"{}\" && go mod tidy && go build" \;

prepublish_go:
	@# Remove go module replacements which are added for local testing
	@# Note: must use `sed -i -e` to be portable - but leaves go.mod-e behind on macos
	find sdk/pulumi-azure-native-sdk -maxdepth 2 -type f -name go.mod -exec sed -i -e '/replace github\.com\/pulumi\/pulumi-azure-native-sdk /d' {} \;
	@# Remove sed backup files if using older sed versions
	find sdk/pulumi-azure-native-sdk -maxdepth 2 -type f -name go.mod-e -delete
	@# Delete go.sum files as these are not used at the point of publishing.
	@# This is because we depend on the root package which will come from the same release commit, that doesn't yet exist.
	find sdk/pulumi-azure-native-sdk -maxdepth 2 -type f -name go.sum -delete
	cp README.md LICENSE sdk/pulumi-azure-native-sdk/

clean:
	rm -rf $$(find sdk/nodejs -mindepth 1 -maxdepth 1 ! -name "go.mod")
	rm -rf $$(find sdk/python -mindepth 1 -maxdepth 1 ! -name "go.mod" ! -name "README.md")
	rm -rf $$(find sdk/dotnet -mindepth 1 -maxdepth 1 ! -name "go.mod")
	rm -rf $$(find sdk/java -mindepth 1 -maxdepth 1 ! -name "go.mod")
	rm -rf sdk/go/azure
	rm -rf sdk/schema

install_provider: bin/pulumictl provider/.mod_download.sentinel provider/cmd/$(PROVIDER)/* $(PROVIDER_PKG)
	(cd provider && go install $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(PROVIDER))

install_dotnet_sdk:
	mkdir -p nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;

install_python_sdk:

install_go_sdk:

install_java_sdk:

install_nodejs_sdk:
	yarn link --cwd sdk/nodejs/bin

test: export PULUMI_LOCAL_NUGET=${WORKING_DIR}/nuget
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
	@wget --quiet -O bin/pulumictl.tar.gz "$(PULUMICTL_URL)"
	@tar -zxf bin/pulumictl.tar.gz -C bin pulumictl
	@touch bin/pulumictl
	@echo "pulumictl" $$(./bin/pulumictl version)

bin/pulumi-java-gen: .pulumi-java-gen.version
	pulumictl download-binary -n pulumi-language-java -v $(shell cat .pulumi-java-gen.version) -r pulumi/pulumi-java

provider/.mod_download.sentinel: provider/go.mod provider/go.sum
	cd provider && GO111MODULE=on go mod download
	@touch provider/.mod_download.sentinel

bin/arm2pulumi: bin/pulumictl provider/.mod_download.sentinel provider/cmd/arm2pulumi/* $(PROVIDER_PKG)
	cd provider && go build -o $(WORKING_DIR)/bin/arm2pulumi $(VERSION_FLAGS) $(PROJECT)/provider/cmd/arm2pulumi

bin/$(CODEGEN): bin/pulumictl provider/.mod_download.sentinel provider/cmd/$(CODEGEN)/* $(PROVIDER_PKG)
	cd provider && go build -o $(WORKING_DIR)/bin/$(CODEGEN) $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(CODEGEN)

bin/$(PROVIDER): bin/pulumictl provider/.mod_download.sentinel provider/cmd/$(PROVIDER)/* $(PROVIDER_PKG)
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

sdk/nodejs/go.mod:
	echo "$$FAKE_MODULE" | sed 's/fake_module/fake_nodejs_module/g' > $@

sdk/java/go.mod:
	echo "$$FAKE_MODULE" | sed 's/fake_module/fake_java_module/g' > $@

sdk/dotnet/go.mod:
	echo "$$FAKE_MODULE" | sed 's/fake_module/fake_dotnet_module/g' > $@

sdk/python/go.mod:
	echo "$$FAKE_MODULE" | sed 's/fake_module/fake_python_module/g' > $@

sdk/java: bin/pulumi-java-gen provider/cmd/$(PROVIDER)/schema.json sdk/java/go.mod
	rm -rf $$(find sdk/java -mindepth 1 -maxdepth 1 ! -name "go.mod")
	bin/$(JAVA_GEN) generate --schema provider/cmd/$(PROVIDER)/schema.json --out sdk/java --build gradle-nexus
	@touch sdk/java

sdk/nodejs: sdk/nodejs/go.mod bin/pulumictl bin/$(CODEGEN)
	rm -rf $$(find sdk/nodejs -mindepth 1 -maxdepth 1 ! -name "go.mod")
	bin/$(CODEGEN) nodejs $(VERSION)
	sed -i.bak -e "s/sourceMap/inlineSourceMap/g" sdk/nodejs/tsconfig.json
	rm sdk/nodejs/tsconfig.json.bak
	@touch sdk/nodejs

sdk/python: sdk/python/go.mod bin/pulumictl bin/$(CODEGEN)
	rm -rf $$(find sdk/python -mindepth 1 -maxdepth 1 ! -name "go.mod")
	bin/$(CODEGEN) python $(VERSION)
	cp README.md sdk/python
	@touch sdk/python

sdk/dotnet: sdk/dotnet/go.mod bin/pulumictl bin/$(CODEGEN)
	rm -rf $$(find sdk/dotnet -mindepth 1 -maxdepth 1 ! -name "go.mod")
	bin/$(CODEGEN) dotnet $(VERSION)
	sed -i.bak -e "s/<\/Nullable>/<\/Nullable>\n    <UseSharedCompilation>false<\/UseSharedCompilation>/g" sdk/dotnet/Pulumi.AzureNative.csproj
	rm sdk/dotnet/Pulumi.AzureNative.csproj.bak
	@touch sdk/dotnet

sdk/go: bin/pulumictl
	rm -rf sdk/go/azure
	bin/$(CODEGEN) go,go-split $(VERSION)
	@# HACK: Strip all comments to make SDK smaller
	find sdk/go -type f -exec sed -i '' -e '/^\/\/.*/g' {} \;
	@touch sdk/go
