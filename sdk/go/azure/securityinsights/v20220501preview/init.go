


package v20220501preview

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
	case "azure-native:securityinsights/v20220501preview:AADDataConnector":
		r = &AADDataConnector{}
	case "azure-native:securityinsights/v20220501preview:AATPDataConnector":
		r = &AATPDataConnector{}
	case "azure-native:securityinsights/v20220501preview:ASCDataConnector":
		r = &ASCDataConnector{}
	case "azure-native:securityinsights/v20220501preview:Action":
		r = &Action{}
	case "azure-native:securityinsights/v20220501preview:ActivityCustomEntityQuery":
		r = &ActivityCustomEntityQuery{}
	case "azure-native:securityinsights/v20220501preview:AlertRule":
		r = &AlertRule{}
	case "azure-native:securityinsights/v20220501preview:Anomalies":
		r = &Anomalies{}
	case "azure-native:securityinsights/v20220501preview:AnomalySecurityMLAnalyticsSettings":
		r = &AnomalySecurityMLAnalyticsSettings{}
	case "azure-native:securityinsights/v20220501preview:AutomationRule":
		r = &AutomationRule{}
	case "azure-native:securityinsights/v20220501preview:AwsCloudTrailDataConnector":
		r = &AwsCloudTrailDataConnector{}
	case "azure-native:securityinsights/v20220501preview:AwsS3DataConnector":
		r = &AwsS3DataConnector{}
	case "azure-native:securityinsights/v20220501preview:Bookmark":
		r = &Bookmark{}
	case "azure-native:securityinsights/v20220501preview:BookmarkRelation":
		r = &BookmarkRelation{}
	case "azure-native:securityinsights/v20220501preview:CodelessApiPollingDataConnector":
		r = &CodelessApiPollingDataConnector{}
	case "azure-native:securityinsights/v20220501preview:CodelessUiDataConnector":
		r = &CodelessUiDataConnector{}
	case "azure-native:securityinsights/v20220501preview:DataConnector":
		r = &DataConnector{}
	case "azure-native:securityinsights/v20220501preview:Dynamics365DataConnector":
		r = &Dynamics365DataConnector{}
	case "azure-native:securityinsights/v20220501preview:EntityAnalytics":
		r = &EntityAnalytics{}
	case "azure-native:securityinsights/v20220501preview:EntityQuery":
		r = &EntityQuery{}
	case "azure-native:securityinsights/v20220501preview:EyesOn":
		r = &EyesOn{}
	case "azure-native:securityinsights/v20220501preview:FusionAlertRule":
		r = &FusionAlertRule{}
	case "azure-native:securityinsights/v20220501preview:Incident":
		r = &Incident{}
	case "azure-native:securityinsights/v20220501preview:IncidentComment":
		r = &IncidentComment{}
	case "azure-native:securityinsights/v20220501preview:IncidentRelation":
		r = &IncidentRelation{}
	case "azure-native:securityinsights/v20220501preview:IoTDataConnector":
		r = &IoTDataConnector{}
	case "azure-native:securityinsights/v20220501preview:MCASDataConnector":
		r = &MCASDataConnector{}
	case "azure-native:securityinsights/v20220501preview:MDATPDataConnector":
		r = &MDATPDataConnector{}
	case "azure-native:securityinsights/v20220501preview:MLBehaviorAnalyticsAlertRule":
		r = &MLBehaviorAnalyticsAlertRule{}
	case "azure-native:securityinsights/v20220501preview:MSTIDataConnector":
		r = &MSTIDataConnector{}
	case "azure-native:securityinsights/v20220501preview:MTPDataConnector":
		r = &MTPDataConnector{}
	case "azure-native:securityinsights/v20220501preview:Metadata":
		r = &Metadata{}
	case "azure-native:securityinsights/v20220501preview:MicrosoftSecurityIncidentCreationAlertRule":
		r = &MicrosoftSecurityIncidentCreationAlertRule{}
	case "azure-native:securityinsights/v20220501preview:NrtAlertRule":
		r = &NrtAlertRule{}
	case "azure-native:securityinsights/v20220501preview:Office365ProjectDataConnector":
		r = &Office365ProjectDataConnector{}
	case "azure-native:securityinsights/v20220501preview:OfficeATPDataConnector":
		r = &OfficeATPDataConnector{}
	case "azure-native:securityinsights/v20220501preview:OfficeDataConnector":
		r = &OfficeDataConnector{}
	case "azure-native:securityinsights/v20220501preview:OfficeIRMDataConnector":
		r = &OfficeIRMDataConnector{}
	case "azure-native:securityinsights/v20220501preview:OfficePowerBIDataConnector":
		r = &OfficePowerBIDataConnector{}
	case "azure-native:securityinsights/v20220501preview:ProductSetting":
		r = &ProductSetting{}
	case "azure-native:securityinsights/v20220501preview:ScheduledAlertRule":
		r = &ScheduledAlertRule{}
	case "azure-native:securityinsights/v20220501preview:SecurityMLAnalyticsSetting":
		r = &SecurityMLAnalyticsSetting{}
	case "azure-native:securityinsights/v20220501preview:SentinelOnboardingState":
		r = &SentinelOnboardingState{}
	case "azure-native:securityinsights/v20220501preview:SourceControl":
		r = &SourceControl{}
	case "azure-native:securityinsights/v20220501preview:TIDataConnector":
		r = &TIDataConnector{}
	case "azure-native:securityinsights/v20220501preview:ThreatIntelligenceAlertRule":
		r = &ThreatIntelligenceAlertRule{}
	case "azure-native:securityinsights/v20220501preview:ThreatIntelligenceIndicator":
		r = &ThreatIntelligenceIndicator{}
	case "azure-native:securityinsights/v20220501preview:TiTaxiiDataConnector":
		r = &TiTaxiiDataConnector{}
	case "azure-native:securityinsights/v20220501preview:Ueba":
		r = &Ueba{}
	case "azure-native:securityinsights/v20220501preview:Watchlist":
		r = &Watchlist{}
	case "azure-native:securityinsights/v20220501preview:WatchlistItem":
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
		"securityinsights/v20220501preview",
		&module{version},
	)
}
