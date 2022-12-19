


package v20221101preview

type AccessKeyPermissions string

const (
	AccessKeyPermissionsRead   = AccessKeyPermissions("Read")
	AccessKeyPermissionsWrite  = AccessKeyPermissions("Write")
	AccessKeyPermissionsListen = AccessKeyPermissions("Listen")
	AccessKeyPermissionsSend   = AccessKeyPermissions("Send")
	AccessKeyPermissionsManage = AccessKeyPermissions("Manage")
)

type ActionType string

const (
	ActionTypeEnable = ActionType("enable")
	ActionTypeOptOut = ActionType("optOut")
)

type AllowType string

const (
	AllowTypeTrue  = AllowType("true")
	AllowTypeFalse = AllowType("false")
)

type AuthType string

const (
	AuthTypeSystemAssignedIdentity      = AuthType("systemAssignedIdentity")
	AuthTypeUserAssignedIdentity        = AuthType("userAssignedIdentity")
	AuthTypeServicePrincipalSecret      = AuthType("servicePrincipalSecret")
	AuthTypeServicePrincipalCertificate = AuthType("servicePrincipalCertificate")
	AuthTypeSecret                      = AuthType("secret")
	AuthTypeAccessKey                   = AuthType("accessKey")
	AuthTypeUserAccount                 = AuthType("userAccount")
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

type DeleteOrUpdateBehavior string

const (
	DeleteOrUpdateBehaviorDefault       = DeleteOrUpdateBehavior("Default")
	DeleteOrUpdateBehaviorForcedCleanup = DeleteOrUpdateBehavior("ForcedCleanup")
)

type DryrunActionName string

const (
	DryrunActionNameCreateOrUpdate = DryrunActionName("createOrUpdate")
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
	TargetServiceTypeSelfHostedServer         = TargetServiceType("SelfHostedServer")
)

type VNetSolutionType string

const (
	VNetSolutionTypeServiceEndpoint = VNetSolutionType("serviceEndpoint")
	VNetSolutionTypePrivateLink     = VNetSolutionType("privateLink")
)

func init() {
}
