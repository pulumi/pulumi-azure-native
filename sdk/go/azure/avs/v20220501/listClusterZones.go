


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListClusterZones(ctx *pulumi.Context, args *ListClusterZonesArgs, opts ...pulumi.InvokeOption) (*ListClusterZonesResult, error) {
	var rv ListClusterZonesResult
	err := ctx.Invoke("azure-native:avs/v20220501:listClusterZones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListClusterZonesArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	PrivateCloudName  string `pulumi:"privateCloudName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListClusterZonesResult struct {
	Zones []ClusterZoneResponse `pulumi:"zones"`
}

func ListClusterZonesOutput(ctx *pulumi.Context, args ListClusterZonesOutputArgs, opts ...pulumi.InvokeOption) ListClusterZonesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListClusterZonesResult, error) {
			args := v.(ListClusterZonesArgs)
			r, err := ListClusterZones(ctx, &args, opts...)
			var s ListClusterZonesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListClusterZonesResultOutput)
}

type ListClusterZonesOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	PrivateCloudName  pulumi.StringInput `pulumi:"privateCloudName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListClusterZonesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListClusterZonesArgs)(nil)).Elem()
}


type ListClusterZonesResultOutput struct{ *pulumi.OutputState }

func (ListClusterZonesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListClusterZonesResult)(nil)).Elem()
}

func (o ListClusterZonesResultOutput) ToListClusterZonesResultOutput() ListClusterZonesResultOutput {
	return o
}

func (o ListClusterZonesResultOutput) ToListClusterZonesResultOutputWithContext(ctx context.Context) ListClusterZonesResultOutput {
	return o
}

func (o ListClusterZonesResultOutput) Zones() ClusterZoneResponseArrayOutput {
	return o.ApplyT(func(v ListClusterZonesResult) []ClusterZoneResponse { return v.Zones }).(ClusterZoneResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListClusterZonesResultOutput{})
}
