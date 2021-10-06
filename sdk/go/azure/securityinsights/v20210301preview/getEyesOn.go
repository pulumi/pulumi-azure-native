


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEyesOn(ctx *pulumi.Context, args *LookupEyesOnArgs, opts ...pulumi.InvokeOption) (*LookupEyesOnResult, error) {
	var rv LookupEyesOnResult
	err := ctx.Invoke("azure-native:securityinsights/v20210301preview:getEyesOn", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEyesOnArgs struct {
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	SettingsName                        string `pulumi:"settingsName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupEyesOnResult struct {
	Etag       *string            `pulumi:"etag"`
	Id         string             `pulumi:"id"`
	IsEnabled  bool               `pulumi:"isEnabled"`
	Kind       string             `pulumi:"kind"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}
