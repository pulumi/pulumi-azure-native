


package v20190101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupProductSetting(ctx *pulumi.Context, args *LookupProductSettingArgs, opts ...pulumi.InvokeOption) (*LookupProductSettingResult, error) {
	var rv LookupProductSettingResult
	err := ctx.Invoke("azure-native:securityinsights/v20190101preview:getProductSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProductSettingArgs struct {
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	SettingsName                        string `pulumi:"settingsName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupProductSettingResult struct {
	Etag *string `pulumi:"etag"`
	Id   string  `pulumi:"id"`
	Kind string  `pulumi:"kind"`
	Name string  `pulumi:"name"`
	Type string  `pulumi:"type"`
}
