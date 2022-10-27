


package v20220501

type AuthType string

const (
	AuthTypeSystemAssignedIdentity      = AuthType("systemAssignedIdentity")
	AuthTypeUserAssignedIdentity        = AuthType("userAssignedIdentity")
	AuthTypeServicePrincipalSecret      = AuthType("servicePrincipalSecret")
	AuthTypeServicePrincipalCertificate = AuthType("servicePrincipalCertificate")
	AuthTypeSecret                      = AuthType("secret")
)

type AzureResourceType string

const (
	AzureResourceTypeKeyVault = AzureResourceType("KeyVault")
)

type ClientType string

const (
	ClientTypeNone              = ClientType("none")
	ClientTypeDotnet            = ClientType("dotnet")
	ClientTypeJava              = ClientType("java")
	ClientTypePython            = ClientType("python")
	ClientTypeGo                = ClientType("go")
	ClientTypePhp               = ClientType("php")
	ClientTypeRuby              = ClientType("ruby")
	ClientTypeDjango            = ClientType("django")
	ClientTypeNodejs            = ClientType("nodejs")
	ClientTypeSpringBoot        = ClientType("springBoot")
	ClientType_Kafka_SpringBoot = ClientType("kafka-springBoot")
)

type SecretType string

const (
	SecretTypeRawValue                = SecretType("rawValue")
	SecretTypeKeyVaultSecretUri       = SecretType("keyVaultSecretUri")
	SecretTypeKeyVaultSecretReference = SecretType("keyVaultSecretReference")
)

type TargetServiceType string

const (
	TargetServiceTypeAzureResource            = TargetServiceType("AzureResource")
	TargetServiceTypeConfluentBootstrapServer = TargetServiceType("ConfluentBootstrapServer")
	TargetServiceTypeConfluentSchemaRegistry  = TargetServiceType("ConfluentSchemaRegistry")
)

type VNetSolutionType string

const (
	VNetSolutionTypeServiceEndpoint = VNetSolutionType("serviceEndpoint")
	VNetSolutionTypePrivateLink     = VNetSolutionType("privateLink")
)

func init() {
}
