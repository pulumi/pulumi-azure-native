


package v20200815preview

type PublicNetworkAccessType string

const (
	// Allows Azure Arc agents to communicate with Azure Arc services over both public (internet) and private endpoints.
	PublicNetworkAccessTypeEnabled = PublicNetworkAccessType("Enabled")
	// Does not allow Azure Arc agents to communicate with Azure Arc services over public (internet) endpoints. The agents must use the private link.
	PublicNetworkAccessTypeDisabled = PublicNetworkAccessType("Disabled")
)

type StatusLevelTypes string

const (
	StatusLevelTypesInfo    = StatusLevelTypes("Info")
	StatusLevelTypesWarning = StatusLevelTypes("Warning")
	StatusLevelTypesError   = StatusLevelTypes("Error")
)

func init() {
}
