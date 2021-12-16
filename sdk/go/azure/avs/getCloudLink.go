


package avs

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCloudLink(ctx *pulumi.Context, args *LookupCloudLinkArgs, opts ...pulumi.InvokeOption) (*LookupCloudLinkResult, error) {
	var rv LookupCloudLinkResult
	err := ctx.Invoke("azure-native:avs:getCloudLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCloudLinkArgs struct {
	CloudLinkName     string `pulumi:"cloudLinkName"`
	PrivateCloudName  string `pulumi:"privateCloudName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCloudLinkResult struct {
	Id          string  `pulumi:"id"`
	LinkedCloud *string `pulumi:"linkedCloud"`
	Name        string  `pulumi:"name"`
	Status      string  `pulumi:"status"`
	Type        string  `pulumi:"type"`
}
