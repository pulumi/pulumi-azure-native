


package v20180101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkWatcher(ctx *pulumi.Context, args *LookupNetworkWatcherArgs, opts ...pulumi.InvokeOption) (*LookupNetworkWatcherResult, error) {
	var rv LookupNetworkWatcherResult
	err := ctx.Invoke("azure-native:network/v20180101:getNetworkWatcher", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupNetworkWatcherArgs struct {
	NetworkWatcherName string `pulumi:"networkWatcherName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupNetworkWatcherResult struct {
	Etag              *string           `pulumi:"etag"`
	Id                string            `pulumi:"id"`
	Location          string            `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState string            `pulumi:"provisioningState"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
}


func (val *LookupNetworkWatcherResult) Defaults() *LookupNetworkWatcherResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Etag) {
		etag_ := "A unique read-only string that changes whenever the resource is updated."
		tmp.Etag = &etag_
	}
	return &tmp
}
