


package v20210701preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplicationType(ctx *pulumi.Context, args *LookupApplicationTypeArgs, opts ...pulumi.InvokeOption) (*LookupApplicationTypeResult, error) {
	var rv LookupApplicationTypeResult
	err := ctx.Invoke("azure-native:servicefabric/v20210701preview:getApplicationType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationTypeArgs struct {
	ApplicationTypeName string `pulumi:"applicationTypeName"`
	ClusterName         string `pulumi:"clusterName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupApplicationTypeResult struct {
	Id                string             `pulumi:"id"`
	Location          *string            `pulumi:"location"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Tags              map[string]string  `pulumi:"tags"`
	Type              string             `pulumi:"type"`
}
