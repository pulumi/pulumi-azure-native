


package v20180901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupChannel(ctx *pulumi.Context, args *LookupChannelArgs, opts ...pulumi.InvokeOption) (*LookupChannelResult, error) {
	var rv LookupChannelResult
	err := ctx.Invoke("azure-native:engagementfabric/v20180901preview:getChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupChannelArgs struct {
	AccountName       string `pulumi:"accountName"`
	ChannelName       string `pulumi:"channelName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupChannelResult struct {
	ChannelFunctions []string          `pulumi:"channelFunctions"`
	ChannelType      string            `pulumi:"channelType"`
	Credentials      map[string]string `pulumi:"credentials"`
	Id               string            `pulumi:"id"`
	Name             string            `pulumi:"name"`
	Type             string            `pulumi:"type"`
}
