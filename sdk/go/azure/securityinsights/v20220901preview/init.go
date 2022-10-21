


package v20220901preview

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
	case "azure-native:securityinsights/v20220901preview:AADDataConnector":
		r = &AADDataConnector{}
	case "azure-native:securityinsights/v20220901preview:AATPDataConnector":
		r = &AATPDataConnector{}
	case "azure-native:securityinsights/v20220901preview:ASCDataConnector":
		r = &ASCDataConnector{}
	case "azure-native:securityinsights/v20220901preview:Action":
		r = &Action{}
	case "azure-native:securityinsights/v20220901preview:ActivityCustomEntityQuery":
		r = &ActivityCustomEntityQuery{}
	case "azure-native:securityinsights/v20220901preview:AlertRule":
		r = &AlertRule{}
	case "azure-native:securityinsights/v20220901preview:Anomalies":
		r = &Anomalies{}
	case "azure-native:securityinsights/v20220901preview:AnomalySecurityMLAnalyticsSettings":
		r = &AnomalySecurityMLAnalyticsSettings{}
	case "azure-native:securityinsights/v20220901preview:AutomationRule":
		r = &AutomationRule{}
	case "azure-native:securityinsights/v20220901preview:AwsCloudTrailDataConnector":
		r = &AwsCloudTrailDataConnector{}
	case "azure-native:securityinsights/v20220901preview:AwsS3DataConnector":
		r = &AwsS3DataConnector{}
	case "azure-native:securityinsights/v20220901preview:Bookmark":
		r = &Bookmark{}
	case "azure-native:securityinsights/v20220901preview:BookmarkRelation":
		r = &BookmarkRelation{}
	case "azure-native:securityinsights/v20220901preview:CodelessApiPollingDataConnector":
		r = &CodelessApiPollingDataConnector{}
	case "azure-native:securityinsights/v20220901preview:CodelessUiDataConnector":
		r = &CodelessUiDataConnector{}
	case "azure-native:securityinsights/v20220901preview:DataConnector":
		r = &DataConnector{}
	case "azure-native:securityinsights/v20220901preview:Dynamics365DataConnector":
		r = &Dynamics365DataConnector{}
	case "azure-native:securityinsights/v20220901preview:EntityAnalytics":
		r = &EntityAnalytics{}
	case "azure-native:securityinsights/v20220901preview:EntityQuery":
		r = &EntityQuery{}
	case "azure-native:securityinsights/v20220901preview:EyesOn":
		r = &EyesOn{}
	case "azure-native:securityinsights/v20220901preview:FileImport":
		r = &FileImport{}
	case "azure-native:securityinsights/v20220901preview:FusionAlertRule":
		r = &FusionAlertRule{}
	case "azure-native:securityinsights/v20220901preview:Incident":
		r = &Incident{}
	case "azure-native:securityinsights/v20220901preview:IncidentComment":
		r = &IncidentComment{}
	case "azure-native:securityinsights/v20220901preview:IncidentRelation":
		r = &IncidentRelation{}
	case "azure-native:securityinsights/v20220901preview:IoTDataConnector":
		r = &IoTDataConnector{}
	case "azure-native:securityinsights/v20220901preview:MCASDataConnector":
		r = &MCASDataConnector{}
	case "azure-native:securityinsights/v20220901preview:MDATPDataConnector":
		r = &MDATPDataConnector{}
	case "azure-native:securityinsights/v20220901preview:MLBehaviorAnalyticsAlertRule":
		r = &MLBehaviorAnalyticsAlertRule{}
	case "azure-native:securityinsights/v20220901preview:MSTIDataConnector":
		r = &MSTIDataConnector{}
	case "azure-native:securityinsights/v20220901preview:MTPDataConnector":
		r = &MTPDataConnector{}
	case "azure-native:securityinsights/v20220901preview:Metadata":
		r = &Metadata{}
	case "azure-native:securityinsights/v20220901preview:MicrosoftSecurityIncidentCreationAlertRule":
		r = &MicrosoftSecurityIncidentCreationAlertRule{}
	case "azure-native:securityinsights/v20220901preview:NrtAlertRule":
		r = &NrtAlertRule{}
	case "azure-native:securityinsights/v20220901preview:Office365ProjectDataConnector":
		r = &Office365ProjectDataConnector{}
	case "azure-native:securityinsights/v20220901preview:OfficeATPDataConnector":
		r = &OfficeATPDataConnector{}
	case "azure-native:securityinsights/v20220901preview:OfficeDataConnector":
		r = &OfficeDataConnector{}
	case "azure-native:securityinsights/v20220901preview:OfficeIRMDataConnector":
		r = &OfficeIRMDataConnector{}
	case "azure-native:securityinsights/v20220901preview:OfficePowerBIDataConnector":
		r = &OfficePowerBIDataConnector{}
	case "azure-native:securityinsights/v20220901preview:ProductSetting":
		r = &ProductSetting{}
	case "azure-native:securityinsights/v20220901preview:ScheduledAlertRule":
		r = &ScheduledAlertRule{}
	case "azure-native:securityinsights/v20220901preview:SecurityMLAnalyticsSetting":
		r = &SecurityMLAnalyticsSetting{}
	case "azure-native:securityinsights/v20220901preview:SentinelOnboardingState":
		r = &SentinelOnboardingState{}
	case "azure-native:securityinsights/v20220901preview:SourceControl":
		r = &SourceControl{}
	case "azure-native:securityinsights/v20220901preview:TIDataConnector":
		r = &TIDataConnector{}
	case "azure-native:securityinsights/v20220901preview:ThreatIntelligenceAlertRule":
		r = &ThreatIntelligenceAlertRule{}
	case "azure-native:securityinsights/v20220901preview:ThreatIntelligenceIndicator":
		r = &ThreatIntelligenceIndicator{}
	case "azure-native:securityinsights/v20220901preview:TiTaxiiDataConnector":
		r = &TiTaxiiDataConnector{}
	case "azure-native:securityinsights/v20220901preview:Ueba":
		r = &Ueba{}
	case "azure-native:securityinsights/v20220901preview:Watchlist":
		r = &Watchlist{}
	case "azure-native:securityinsights/v20220901preview:WatchlistItem":
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
		"securityinsights/v20220901preview",
		&module{version},
	)
}
