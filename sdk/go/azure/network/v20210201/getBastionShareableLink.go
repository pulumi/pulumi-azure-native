


package v20210201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetBastionShareableLink(ctx *pulumi.Context, args *GetBastionShareableLinkArgs, opts ...pulumi.InvokeOption) (*GetBastionShareableLinkResult, error) {
	var rv GetBastionShareableLinkResult
	err := ctx.Invoke("azure-native:network/v20210201:getBastionShareableLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetBastionShareableLinkArgs struct {
	BastionHostName   string                 `pulumi:"bastionHostName"`
	ResourceGroupName string                 `pulumi:"resourceGroupName"`
	Vms               []BastionShareableLink `pulumi:"vms"`
}


type GetBastionShareableLinkResult struct {
	NextLink *string                        `pulumi:"nextLink"`
	Value    []BastionShareableLinkResponse `pulumi:"value"`
}
