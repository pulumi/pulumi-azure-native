module github.com/pulumi/pulumi-azure-native/v2/provider

go 1.24.7

toolchain go1.24.10

require (
	github.com/Azure/azure-sdk-for-go v68.0.0+incompatible
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.20.0
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.13.1
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3 v3.0.0-beta.2
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault v1.5.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4 v4.2.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions v1.3.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage v1.8.1
	github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azkeys v1.4.0
	github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azsecrets v1.4.0
	github.com/Azure/azure-sdk-for-go/sdk/storage/azblob v1.6.3
	github.com/Azure/go-autorest/autorest v0.11.29
	github.com/blang/semver v3.5.1+incompatible
	github.com/brianvoe/gofakeit/v6 v6.16.0
	github.com/deckarep/golang-set/v2 v2.5.0
	github.com/gedex/inflector v0.0.0-20170307190818-16278e9db813
	github.com/gkampitakis/go-snaps v0.4.10
	github.com/go-openapi/jsonreference v0.19.6
	github.com/go-openapi/spec v0.20.4
	github.com/go-openapi/swag v0.21.1
	github.com/golang-jwt/jwt v3.2.1+incompatible
	github.com/golang/protobuf v1.5.4
	github.com/google/uuid v1.6.0
	github.com/hashicorp/go-azure-helpers v0.51.0
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hashicorp/hcl/v2 v2.24.0
	github.com/manicminer/hamilton v0.57.1
	github.com/manicminer/hamilton-autorest v0.3.0
	github.com/microsoftgraph/msgraph-sdk-go v1.81.0
	github.com/microsoftgraph/msgraph-sdk-go-core v1.3.2
	github.com/pkg/errors v0.9.1
	github.com/pulumi/providertest v0.4.1
	github.com/pulumi/pulumi-java/pkg v1.17.0
	github.com/pulumi/pulumi-yaml v1.25.0
	github.com/pulumi/pulumi/pkg/v3 v3.210.0
	github.com/pulumi/pulumi/sdk/v3 v3.210.0
	github.com/schollz/progressbar/v3 v3.5.0
	github.com/segmentio/encoding v0.3.5
	github.com/stretchr/testify v1.11.1
	github.com/tombuildsstuff/giovanni v0.15.1
	github.com/zclconf/go-cty v1.16.3
	golang.org/x/exp v0.0.0-20250718183923-645b1fa84792
	google.golang.org/grpc v1.76.0
	google.golang.org/protobuf v1.36.6
	gopkg.in/yaml.v3 v3.0.1
	pgregory.net/rapid v1.1.0
)

require (
	dario.cat/mergo v1.0.1 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.11.2 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/internal v1.2.0 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.23 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.6.0 // indirect
	github.com/BurntSushi/toml v1.4.0 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/atotto/clipboard v0.1.4 // indirect
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/charmbracelet/bubbles v0.20.0 // indirect
	github.com/charmbracelet/bubbletea v1.3.4 // indirect
	github.com/charmbracelet/lipgloss v1.0.0 // indirect
	github.com/charmbracelet/x/ansi v0.8.0 // indirect
	github.com/charmbracelet/x/term v0.2.1 // indirect
	github.com/cloudflare/circl v1.6.1 // indirect
	github.com/cyphar/filepath-securejoin v0.4.1 // indirect
	github.com/erikgeiser/coninput v0.0.0-20211004153227-1c3628e74d0f // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/gkampitakis/ciinfo v0.3.0 // indirect
	github.com/gkampitakis/go-diff v1.3.2 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang-jwt/jwt/v5 v5.3.0 // indirect
	github.com/iwdgo/sigintwindows v0.2.2 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-localereader v0.0.1 // indirect
	github.com/microsoft/kiota-abstractions-go v1.9.3 // indirect
	github.com/microsoft/kiota-authentication-azure-go v1.3.0 // indirect
	github.com/microsoft/kiota-http-go v1.5.2 // indirect
	github.com/microsoft/kiota-serialization-form-go v1.1.2 // indirect
	github.com/microsoft/kiota-serialization-json-go v1.1.2 // indirect
	github.com/microsoft/kiota-serialization-multipart-go v1.1.2 // indirect
	github.com/microsoft/kiota-serialization-text-go v1.1.2 // indirect
	github.com/muesli/ansi v0.0.0-20230316100256-276c6243b2f6 // indirect
	github.com/muesli/cancelreader v0.2.2 // indirect
	github.com/muesli/termenv v0.16.0 // indirect
	github.com/nxadm/tail v1.4.11 // indirect
	github.com/pgavlin/fx v0.1.6 // indirect
	github.com/pgavlin/fx/v2 v2.0.10 // indirect
	github.com/pjbgf/sha1cd v0.3.2 // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/pulumi/appdash v0.0.0-20231130102222-75f619a67231 // indirect
	github.com/pulumi/esc v0.20.0 // indirect
	github.com/pulumi/inflector v0.2.1 // indirect
	github.com/skeema/knownhosts v1.3.1 // indirect
	github.com/spf13/afero v1.9.5 // indirect
	github.com/std-uritemplate/std-uritemplate/go/v2 v2.0.3 // indirect
	github.com/tidwall/gjson v1.17.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/sjson v1.2.5 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel v1.37.0 // indirect
	go.opentelemetry.io/otel/metric v1.37.0 // indirect
	go.opentelemetry.io/otel/trace v1.37.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250804133106-a7a43d27e69b // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	software.sslmate.com/src/go-pkcs12 v0.2.0 // indirect
)

require (
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest/azure/cli v0.4.6 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/ProtonMail/go-crypto v1.1.6 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/cheggaaa/pb v1.0.29 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/davegardnerisme/deephash v0.0.0-20210406090112-6d072427d830
	github.com/dimchansky/utfbom v1.1.1 // indirect
	github.com/djherbis/times v1.6.0 // indirect
	github.com/edsrzf/mmap-go v1.1.0 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/ettle/strcase v0.1.1 // indirect
	github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 // indirect
	github.com/go-git/go-billy/v5 v5.6.2 // indirect
	github.com/go-git/go-git/v5 v5.16.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.2 // indirect
	github.com/golang/glog v1.2.5 // indirect
	github.com/golang/groupcache v0.0.0-20241129210726-2c02b8208cf8 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.7 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/go-version v1.6.0
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-ps v1.0.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/natefinch/atomic v1.0.1 // indirect
	github.com/opentracing/basictracer-go v1.1.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pgavlin/goldmark v1.1.33-0.20200616210433-b5eb04559386 // indirect
	github.com/pkg/term v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	github.com/sabhiram/go-gitignore v0.0.0-20210923224102-525f6e181f06 // indirect
	github.com/santhosh-tekuri/jsonschema/v5 v5.3.1 // indirect
	github.com/segmentio/asm v1.1.3 // indirect
	github.com/sergi/go-diff v1.4.0 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/cobra v1.10.1 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
	github.com/texttheater/golang-levenshtein v1.0.1 // indirect
	github.com/uber/jaeger-client-go v2.30.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/crypto v0.45.0 // indirect
	golang.org/x/mod v0.29.0 // indirect
	golang.org/x/net v0.47.0 // indirect
	golang.org/x/oauth2 v0.30.0 // indirect
	golang.org/x/sync v0.18.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/term v0.37.0 // indirect
	golang.org/x/text v0.31.0 // indirect
	golang.org/x/tools v0.38.0 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v2 v2.4.0
	lukechampine.com/frand v1.5.1 // indirect
)
