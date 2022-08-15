


package v20181101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDdosProtectionPlan(ctx *pulumi.Context, args *LookupDdosProtectionPlanArgs, opts ...pulumi.InvokeOption) (*LookupDdosProtectionPlanResult, error) {
	var rv LookupDdosProtectionPlanResult
	err := ctx.Invoke("azure-native:network/v20181101:getDdosProtectionPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDdosProtectionPlanArgs struct {
	DdosProtectionPlanName string `pulumi:"ddosProtectionPlanName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupDdosProtectionPlanResult struct {
	Etag              string                `pulumi:"etag"`
	Id                string                `pulumi:"id"`
	Location          *string               `pulumi:"location"`
	Name              string                `pulumi:"name"`
	ProvisioningState string                `pulumi:"provisioningState"`
	ResourceGuid      string                `pulumi:"resourceGuid"`
	Tags              map[string]string     `pulumi:"tags"`
	Type              string                `pulumi:"type"`
	VirtualNetworks   []SubResourceResponse `pulumi:"virtualNetworks"`
}

func LookupDdosProtectionPlanOutput(ctx *pulumi.Context, args LookupDdosProtectionPlanOutputArgs, opts ...pulumi.InvokeOption) LookupDdosProtectionPlanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDdosProtectionPlanResult, error) {
			args := v.(LookupDdosProtectionPlanArgs)
			r, err := LookupDdosProtectionPlan(ctx, &args, opts...)
			var s LookupDdosProtectionPlanResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDdosProtectionPlanResultOutput)
}

type LookupDdosProtectionPlanOutputArgs struct {
	DdosProtectionPlanName pulumi.StringInput `pulumi:"ddosProtectionPlanName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDdosProtectionPlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDdosProtectionPlanArgs)(nil)).Elem()
}


type LookupDdosProtectionPlanResultOutput struct{ *pulumi.OutputState }

func (LookupDdosProtectionPlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDdosProtectionPlanResult)(nil)).Elem()
}

func (o LookupDdosProtectionPlanResultOutput) ToLookupDdosProtectionPlanResultOutput() LookupDdosProtectionPlanResultOutput {
	return o
}

func (o LookupDdosProtectionPlanResultOutput) ToLookupDdosProtectionPlanResultOutputWithContext(ctx context.Context) LookupDdosProtectionPlanResultOutput {
	return o
}

func (o LookupDdosProtectionPlanResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdosProtectionPlanResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupDdosProtectionPlanResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdosProtectionPlanResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDdosProtectionPlanResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDdosProtectionPlanResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDdosProtectionPlanResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdosProtectionPlanResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDdosProtectionPlanResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdosProtectionPlanResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDdosProtectionPlanResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdosProtectionPlanResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupDdosProtectionPlanResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDdosProtectionPlanResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDdosProtectionPlanResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdosProtectionPlanResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDdosProtectionPlanResultOutput) VirtualNetworks() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupDdosProtectionPlanResult) []SubResourceResponse { return v.VirtualNetworks }).(SubResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDdosProtectionPlanResultOutput{})
}
