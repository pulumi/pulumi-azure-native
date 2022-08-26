


package v20210701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAvailabilitySet(ctx *pulumi.Context, args *LookupAvailabilitySetArgs, opts ...pulumi.InvokeOption) (*LookupAvailabilitySetResult, error) {
	var rv LookupAvailabilitySetResult
	err := ctx.Invoke("azure-native:compute/v20210701:getAvailabilitySet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAvailabilitySetArgs struct {
	AvailabilitySetName string `pulumi:"availabilitySetName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupAvailabilitySetResult struct {
	Id                        string                       `pulumi:"id"`
	Location                  string                       `pulumi:"location"`
	Name                      string                       `pulumi:"name"`
	PlatformFaultDomainCount  *int                         `pulumi:"platformFaultDomainCount"`
	PlatformUpdateDomainCount *int                         `pulumi:"platformUpdateDomainCount"`
	ProximityPlacementGroup   *SubResourceResponse         `pulumi:"proximityPlacementGroup"`
	Sku                       *SkuResponse                 `pulumi:"sku"`
	Statuses                  []InstanceViewStatusResponse `pulumi:"statuses"`
	Tags                      map[string]string            `pulumi:"tags"`
	Type                      string                       `pulumi:"type"`
	VirtualMachines           []SubResourceResponse        `pulumi:"virtualMachines"`
}

func LookupAvailabilitySetOutput(ctx *pulumi.Context, args LookupAvailabilitySetOutputArgs, opts ...pulumi.InvokeOption) LookupAvailabilitySetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAvailabilitySetResult, error) {
			args := v.(LookupAvailabilitySetArgs)
			r, err := LookupAvailabilitySet(ctx, &args, opts...)
			var s LookupAvailabilitySetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAvailabilitySetResultOutput)
}

type LookupAvailabilitySetOutputArgs struct {
	AvailabilitySetName pulumi.StringInput `pulumi:"availabilitySetName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAvailabilitySetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAvailabilitySetArgs)(nil)).Elem()
}


type LookupAvailabilitySetResultOutput struct{ *pulumi.OutputState }

func (LookupAvailabilitySetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAvailabilitySetResult)(nil)).Elem()
}

func (o LookupAvailabilitySetResultOutput) ToLookupAvailabilitySetResultOutput() LookupAvailabilitySetResultOutput {
	return o
}

func (o LookupAvailabilitySetResultOutput) ToLookupAvailabilitySetResultOutputWithContext(ctx context.Context) LookupAvailabilitySetResultOutput {
	return o
}

func (o LookupAvailabilitySetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAvailabilitySetResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAvailabilitySetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAvailabilitySetResultOutput) PlatformFaultDomainCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) *int { return v.PlatformFaultDomainCount }).(pulumi.IntPtrOutput)
}

func (o LookupAvailabilitySetResultOutput) PlatformUpdateDomainCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) *int { return v.PlatformUpdateDomainCount }).(pulumi.IntPtrOutput)
}

func (o LookupAvailabilitySetResultOutput) ProximityPlacementGroup() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) *SubResourceResponse { return v.ProximityPlacementGroup }).(SubResourceResponsePtrOutput)
}

func (o LookupAvailabilitySetResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupAvailabilitySetResultOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) []InstanceViewStatusResponse { return v.Statuses }).(InstanceViewStatusResponseArrayOutput)
}

func (o LookupAvailabilitySetResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAvailabilitySetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupAvailabilitySetResultOutput) VirtualMachines() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) []SubResourceResponse { return v.VirtualMachines }).(SubResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAvailabilitySetResultOutput{})
}
