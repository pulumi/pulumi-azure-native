


package v20210501preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSubscriptionDiagnosticSetting(ctx *pulumi.Context, args *LookupSubscriptionDiagnosticSettingArgs, opts ...pulumi.InvokeOption) (*LookupSubscriptionDiagnosticSettingResult, error) {
	var rv LookupSubscriptionDiagnosticSettingResult
	err := ctx.Invoke("azure-native:insights/v20210501preview:getSubscriptionDiagnosticSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubscriptionDiagnosticSettingArgs struct {
	Name string `pulumi:"name"`
}


type LookupSubscriptionDiagnosticSettingResult struct {
	EventHubAuthorizationRuleId *string                           `pulumi:"eventHubAuthorizationRuleId"`
	EventHubName                *string                           `pulumi:"eventHubName"`
	Id                          string                            `pulumi:"id"`
	Logs                        []SubscriptionLogSettingsResponse `pulumi:"logs"`
	MarketplacePartnerId        *string                           `pulumi:"marketplacePartnerId"`
	Name                        string                            `pulumi:"name"`
	ServiceBusRuleId            *string                           `pulumi:"serviceBusRuleId"`
	StorageAccountId            *string                           `pulumi:"storageAccountId"`
	SystemData                  SystemDataResponse                `pulumi:"systemData"`
	Type                        string                            `pulumi:"type"`
	WorkspaceId                 *string                           `pulumi:"workspaceId"`
}
