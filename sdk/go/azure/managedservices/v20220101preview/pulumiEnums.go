


package v20220101preview

type MultiFactorAuthProvider string

const (
	MultiFactorAuthProviderAzure = MultiFactorAuthProvider("Azure")
	MultiFactorAuthProviderNone  = MultiFactorAuthProvider("None")
)

func init() {
}
