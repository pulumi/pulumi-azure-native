


package v20211101preview

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

func init() {
}
