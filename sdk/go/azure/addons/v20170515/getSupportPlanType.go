


package v20170515

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSupportPlanType(ctx *pulumi.Context, args *LookupSupportPlanTypeArgs, opts ...pulumi.InvokeOption) (*LookupSupportPlanTypeResult, error) {
	var rv LookupSupportPlanTypeResult
	err := ctx.Invoke("azure-native:addons/v20170515:getSupportPlanType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSupportPlanTypeArgs struct {
	PlanTypeName string `pulumi:"planTypeName"`
	ProviderName string `pulumi:"providerName"`
}


type LookupSupportPlanTypeResult struct {
	Id                string  `pulumi:"id"`
	Name              string  `pulumi:"name"`
	ProvisioningState *string `pulumi:"provisioningState"`
	Type              string  `pulumi:"type"`
}

func LookupSupportPlanTypeOutput(ctx *pulumi.Context, args LookupSupportPlanTypeOutputArgs, opts ...pulumi.InvokeOption) LookupSupportPlanTypeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSupportPlanTypeResult, error) {
			args := v.(LookupSupportPlanTypeArgs)
			r, err := LookupSupportPlanType(ctx, &args, opts...)
			var s LookupSupportPlanTypeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSupportPlanTypeResultOutput)
}

type LookupSupportPlanTypeOutputArgs struct {
	PlanTypeName pulumi.StringInput `pulumi:"planTypeName"`
	ProviderName pulumi.StringInput `pulumi:"providerName"`
}

func (LookupSupportPlanTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSupportPlanTypeArgs)(nil)).Elem()
}


type LookupSupportPlanTypeResultOutput struct{ *pulumi.OutputState }

func (LookupSupportPlanTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSupportPlanTypeResult)(nil)).Elem()
}

func (o LookupSupportPlanTypeResultOutput) ToLookupSupportPlanTypeResultOutput() LookupSupportPlanTypeResultOutput {
	return o
}

func (o LookupSupportPlanTypeResultOutput) ToLookupSupportPlanTypeResultOutputWithContext(ctx context.Context) LookupSupportPlanTypeResultOutput {
	return o
}

func (o LookupSupportPlanTypeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSupportPlanTypeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSupportPlanTypeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSupportPlanTypeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSupportPlanTypeResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSupportPlanTypeResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupSupportPlanTypeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSupportPlanTypeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSupportPlanTypeResultOutput{})
}
