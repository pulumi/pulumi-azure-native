


package v20221101

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure-native:securityinsights/v20221101:AADDataConnector":
		r = &AADDataConnector{}
	case "azure-native:securityinsights/v20221101:AATPDataConnector":
		r = &AATPDataConnector{}
	case "azure-native:securityinsights/v20221101:ASCDataConnector":
		r = &ASCDataConnector{}
	case "azure-native:securityinsights/v20221101:Action":
		r = &Action{}
	case "azure-native:securityinsights/v20221101:AlertRule":
		r = &AlertRule{}
	case "azure-native:securityinsights/v20221101:AnomalySecurityMLAnalyticsSettings":
		r = &AnomalySecurityMLAnalyticsSettings{}
	case "azure-native:securityinsights/v20221101:AutomationRule":
		r = &AutomationRule{}
	case "azure-native:securityinsights/v20221101:AwsCloudTrailDataConnector":
		r = &AwsCloudTrailDataConnector{}
	case "azure-native:securityinsights/v20221101:Bookmark":
		r = &Bookmark{}
	case "azure-native:securityinsights/v20221101:DataConnector":
		r = &DataConnector{}
	case "azure-native:securityinsights/v20221101:FusionAlertRule":
		r = &FusionAlertRule{}
	case "azure-native:securityinsights/v20221101:Incident":
		r = &Incident{}
	case "azure-native:securityinsights/v20221101:IncidentComment":
		r = &IncidentComment{}
	case "azure-native:securityinsights/v20221101:IncidentRelation":
		r = &IncidentRelation{}
	case "azure-native:securityinsights/v20221101:MCASDataConnector":
		r = &MCASDataConnector{}
	case "azure-native:securityinsights/v20221101:MDATPDataConnector":
		r = &MDATPDataConnector{}
	case "azure-native:securityinsights/v20221101:MicrosoftSecurityIncidentCreationAlertRule":
		r = &MicrosoftSecurityIncidentCreationAlertRule{}
	case "azure-native:securityinsights/v20221101:OfficeDataConnector":
		r = &OfficeDataConnector{}
	case "azure-native:securityinsights/v20221101:ScheduledAlertRule":
		r = &ScheduledAlertRule{}
	case "azure-native:securityinsights/v20221101:SecurityMLAnalyticsSetting":
		r = &SecurityMLAnalyticsSetting{}
	case "azure-native:securityinsights/v20221101:SentinelOnboardingState":
		r = &SentinelOnboardingState{}
	case "azure-native:securityinsights/v20221101:TIDataConnector":
		r = &TIDataConnector{}
	case "azure-native:securityinsights/v20221101:ThreatIntelligenceIndicator":
		r = &ThreatIntelligenceIndicator{}
	case "azure-native:securityinsights/v20221101:Watchlist":
		r = &Watchlist{}
	case "azure-native:securityinsights/v20221101:WatchlistItem":
		r = &WatchlistItem{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"securityinsights/v20221101",
		&module{version},
	)
}
