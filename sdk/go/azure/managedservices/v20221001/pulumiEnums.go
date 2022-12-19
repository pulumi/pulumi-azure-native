


package v20221001

type MultiFactorAuthProvider string

const (
	MultiFactorAuthProviderAzure = MultiFactorAuthProvider("Azure")
	MultiFactorAuthProviderNone  = MultiFactorAuthProvider("None")
)

func init() {
}
