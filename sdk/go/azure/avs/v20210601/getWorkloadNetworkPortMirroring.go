


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkloadNetworkPortMirroring(ctx *pulumi.Context, args *LookupWorkloadNetworkPortMirroringArgs, opts ...pulumi.InvokeOption) (*LookupWorkloadNetworkPortMirroringResult, error) {
	var rv LookupWorkloadNetworkPortMirroringResult
	err := ctx.Invoke("azure-native:avs/v20210601:getWorkloadNetworkPortMirroring", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkloadNetworkPortMirroringArgs struct {
	PortMirroringId   string `pulumi:"portMirroringId"`
	PrivateCloudName  string `pulumi:"privateCloudName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWorkloadNetworkPortMirroringResult struct {
	Destination       *string  `pulumi:"destination"`
	Direction         *string  `pulumi:"direction"`
	DisplayName       *string  `pulumi:"displayName"`
	Id                string   `pulumi:"id"`
	Name              string   `pulumi:"name"`
	ProvisioningState string   `pulumi:"provisioningState"`
	Revision          *float64 `pulumi:"revision"`
	Source            *string  `pulumi:"source"`
	Status            string   `pulumi:"status"`
	Type              string   `pulumi:"type"`
}

func LookupWorkloadNetworkPortMirroringOutput(ctx *pulumi.Context, args LookupWorkloadNetworkPortMirroringOutputArgs, opts ...pulumi.InvokeOption) LookupWorkloadNetworkPortMirroringResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkloadNetworkPortMirroringResult, error) {
			args := v.(LookupWorkloadNetworkPortMirroringArgs)
			r, err := LookupWorkloadNetworkPortMirroring(ctx, &args, opts...)
			var s LookupWorkloadNetworkPortMirroringResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkloadNetworkPortMirroringResultOutput)
}

type LookupWorkloadNetworkPortMirroringOutputArgs struct {
	PortMirroringId   pulumi.StringInput `pulumi:"portMirroringId"`
	PrivateCloudName  pulumi.StringInput `pulumi:"privateCloudName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWorkloadNetworkPortMirroringOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadNetworkPortMirroringArgs)(nil)).Elem()
}


type LookupWorkloadNetworkPortMirroringResultOutput struct{ *pulumi.OutputState }

func (LookupWorkloadNetworkPortMirroringResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadNetworkPortMirroringResult)(nil)).Elem()
}

func (o LookupWorkloadNetworkPortMirroringResultOutput) ToLookupWorkloadNetworkPortMirroringResultOutput() LookupWorkloadNetworkPortMirroringResultOutput {
	return o
}

func (o LookupWorkloadNetworkPortMirroringResultOutput) ToLookupWorkloadNetworkPortMirroringResultOutputWithContext(ctx context.Context) LookupWorkloadNetworkPortMirroringResultOutput {
	return o
}

func (o LookupWorkloadNetworkPortMirroringResultOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkPortMirroringResult) *string { return v.Destination }).(pulumi.StringPtrOutput)
}

func (o LookupWorkloadNetworkPortMirroringResultOutput) Direction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkPortMirroringResult) *string { return v.Direction }).(pulumi.StringPtrOutput)
}

func (o LookupWorkloadNetworkPortMirroringResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkPortMirroringResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupWorkloadNetworkPortMirroringResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkPortMirroringResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkPortMirroringResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkPortMirroringResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkPortMirroringResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkPortMirroringResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkPortMirroringResultOutput) Revision() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkPortMirroringResult) *float64 { return v.Revision }).(pulumi.Float64PtrOutput)
}

func (o LookupWorkloadNetworkPortMirroringResultOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkPortMirroringResult) *string { return v.Source }).(pulumi.StringPtrOutput)
}

func (o LookupWorkloadNetworkPortMirroringResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkPortMirroringResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkPortMirroringResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkPortMirroringResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkloadNetworkPortMirroringResultOutput{})
}
