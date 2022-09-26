


package v20200801preview

type PrincipalType string

const (
	PrincipalTypeUser             = PrincipalType("User")
	PrincipalTypeGroup            = PrincipalType("Group")
	PrincipalTypeServicePrincipal = PrincipalType("ServicePrincipal")
	PrincipalTypeForeignGroup     = PrincipalType("ForeignGroup")
)

func init() {
}
