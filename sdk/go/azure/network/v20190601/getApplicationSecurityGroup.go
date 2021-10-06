


package v20190601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplicationSecurityGroup(ctx *pulumi.Context, args *LookupApplicationSecurityGroupArgs, opts ...pulumi.InvokeOption) (*LookupApplicationSecurityGroupResult, error) {
	var rv LookupApplicationSecurityGroupResult
	err := ctx.Invoke("azure-native:network/v20190601:getApplicationSecurityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationSecurityGroupArgs struct {
	ApplicationSecurityGroupName string `pulumi:"applicationSecurityGroupName"`
	ResourceGroupName            string `pulumi:"resourceGroupName"`
}


type LookupApplicationSecurityGroupResult struct {
	Etag              string            `pulumi:"etag"`
	Id                *string           `pulumi:"id"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState string            `pulumi:"provisioningState"`
	ResourceGuid      string            `pulumi:"resourceGuid"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
}
