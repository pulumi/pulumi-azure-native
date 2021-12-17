


package v20200202preview

type ApplicationType string

const (
	ApplicationTypeWeb   = ApplicationType("web")
	ApplicationTypeOther = ApplicationType("other")
)

type FlowType string

const (
	FlowTypeBluefield = FlowType("Bluefield")
)

type IngestionMode string

const (
	IngestionModeApplicationInsights                       = IngestionMode("ApplicationInsights")
	IngestionModeApplicationInsightsWithDiagnosticSettings = IngestionMode("ApplicationInsightsWithDiagnosticSettings")
	IngestionModeLogAnalytics                              = IngestionMode("LogAnalytics")
)

type PublicNetworkAccessType string

const (
	// Enables connectivity to Application Insights through public DNS.
	PublicNetworkAccessTypeEnabled = PublicNetworkAccessType("Enabled")
	// Disables public connectivity to Application Insights through public DNS.
	PublicNetworkAccessTypeDisabled = PublicNetworkAccessType("Disabled")
)

type RequestSource string

const (
	RequestSourceRest = RequestSource("rest")
)

func init() {
}
