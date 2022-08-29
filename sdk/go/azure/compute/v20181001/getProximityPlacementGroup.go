


package v20181001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupProximityPlacementGroup(ctx *pulumi.Context, args *LookupProximityPlacementGroupArgs, opts ...pulumi.InvokeOption) (*LookupProximityPlacementGroupResult, error) {
	var rv LookupProximityPlacementGroupResult
	err := ctx.Invoke("azure-native:compute/v20181001:getProximityPlacementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProximityPlacementGroupArgs struct {
	ProximityPlacementGroupName string `pulumi:"proximityPlacementGroupName"`
	ResourceGroupName           string `pulumi:"resourceGroupName"`
}


type LookupProximityPlacementGroupResult struct {
	AvailabilitySets            []SubResourceResponse `pulumi:"availabilitySets"`
	Id                          string                `pulumi:"id"`
	Location                    string                `pulumi:"location"`
	Name                        string                `pulumi:"name"`
	ProximityPlacementGroupType *string               `pulumi:"proximityPlacementGroupType"`
	Tags                        map[string]string     `pulumi:"tags"`
	Type                        string                `pulumi:"type"`
	VirtualMachineScaleSets     []SubResourceResponse `pulumi:"virtualMachineScaleSets"`
	VirtualMachines             []SubResourceResponse `pulumi:"virtualMachines"`
}

func LookupProximityPlacementGroupOutput(ctx *pulumi.Context, args LookupProximityPlacementGroupOutputArgs, opts ...pulumi.InvokeOption) LookupProximityPlacementGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProximityPlacementGroupResult, error) {
			args := v.(LookupProximityPlacementGroupArgs)
			r, err := LookupProximityPlacementGroup(ctx, &args, opts...)
			var s LookupProximityPlacementGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProximityPlacementGroupResultOutput)
}

type LookupProximityPlacementGroupOutputArgs struct {
	ProximityPlacementGroupName pulumi.StringInput `pulumi:"proximityPlacementGroupName"`
	ResourceGroupName           pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupProximityPlacementGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProximityPlacementGroupArgs)(nil)).Elem()
}


type LookupProximityPlacementGroupResultOutput struct{ *pulumi.OutputState }

func (LookupProximityPlacementGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProximityPlacementGroupResult)(nil)).Elem()
}

func (o LookupProximityPlacementGroupResultOutput) ToLookupProximityPlacementGroupResultOutput() LookupProximityPlacementGroupResultOutput {
	return o
}

func (o LookupProximityPlacementGroupResultOutput) ToLookupProximityPlacementGroupResultOutputWithContext(ctx context.Context) LookupProximityPlacementGroupResultOutput {
	return o
}

func (o LookupProximityPlacementGroupResultOutput) AvailabilitySets() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupProximityPlacementGroupResult) []SubResourceResponse { return v.AvailabilitySets }).(SubResourceResponseArrayOutput)
}

func (o LookupProximityPlacementGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProximityPlacementGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProximityPlacementGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProximityPlacementGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupProximityPlacementGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProximityPlacementGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupProximityPlacementGroupResultOutput) ProximityPlacementGroupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProximityPlacementGroupResult) *string { return v.ProximityPlacementGroupType }).(pulumi.StringPtrOutput)
}

func (o LookupProximityPlacementGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProximityPlacementGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupProximityPlacementGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProximityPlacementGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupProximityPlacementGroupResultOutput) VirtualMachineScaleSets() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupProximityPlacementGroupResult) []SubResourceResponse { return v.VirtualMachineScaleSets }).(SubResourceResponseArrayOutput)
}

func (o LookupProximityPlacementGroupResultOutput) VirtualMachines() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupProximityPlacementGroupResult) []SubResourceResponse { return v.VirtualMachines }).(SubResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProximityPlacementGroupResultOutput{})
}
