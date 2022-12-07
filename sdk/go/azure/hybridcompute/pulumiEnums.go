


package hybridcompute

type PublicNetworkAccessType string

const (
	// Allows Azure Arc agents to communicate with Azure Arc services over both public (internet) and private endpoints.
	PublicNetworkAccessTypeEnabled = PublicNetworkAccessType("Enabled")
	// Does not allow Azure Arc agents to communicate with Azure Arc services over public (internet) endpoints. The agents must use the private link.
	PublicNetworkAccessTypeDisabled = PublicNetworkAccessType("Disabled")
)

func init() {
}
