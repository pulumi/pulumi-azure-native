


package v20221101preview

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
	case "azure-native:securityinsights/v20221101preview:AADDataConnector":
		r = &AADDataConnector{}
	case "azure-native:securityinsights/v20221101preview:AATPDataConnector":
		r = &AATPDataConnector{}
	case "azure-native:securityinsights/v20221101preview:ASCDataConnector":
		r = &ASCDataConnector{}
	case "azure-native:securityinsights/v20221101preview:Action":
		r = &Action{}
	case "azure-native:securityinsights/v20221101preview:ActivityCustomEntityQuery":
		r = &ActivityCustomEntityQuery{}
	case "azure-native:securityinsights/v20221101preview:AlertRule":
		r = &AlertRule{}
	case "azure-native:securityinsights/v20221101preview:Anomalies":
		r = &Anomalies{}
	case "azure-native:securityinsights/v20221101preview:AnomalySecurityMLAnalyticsSettings":
		r = &AnomalySecurityMLAnalyticsSettings{}
	case "azure-native:securityinsights/v20221101preview:AutomationRule":
		r = &AutomationRule{}
	case "azure-native:securityinsights/v20221101preview:AwsCloudTrailDataConnector":
		r = &AwsCloudTrailDataConnector{}
	case "azure-native:securityinsights/v20221101preview:AwsS3DataConnector":
		r = &AwsS3DataConnector{}
	case "azure-native:securityinsights/v20221101preview:Bookmark":
		r = &Bookmark{}
	case "azure-native:securityinsights/v20221101preview:BookmarkRelation":
		r = &BookmarkRelation{}
	case "azure-native:securityinsights/v20221101preview:CodelessApiPollingDataConnector":
		r = &CodelessApiPollingDataConnector{}
	case "azure-native:securityinsights/v20221101preview:CodelessUiDataConnector":
		r = &CodelessUiDataConnector{}
	case "azure-native:securityinsights/v20221101preview:DataConnector":
		r = &DataConnector{}
	case "azure-native:securityinsights/v20221101preview:Dynamics365DataConnector":
		r = &Dynamics365DataConnector{}
	case "azure-native:securityinsights/v20221101preview:EntityAnalytics":
		r = &EntityAnalytics{}
	case "azure-native:securityinsights/v20221101preview:EntityQuery":
		r = &EntityQuery{}
	case "azure-native:securityinsights/v20221101preview:EyesOn":
		r = &EyesOn{}
	case "azure-native:securityinsights/v20221101preview:FileImport":
		r = &FileImport{}
	case "azure-native:securityinsights/v20221101preview:FusionAlertRule":
		r = &FusionAlertRule{}
	case "azure-native:securityinsights/v20221101preview:Incident":
		r = &Incident{}
	case "azure-native:securityinsights/v20221101preview:IncidentComment":
		r = &IncidentComment{}
	case "azure-native:securityinsights/v20221101preview:IncidentRelation":
		r = &IncidentRelation{}
	case "azure-native:securityinsights/v20221101preview:IoTDataConnector":
		r = &IoTDataConnector{}
	case "azure-native:securityinsights/v20221101preview:MCASDataConnector":
		r = &MCASDataConnector{}
	case "azure-native:securityinsights/v20221101preview:MDATPDataConnector":
		r = &MDATPDataConnector{}
	case "azure-native:securityinsights/v20221101preview:MLBehaviorAnalyticsAlertRule":
		r = &MLBehaviorAnalyticsAlertRule{}
	case "azure-native:securityinsights/v20221101preview:MSTIDataConnector":
		r = &MSTIDataConnector{}
	case "azure-native:securityinsights/v20221101preview:MTPDataConnector":
		r = &MTPDataConnector{}
	case "azure-native:securityinsights/v20221101preview:Metadata":
		r = &Metadata{}
	case "azure-native:securityinsights/v20221101preview:MicrosoftSecurityIncidentCreationAlertRule":
		r = &MicrosoftSecurityIncidentCreationAlertRule{}
	case "azure-native:securityinsights/v20221101preview:NrtAlertRule":
		r = &NrtAlertRule{}
	case "azure-native:securityinsights/v20221101preview:Office365ProjectDataConnector":
		r = &Office365ProjectDataConnector{}
	case "azure-native:securityinsights/v20221101preview:OfficeATPDataConnector":
		r = &OfficeATPDataConnector{}
	case "azure-native:securityinsights/v20221101preview:OfficeDataConnector":
		r = &OfficeDataConnector{}
	case "azure-native:securityinsights/v20221101preview:OfficeIRMDataConnector":
		r = &OfficeIRMDataConnector{}
	case "azure-native:securityinsights/v20221101preview:OfficePowerBIDataConnector":
		r = &OfficePowerBIDataConnector{}
	case "azure-native:securityinsights/v20221101preview:ProductSetting":
		r = &ProductSetting{}
	case "azure-native:securityinsights/v20221101preview:ScheduledAlertRule":
		r = &ScheduledAlertRule{}
	case "azure-native:securityinsights/v20221101preview:SecurityMLAnalyticsSetting":
		r = &SecurityMLAnalyticsSetting{}
	case "azure-native:securityinsights/v20221101preview:SentinelOnboardingState":
		r = &SentinelOnboardingState{}
	case "azure-native:securityinsights/v20221101preview:SourceControl":
		r = &SourceControl{}
	case "azure-native:securityinsights/v20221101preview:TIDataConnector":
		r = &TIDataConnector{}
	case "azure-native:securityinsights/v20221101preview:ThreatIntelligenceAlertRule":
		r = &ThreatIntelligenceAlertRule{}
	case "azure-native:securityinsights/v20221101preview:ThreatIntelligenceIndicator":
		r = &ThreatIntelligenceIndicator{}
	case "azure-native:securityinsights/v20221101preview:TiTaxiiDataConnector":
		r = &TiTaxiiDataConnector{}
	case "azure-native:securityinsights/v20221101preview:Ueba":
		r = &Ueba{}
	case "azure-native:securityinsights/v20221101preview:Watchlist":
		r = &Watchlist{}
	case "azure-native:securityinsights/v20221101preview:WatchlistItem":
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
		"securityinsights/v20221101preview",
		&module{version},
	)
}
