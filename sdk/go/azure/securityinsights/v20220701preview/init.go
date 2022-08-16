


package v20220701preview

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
	case "azure-native:securityinsights/v20220701preview:AADDataConnector":
		r = &AADDataConnector{}
	case "azure-native:securityinsights/v20220701preview:AATPDataConnector":
		r = &AATPDataConnector{}
	case "azure-native:securityinsights/v20220701preview:ASCDataConnector":
		r = &ASCDataConnector{}
	case "azure-native:securityinsights/v20220701preview:Action":
		r = &Action{}
	case "azure-native:securityinsights/v20220701preview:ActivityCustomEntityQuery":
		r = &ActivityCustomEntityQuery{}
	case "azure-native:securityinsights/v20220701preview:AlertRule":
		r = &AlertRule{}
	case "azure-native:securityinsights/v20220701preview:Anomalies":
		r = &Anomalies{}
	case "azure-native:securityinsights/v20220701preview:AnomalySecurityMLAnalyticsSettings":
		r = &AnomalySecurityMLAnalyticsSettings{}
	case "azure-native:securityinsights/v20220701preview:AutomationRule":
		r = &AutomationRule{}
	case "azure-native:securityinsights/v20220701preview:AwsCloudTrailDataConnector":
		r = &AwsCloudTrailDataConnector{}
	case "azure-native:securityinsights/v20220701preview:AwsS3DataConnector":
		r = &AwsS3DataConnector{}
	case "azure-native:securityinsights/v20220701preview:Bookmark":
		r = &Bookmark{}
	case "azure-native:securityinsights/v20220701preview:BookmarkRelation":
		r = &BookmarkRelation{}
	case "azure-native:securityinsights/v20220701preview:CodelessApiPollingDataConnector":
		r = &CodelessApiPollingDataConnector{}
	case "azure-native:securityinsights/v20220701preview:CodelessUiDataConnector":
		r = &CodelessUiDataConnector{}
	case "azure-native:securityinsights/v20220701preview:DataConnector":
		r = &DataConnector{}
	case "azure-native:securityinsights/v20220701preview:Dynamics365DataConnector":
		r = &Dynamics365DataConnector{}
	case "azure-native:securityinsights/v20220701preview:EntityAnalytics":
		r = &EntityAnalytics{}
	case "azure-native:securityinsights/v20220701preview:EntityQuery":
		r = &EntityQuery{}
	case "azure-native:securityinsights/v20220701preview:EyesOn":
		r = &EyesOn{}
	case "azure-native:securityinsights/v20220701preview:FusionAlertRule":
		r = &FusionAlertRule{}
	case "azure-native:securityinsights/v20220701preview:Incident":
		r = &Incident{}
	case "azure-native:securityinsights/v20220701preview:IncidentComment":
		r = &IncidentComment{}
	case "azure-native:securityinsights/v20220701preview:IncidentRelation":
		r = &IncidentRelation{}
	case "azure-native:securityinsights/v20220701preview:IoTDataConnector":
		r = &IoTDataConnector{}
	case "azure-native:securityinsights/v20220701preview:MCASDataConnector":
		r = &MCASDataConnector{}
	case "azure-native:securityinsights/v20220701preview:MDATPDataConnector":
		r = &MDATPDataConnector{}
	case "azure-native:securityinsights/v20220701preview:MLBehaviorAnalyticsAlertRule":
		r = &MLBehaviorAnalyticsAlertRule{}
	case "azure-native:securityinsights/v20220701preview:MSTIDataConnector":
		r = &MSTIDataConnector{}
	case "azure-native:securityinsights/v20220701preview:MTPDataConnector":
		r = &MTPDataConnector{}
	case "azure-native:securityinsights/v20220701preview:Metadata":
		r = &Metadata{}
	case "azure-native:securityinsights/v20220701preview:MicrosoftSecurityIncidentCreationAlertRule":
		r = &MicrosoftSecurityIncidentCreationAlertRule{}
	case "azure-native:securityinsights/v20220701preview:NrtAlertRule":
		r = &NrtAlertRule{}
	case "azure-native:securityinsights/v20220701preview:Office365ProjectDataConnector":
		r = &Office365ProjectDataConnector{}
	case "azure-native:securityinsights/v20220701preview:OfficeATPDataConnector":
		r = &OfficeATPDataConnector{}
	case "azure-native:securityinsights/v20220701preview:OfficeDataConnector":
		r = &OfficeDataConnector{}
	case "azure-native:securityinsights/v20220701preview:OfficeIRMDataConnector":
		r = &OfficeIRMDataConnector{}
	case "azure-native:securityinsights/v20220701preview:OfficePowerBIDataConnector":
		r = &OfficePowerBIDataConnector{}
	case "azure-native:securityinsights/v20220701preview:ProductSetting":
		r = &ProductSetting{}
	case "azure-native:securityinsights/v20220701preview:ScheduledAlertRule":
		r = &ScheduledAlertRule{}
	case "azure-native:securityinsights/v20220701preview:SecurityMLAnalyticsSetting":
		r = &SecurityMLAnalyticsSetting{}
	case "azure-native:securityinsights/v20220701preview:SentinelOnboardingState":
		r = &SentinelOnboardingState{}
	case "azure-native:securityinsights/v20220701preview:SourceControl":
		r = &SourceControl{}
	case "azure-native:securityinsights/v20220701preview:TIDataConnector":
		r = &TIDataConnector{}
	case "azure-native:securityinsights/v20220701preview:ThreatIntelligenceAlertRule":
		r = &ThreatIntelligenceAlertRule{}
	case "azure-native:securityinsights/v20220701preview:ThreatIntelligenceIndicator":
		r = &ThreatIntelligenceIndicator{}
	case "azure-native:securityinsights/v20220701preview:TiTaxiiDataConnector":
		r = &TiTaxiiDataConnector{}
	case "azure-native:securityinsights/v20220701preview:Ueba":
		r = &Ueba{}
	case "azure-native:securityinsights/v20220701preview:Watchlist":
		r = &Watchlist{}
	case "azure-native:securityinsights/v20220701preview:WatchlistItem":
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
		"securityinsights/v20220701preview",
		&module{version},
	)
}
