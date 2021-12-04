


package insights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDiagnosticSetting(ctx *pulumi.Context, args *LookupDiagnosticSettingArgs, opts ...pulumi.InvokeOption) (*LookupDiagnosticSettingResult, error) {
	var rv LookupDiagnosticSettingResult
	err := ctx.Invoke("azure-native:insights:getDiagnosticSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiagnosticSettingArgs struct {
	Name        string `pulumi:"name"`
	ResourceUri string `pulumi:"resourceUri"`
}


type LookupDiagnosticSettingResult struct {
	EventHubAuthorizationRuleId *string                  `pulumi:"eventHubAuthorizationRuleId"`
	EventHubName                *string                  `pulumi:"eventHubName"`
	Id                          string                   `pulumi:"id"`
	LogAnalyticsDestinationType *string                  `pulumi:"logAnalyticsDestinationType"`
	Logs                        []LogSettingsResponse    `pulumi:"logs"`
	Metrics                     []MetricSettingsResponse `pulumi:"metrics"`
	Name                        string                   `pulumi:"name"`
	ServiceBusRuleId            *string                  `pulumi:"serviceBusRuleId"`
	StorageAccountId            *string                  `pulumi:"storageAccountId"`
	Type                        string                   `pulumi:"type"`
	WorkspaceId                 *string                  `pulumi:"workspaceId"`
}
