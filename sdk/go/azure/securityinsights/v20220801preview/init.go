


package v20220801preview

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
	case "azure-native:securityinsights/v20220801preview:AADDataConnector":
		r = &AADDataConnector{}
	case "azure-native:securityinsights/v20220801preview:AATPDataConnector":
		r = &AATPDataConnector{}
	case "azure-native:securityinsights/v20220801preview:ASCDataConnector":
		r = &ASCDataConnector{}
	case "azure-native:securityinsights/v20220801preview:Action":
		r = &Action{}
	case "azure-native:securityinsights/v20220801preview:ActivityCustomEntityQuery":
		r = &ActivityCustomEntityQuery{}
	case "azure-native:securityinsights/v20220801preview:AlertRule":
		r = &AlertRule{}
	case "azure-native:securityinsights/v20220801preview:Anomalies":
		r = &Anomalies{}
	case "azure-native:securityinsights/v20220801preview:AnomalySecurityMLAnalyticsSettings":
		r = &AnomalySecurityMLAnalyticsSettings{}
	case "azure-native:securityinsights/v20220801preview:AutomationRule":
		r = &AutomationRule{}
	case "azure-native:securityinsights/v20220801preview:AwsCloudTrailDataConnector":
		r = &AwsCloudTrailDataConnector{}
	case "azure-native:securityinsights/v20220801preview:AwsS3DataConnector":
		r = &AwsS3DataConnector{}
	case "azure-native:securityinsights/v20220801preview:Bookmark":
		r = &Bookmark{}
	case "azure-native:securityinsights/v20220801preview:BookmarkRelation":
		r = &BookmarkRelation{}
	case "azure-native:securityinsights/v20220801preview:CodelessApiPollingDataConnector":
		r = &CodelessApiPollingDataConnector{}
	case "azure-native:securityinsights/v20220801preview:CodelessUiDataConnector":
		r = &CodelessUiDataConnector{}
	case "azure-native:securityinsights/v20220801preview:DataConnector":
		r = &DataConnector{}
	case "azure-native:securityinsights/v20220801preview:Dynamics365DataConnector":
		r = &Dynamics365DataConnector{}
	case "azure-native:securityinsights/v20220801preview:EntityAnalytics":
		r = &EntityAnalytics{}
	case "azure-native:securityinsights/v20220801preview:EntityQuery":
		r = &EntityQuery{}
	case "azure-native:securityinsights/v20220801preview:EyesOn":
		r = &EyesOn{}
	case "azure-native:securityinsights/v20220801preview:FileImport":
		r = &FileImport{}
	case "azure-native:securityinsights/v20220801preview:FusionAlertRule":
		r = &FusionAlertRule{}
	case "azure-native:securityinsights/v20220801preview:Incident":
		r = &Incident{}
	case "azure-native:securityinsights/v20220801preview:IncidentComment":
		r = &IncidentComment{}
	case "azure-native:securityinsights/v20220801preview:IncidentRelation":
		r = &IncidentRelation{}
	case "azure-native:securityinsights/v20220801preview:IoTDataConnector":
		r = &IoTDataConnector{}
	case "azure-native:securityinsights/v20220801preview:MCASDataConnector":
		r = &MCASDataConnector{}
	case "azure-native:securityinsights/v20220801preview:MDATPDataConnector":
		r = &MDATPDataConnector{}
	case "azure-native:securityinsights/v20220801preview:MLBehaviorAnalyticsAlertRule":
		r = &MLBehaviorAnalyticsAlertRule{}
	case "azure-native:securityinsights/v20220801preview:MSTIDataConnector":
		r = &MSTIDataConnector{}
	case "azure-native:securityinsights/v20220801preview:MTPDataConnector":
		r = &MTPDataConnector{}
	case "azure-native:securityinsights/v20220801preview:Metadata":
		r = &Metadata{}
	case "azure-native:securityinsights/v20220801preview:MicrosoftSecurityIncidentCreationAlertRule":
		r = &MicrosoftSecurityIncidentCreationAlertRule{}
	case "azure-native:securityinsights/v20220801preview:NrtAlertRule":
		r = &NrtAlertRule{}
	case "azure-native:securityinsights/v20220801preview:Office365ProjectDataConnector":
		r = &Office365ProjectDataConnector{}
	case "azure-native:securityinsights/v20220801preview:OfficeATPDataConnector":
		r = &OfficeATPDataConnector{}
	case "azure-native:securityinsights/v20220801preview:OfficeDataConnector":
		r = &OfficeDataConnector{}
	case "azure-native:securityinsights/v20220801preview:OfficeIRMDataConnector":
		r = &OfficeIRMDataConnector{}
	case "azure-native:securityinsights/v20220801preview:OfficePowerBIDataConnector":
		r = &OfficePowerBIDataConnector{}
	case "azure-native:securityinsights/v20220801preview:ProductSetting":
		r = &ProductSetting{}
	case "azure-native:securityinsights/v20220801preview:ScheduledAlertRule":
		r = &ScheduledAlertRule{}
	case "azure-native:securityinsights/v20220801preview:SecurityMLAnalyticsSetting":
		r = &SecurityMLAnalyticsSetting{}
	case "azure-native:securityinsights/v20220801preview:SentinelOnboardingState":
		r = &SentinelOnboardingState{}
	case "azure-native:securityinsights/v20220801preview:SourceControl":
		r = &SourceControl{}
	case "azure-native:securityinsights/v20220801preview:TIDataConnector":
		r = &TIDataConnector{}
	case "azure-native:securityinsights/v20220801preview:ThreatIntelligenceAlertRule":
		r = &ThreatIntelligenceAlertRule{}
	case "azure-native:securityinsights/v20220801preview:ThreatIntelligenceIndicator":
		r = &ThreatIntelligenceIndicator{}
	case "azure-native:securityinsights/v20220801preview:TiTaxiiDataConnector":
		r = &TiTaxiiDataConnector{}
	case "azure-native:securityinsights/v20220801preview:Ueba":
		r = &Ueba{}
	case "azure-native:securityinsights/v20220801preview:Watchlist":
		r = &Watchlist{}
	case "azure-native:securityinsights/v20220801preview:WatchlistItem":
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
		"securityinsights/v20220801preview",
		&module{version},
	)
}
