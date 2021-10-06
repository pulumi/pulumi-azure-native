


package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDefenderSetting(ctx *pulumi.Context, args *LookupDefenderSettingArgs, opts ...pulumi.InvokeOption) (*LookupDefenderSettingResult, error) {
	var rv LookupDefenderSettingResult
	err := ctx.Invoke("azure-native:iotsecurity/v20210201preview:getDefenderSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDefenderSettingArgs struct {
}


type LookupDefenderSettingResult struct {
	DeviceQuota                  int                                              `pulumi:"deviceQuota"`
	EvaluationEndTime            string                                           `pulumi:"evaluationEndTime"`
	Id                           string                                           `pulumi:"id"`
	MdeIntegration               DefenderSettingsPropertiesResponseMdeIntegration `pulumi:"mdeIntegration"`
	Name                         string                                           `pulumi:"name"`
	OnboardingKind               string                                           `pulumi:"onboardingKind"`
	SentinelWorkspaceResourceIds []string                                         `pulumi:"sentinelWorkspaceResourceIds"`
	Type                         string                                           `pulumi:"type"`
}
