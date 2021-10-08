


package labservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLabAccount(ctx *pulumi.Context, args *LookupLabAccountArgs, opts ...pulumi.InvokeOption) (*LookupLabAccountResult, error) {
	var rv LookupLabAccountResult
	err := ctx.Invoke("azure-native:labservices:getLabAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLabAccountArgs struct {
	Expand            *string `pulumi:"expand"`
	LabAccountName    string  `pulumi:"labAccountName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupLabAccountResult struct {
	EnabledRegionSelection *bool                               `pulumi:"enabledRegionSelection"`
	Id                     string                              `pulumi:"id"`
	LatestOperationResult  LatestOperationResultResponse       `pulumi:"latestOperationResult"`
	Location               *string                             `pulumi:"location"`
	Name                   string                              `pulumi:"name"`
	ProvisioningState      *string                             `pulumi:"provisioningState"`
	SizeConfiguration      SizeConfigurationPropertiesResponse `pulumi:"sizeConfiguration"`
	Tags                   map[string]string                   `pulumi:"tags"`
	Type                   string                              `pulumi:"type"`
	UniqueIdentifier       *string                             `pulumi:"uniqueIdentifier"`
}
