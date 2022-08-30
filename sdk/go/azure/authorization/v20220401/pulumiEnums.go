


package v20220401

type PrincipalType string

const (
	PrincipalTypeUser             = PrincipalType("User")
	PrincipalTypeGroup            = PrincipalType("Group")
	PrincipalTypeServicePrincipal = PrincipalType("ServicePrincipal")
	PrincipalTypeForeignGroup     = PrincipalType("ForeignGroup")
	PrincipalTypeDevice           = PrincipalType("Device")
)

func init() {
}
