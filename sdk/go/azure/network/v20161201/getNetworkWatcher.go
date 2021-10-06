


package v20161201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkWatcher(ctx *pulumi.Context, args *LookupNetworkWatcherArgs, opts ...pulumi.InvokeOption) (*LookupNetworkWatcherResult, error) {
	var rv LookupNetworkWatcherResult
	err := ctx.Invoke("azure-native:network/v20161201:getNetworkWatcher", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkWatcherArgs struct {
	NetworkWatcherName string `pulumi:"networkWatcherName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupNetworkWatcherResult struct {
	Etag              *string           `pulumi:"etag"`
	Id                *string           `pulumi:"id"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState string            `pulumi:"provisioningState"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
}
