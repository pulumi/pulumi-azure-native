


package v20160901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupAppliance(ctx *pulumi.Context, args *LookupApplianceArgs, opts ...pulumi.InvokeOption) (*LookupApplianceResult, error) {
	var rv LookupApplianceResult
	err := ctx.Invoke("azure-native:solutions/v20160901preview:getAppliance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplianceArgs struct {
	ApplianceName     string `pulumi:"applianceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupApplianceResult struct {
	ApplianceDefinitionId  *string           `pulumi:"applianceDefinitionId"`
	Id                     string            `pulumi:"id"`
	Identity               *IdentityResponse `pulumi:"identity"`
	Kind                   *string           `pulumi:"kind"`
	Location               *string           `pulumi:"location"`
	ManagedBy              *string           `pulumi:"managedBy"`
	ManagedResourceGroupId string            `pulumi:"managedResourceGroupId"`
	Name                   string            `pulumi:"name"`
	Outputs                interface{}       `pulumi:"outputs"`
	Parameters             interface{}       `pulumi:"parameters"`
	Plan                   *PlanResponse     `pulumi:"plan"`
	ProvisioningState      string            `pulumi:"provisioningState"`
	Sku                    *SkuResponse      `pulumi:"sku"`
	Tags                   map[string]string `pulumi:"tags"`
	Type                   string            `pulumi:"type"`
	UiDefinitionUri        *string           `pulumi:"uiDefinitionUri"`
}

func LookupApplianceOutput(ctx *pulumi.Context, args LookupApplianceOutputArgs, opts ...pulumi.InvokeOption) LookupApplianceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplianceResult, error) {
			args := v.(LookupApplianceArgs)
			r, err := LookupAppliance(ctx, &args, opts...)
			var s LookupApplianceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplianceResultOutput)
}

type LookupApplianceOutputArgs struct {
	ApplianceName     pulumi.StringInput `pulumi:"applianceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplianceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplianceArgs)(nil)).Elem()
}


type LookupApplianceResultOutput struct{ *pulumi.OutputState }

func (LookupApplianceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplianceResult)(nil)).Elem()
}

func (o LookupApplianceResultOutput) ToLookupApplianceResultOutput() LookupApplianceResultOutput {
	return o
}

func (o LookupApplianceResultOutput) ToLookupApplianceResultOutputWithContext(ctx context.Context) LookupApplianceResultOutput {
	return o
}

func (o LookupApplianceResultOutput) ApplianceDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplianceResult) *string { return v.ApplianceDefinitionId }).(pulumi.StringPtrOutput)
}

func (o LookupApplianceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplianceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApplianceResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupApplianceResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupApplianceResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplianceResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupApplianceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplianceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupApplianceResultOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplianceResult) *string { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

func (o LookupApplianceResultOutput) ManagedResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplianceResult) string { return v.ManagedResourceGroupId }).(pulumi.StringOutput)
}

func (o LookupApplianceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplianceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApplianceResultOutput) Outputs() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupApplianceResult) interface{} { return v.Outputs }).(pulumi.AnyOutput)
}

func (o LookupApplianceResultOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupApplianceResult) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

func (o LookupApplianceResultOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v LookupApplianceResult) *PlanResponse { return v.Plan }).(PlanResponsePtrOutput)
}

func (o LookupApplianceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplianceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupApplianceResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupApplianceResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupApplianceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplianceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupApplianceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplianceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupApplianceResultOutput) UiDefinitionUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplianceResult) *string { return v.UiDefinitionUri }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplianceResultOutput{})
}
