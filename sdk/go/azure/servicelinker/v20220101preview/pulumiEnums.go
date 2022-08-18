


package v20220101preview

type AuthType string

const (
	AuthTypeSystemAssignedIdentity      = AuthType("systemAssignedIdentity")
	AuthTypeUserAssignedIdentity        = AuthType("userAssignedIdentity")
	AuthTypeServicePrincipalSecret      = AuthType("servicePrincipalSecret")
	AuthTypeServicePrincipalCertificate = AuthType("servicePrincipalCertificate")
	AuthTypeSecret                      = AuthType("secret")
)

type ClientType string

const (
	ClientTypeNone       = ClientType("none")
	ClientTypeDotnet     = ClientType("dotnet")
	ClientTypeJava       = ClientType("java")
	ClientTypePython     = ClientType("python")
	ClientTypeGo         = ClientType("go")
	ClientTypePhp        = ClientType("php")
	ClientTypeRuby       = ClientType("ruby")
	ClientTypeDjango     = ClientType("django")
	ClientTypeNodejs     = ClientType("nodejs")
	ClientTypeSpringBoot = ClientType("springBoot")
)

type SecretType string

const (
	SecretTypeRawValue                = SecretType("rawValue")
	SecretTypeKeyVaultSecretUri       = SecretType("keyVaultSecretUri")
	SecretTypeKeyVaultSecretReference = SecretType("keyVaultSecretReference")
)

type Type string

const (
	TypeAzureResource            = Type("AzureResource")
	TypeConfluentBootstrapServer = Type("ConfluentBootstrapServer")
	TypeConfluentSchemaRegistry  = Type("ConfluentSchemaRegistry")
)

type VNetSolutionType string

const (
	VNetSolutionTypeServiceEndpoint = VNetSolutionType("serviceEndpoint")
	VNetSolutionTypePrivateLink     = VNetSolutionType("privateLink")
)

func init() {
}
