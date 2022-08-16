


package v20200717preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkloadNetworkSegment(ctx *pulumi.Context, args *LookupWorkloadNetworkSegmentArgs, opts ...pulumi.InvokeOption) (*LookupWorkloadNetworkSegmentResult, error) {
	var rv LookupWorkloadNetworkSegmentResult
	err := ctx.Invoke("azure-native:avs/v20200717preview:getWorkloadNetworkSegment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkloadNetworkSegmentArgs struct {
	PrivateCloudName  string `pulumi:"privateCloudName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SegmentId         string `pulumi:"segmentId"`
}


type LookupWorkloadNetworkSegmentResult struct {
	ConnectedGateway  *string                                 `pulumi:"connectedGateway"`
	DisplayName       *string                                 `pulumi:"displayName"`
	Id                string                                  `pulumi:"id"`
	Name              string                                  `pulumi:"name"`
	PortVif           []WorkloadNetworkSegmentPortVifResponse `pulumi:"portVif"`
	ProvisioningState string                                  `pulumi:"provisioningState"`
	Revision          *float64                                `pulumi:"revision"`
	Status            string                                  `pulumi:"status"`
	Subnet            *WorkloadNetworkSegmentSubnetResponse   `pulumi:"subnet"`
	Type              string                                  `pulumi:"type"`
}

func LookupWorkloadNetworkSegmentOutput(ctx *pulumi.Context, args LookupWorkloadNetworkSegmentOutputArgs, opts ...pulumi.InvokeOption) LookupWorkloadNetworkSegmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkloadNetworkSegmentResult, error) {
			args := v.(LookupWorkloadNetworkSegmentArgs)
			r, err := LookupWorkloadNetworkSegment(ctx, &args, opts...)
			var s LookupWorkloadNetworkSegmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkloadNetworkSegmentResultOutput)
}

type LookupWorkloadNetworkSegmentOutputArgs struct {
	PrivateCloudName  pulumi.StringInput `pulumi:"privateCloudName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SegmentId         pulumi.StringInput `pulumi:"segmentId"`
}

func (LookupWorkloadNetworkSegmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadNetworkSegmentArgs)(nil)).Elem()
}


type LookupWorkloadNetworkSegmentResultOutput struct{ *pulumi.OutputState }

func (LookupWorkloadNetworkSegmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadNetworkSegmentResult)(nil)).Elem()
}

func (o LookupWorkloadNetworkSegmentResultOutput) ToLookupWorkloadNetworkSegmentResultOutput() LookupWorkloadNetworkSegmentResultOutput {
	return o
}

func (o LookupWorkloadNetworkSegmentResultOutput) ToLookupWorkloadNetworkSegmentResultOutputWithContext(ctx context.Context) LookupWorkloadNetworkSegmentResultOutput {
	return o
}

func (o LookupWorkloadNetworkSegmentResultOutput) ConnectedGateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkSegmentResult) *string { return v.ConnectedGateway }).(pulumi.StringPtrOutput)
}

func (o LookupWorkloadNetworkSegmentResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkSegmentResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupWorkloadNetworkSegmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkSegmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkSegmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkSegmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkSegmentResultOutput) PortVif() WorkloadNetworkSegmentPortVifResponseArrayOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkSegmentResult) []WorkloadNetworkSegmentPortVifResponse { return v.PortVif }).(WorkloadNetworkSegmentPortVifResponseArrayOutput)
}

func (o LookupWorkloadNetworkSegmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkSegmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkSegmentResultOutput) Revision() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkSegmentResult) *float64 { return v.Revision }).(pulumi.Float64PtrOutput)
}

func (o LookupWorkloadNetworkSegmentResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkSegmentResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkSegmentResultOutput) Subnet() WorkloadNetworkSegmentSubnetResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkSegmentResult) *WorkloadNetworkSegmentSubnetResponse { return v.Subnet }).(WorkloadNetworkSegmentSubnetResponsePtrOutput)
}

func (o LookupWorkloadNetworkSegmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkSegmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkloadNetworkSegmentResultOutput{})
}
