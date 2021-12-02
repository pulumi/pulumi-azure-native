


package v20180901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateZone(ctx *pulumi.Context, args *LookupPrivateZoneArgs, opts ...pulumi.InvokeOption) (*LookupPrivateZoneResult, error) {
	var rv LookupPrivateZoneResult
	err := ctx.Invoke("azure-native:network/v20180901:getPrivateZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateZoneArgs struct {
	PrivateZoneName   string `pulumi:"privateZoneName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPrivateZoneResult struct {
	Etag                                           *string           `pulumi:"etag"`
	Id                                             string            `pulumi:"id"`
	Location                                       *string           `pulumi:"location"`
	MaxNumberOfRecordSets                          float64           `pulumi:"maxNumberOfRecordSets"`
	MaxNumberOfVirtualNetworkLinks                 float64           `pulumi:"maxNumberOfVirtualNetworkLinks"`
	MaxNumberOfVirtualNetworkLinksWithRegistration float64           `pulumi:"maxNumberOfVirtualNetworkLinksWithRegistration"`
	Name                                           string            `pulumi:"name"`
	NumberOfRecordSets                             float64           `pulumi:"numberOfRecordSets"`
	NumberOfVirtualNetworkLinks                    float64           `pulumi:"numberOfVirtualNetworkLinks"`
	NumberOfVirtualNetworkLinksWithRegistration    float64           `pulumi:"numberOfVirtualNetworkLinksWithRegistration"`
	ProvisioningState                              string            `pulumi:"provisioningState"`
	Tags                                           map[string]string `pulumi:"tags"`
	Type                                           string            `pulumi:"type"`
}
