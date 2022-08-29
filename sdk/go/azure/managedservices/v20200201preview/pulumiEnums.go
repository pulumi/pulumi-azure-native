


package v20200201preview

type MultiFactorAuthProvider string

const (
	MultiFactorAuthProviderAzure = MultiFactorAuthProvider("Azure")
	MultiFactorAuthProviderNone  = MultiFactorAuthProvider("None")
)

func init() {
}
