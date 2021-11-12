


package v20200701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateDnsZoneGroup(ctx *pulumi.Context, args *LookupPrivateDnsZoneGroupArgs, opts ...pulumi.InvokeOption) (*LookupPrivateDnsZoneGroupResult, error) {
	var rv LookupPrivateDnsZoneGroupResult
	err := ctx.Invoke("azure-native:network/v20200701:getPrivateDnsZoneGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateDnsZoneGroupArgs struct {
	PrivateDnsZoneGroupName string `pulumi:"privateDnsZoneGroupName"`
	PrivateEndpointName     string `pulumi:"privateEndpointName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupPrivateDnsZoneGroupResult struct {
	Etag                  string                         `pulumi:"etag"`
	Id                    *string                        `pulumi:"id"`
	Name                  *string                        `pulumi:"name"`
	PrivateDnsZoneConfigs []PrivateDnsZoneConfigResponse `pulumi:"privateDnsZoneConfigs"`
	ProvisioningState     string                         `pulumi:"provisioningState"`
}
