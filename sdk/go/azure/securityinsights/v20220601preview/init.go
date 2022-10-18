


package v20220601preview

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
	case "azure-native:securityinsights/v20220601preview:AADDataConnector":
		r = &AADDataConnector{}
	case "azure-native:securityinsights/v20220601preview:AATPDataConnector":
		r = &AATPDataConnector{}
	case "azure-native:securityinsights/v20220601preview:ASCDataConnector":
		r = &ASCDataConnector{}
	case "azure-native:securityinsights/v20220601preview:Action":
		r = &Action{}
	case "azure-native:securityinsights/v20220601preview:ActivityCustomEntityQuery":
		r = &ActivityCustomEntityQuery{}
	case "azure-native:securityinsights/v20220601preview:AlertRule":
		r = &AlertRule{}
	case "azure-native:securityinsights/v20220601preview:Anomalies":
		r = &Anomalies{}
	case "azure-native:securityinsights/v20220601preview:AnomalySecurityMLAnalyticsSettings":
		r = &AnomalySecurityMLAnalyticsSettings{}
	case "azure-native:securityinsights/v20220601preview:AutomationRule":
		r = &AutomationRule{}
	case "azure-native:securityinsights/v20220601preview:AwsCloudTrailDataConnector":
		r = &AwsCloudTrailDataConnector{}
	case "azure-native:securityinsights/v20220601preview:AwsS3DataConnector":
		r = &AwsS3DataConnector{}
	case "azure-native:securityinsights/v20220601preview:Bookmark":
		r = &Bookmark{}
	case "azure-native:securityinsights/v20220601preview:BookmarkRelation":
		r = &BookmarkRelation{}
	case "azure-native:securityinsights/v20220601preview:CodelessApiPollingDataConnector":
		r = &CodelessApiPollingDataConnector{}
	case "azure-native:securityinsights/v20220601preview:CodelessUiDataConnector":
		r = &CodelessUiDataConnector{}
	case "azure-native:securityinsights/v20220601preview:DataConnector":
		r = &DataConnector{}
	case "azure-native:securityinsights/v20220601preview:Dynamics365DataConnector":
		r = &Dynamics365DataConnector{}
	case "azure-native:securityinsights/v20220601preview:EntityAnalytics":
		r = &EntityAnalytics{}
	case "azure-native:securityinsights/v20220601preview:EntityQuery":
		r = &EntityQuery{}
	case "azure-native:securityinsights/v20220601preview:EyesOn":
		r = &EyesOn{}
	case "azure-native:securityinsights/v20220601preview:FusionAlertRule":
		r = &FusionAlertRule{}
	case "azure-native:securityinsights/v20220601preview:Incident":
		r = &Incident{}
	case "azure-native:securityinsights/v20220601preview:IncidentComment":
		r = &IncidentComment{}
	case "azure-native:securityinsights/v20220601preview:IncidentRelation":
		r = &IncidentRelation{}
	case "azure-native:securityinsights/v20220601preview:IoTDataConnector":
		r = &IoTDataConnector{}
	case "azure-native:securityinsights/v20220601preview:MCASDataConnector":
		r = &MCASDataConnector{}
	case "azure-native:securityinsights/v20220601preview:MDATPDataConnector":
		r = &MDATPDataConnector{}
	case "azure-native:securityinsights/v20220601preview:MLBehaviorAnalyticsAlertRule":
		r = &MLBehaviorAnalyticsAlertRule{}
	case "azure-native:securityinsights/v20220601preview:MSTIDataConnector":
		r = &MSTIDataConnector{}
	case "azure-native:securityinsights/v20220601preview:MTPDataConnector":
		r = &MTPDataConnector{}
	case "azure-native:securityinsights/v20220601preview:Metadata":
		r = &Metadata{}
	case "azure-native:securityinsights/v20220601preview:MicrosoftSecurityIncidentCreationAlertRule":
		r = &MicrosoftSecurityIncidentCreationAlertRule{}
	case "azure-native:securityinsights/v20220601preview:NrtAlertRule":
		r = &NrtAlertRule{}
	case "azure-native:securityinsights/v20220601preview:Office365ProjectDataConnector":
		r = &Office365ProjectDataConnector{}
	case "azure-native:securityinsights/v20220601preview:OfficeATPDataConnector":
		r = &OfficeATPDataConnector{}
	case "azure-native:securityinsights/v20220601preview:OfficeDataConnector":
		r = &OfficeDataConnector{}
	case "azure-native:securityinsights/v20220601preview:OfficeIRMDataConnector":
		r = &OfficeIRMDataConnector{}
	case "azure-native:securityinsights/v20220601preview:OfficePowerBIDataConnector":
		r = &OfficePowerBIDataConnector{}
	case "azure-native:securityinsights/v20220601preview:ProductSetting":
		r = &ProductSetting{}
	case "azure-native:securityinsights/v20220601preview:ScheduledAlertRule":
		r = &ScheduledAlertRule{}
	case "azure-native:securityinsights/v20220601preview:SecurityMLAnalyticsSetting":
		r = &SecurityMLAnalyticsSetting{}
	case "azure-native:securityinsights/v20220601preview:SentinelOnboardingState":
		r = &SentinelOnboardingState{}
	case "azure-native:securityinsights/v20220601preview:SourceControl":
		r = &SourceControl{}
	case "azure-native:securityinsights/v20220601preview:TIDataConnector":
		r = &TIDataConnector{}
	case "azure-native:securityinsights/v20220601preview:ThreatIntelligenceAlertRule":
		r = &ThreatIntelligenceAlertRule{}
	case "azure-native:securityinsights/v20220601preview:ThreatIntelligenceIndicator":
		r = &ThreatIntelligenceIndicator{}
	case "azure-native:securityinsights/v20220601preview:TiTaxiiDataConnector":
		r = &TiTaxiiDataConnector{}
	case "azure-native:securityinsights/v20220601preview:Ueba":
		r = &Ueba{}
	case "azure-native:securityinsights/v20220601preview:Watchlist":
		r = &Watchlist{}
	case "azure-native:securityinsights/v20220601preview:WatchlistItem":
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
		"securityinsights/v20220601preview",
		&module{version},
	)
}
