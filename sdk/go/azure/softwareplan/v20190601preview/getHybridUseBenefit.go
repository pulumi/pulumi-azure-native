


package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHybridUseBenefit(ctx *pulumi.Context, args *LookupHybridUseBenefitArgs, opts ...pulumi.InvokeOption) (*LookupHybridUseBenefitResult, error) {
	var rv LookupHybridUseBenefitResult
	err := ctx.Invoke("azure-native:softwareplan/v20190601preview:getHybridUseBenefit", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHybridUseBenefitArgs struct {
	PlanId string `pulumi:"planId"`
	Scope  string `pulumi:"scope"`
}


type LookupHybridUseBenefitResult struct {
	CreatedDate       string      `pulumi:"createdDate"`
	Etag              int         `pulumi:"etag"`
	Id                string      `pulumi:"id"`
	LastUpdatedDate   string      `pulumi:"lastUpdatedDate"`
	Name              string      `pulumi:"name"`
	ProvisioningState string      `pulumi:"provisioningState"`
	Sku               SkuResponse `pulumi:"sku"`
	Type              string      `pulumi:"type"`
}

func LookupHybridUseBenefitOutput(ctx *pulumi.Context, args LookupHybridUseBenefitOutputArgs, opts ...pulumi.InvokeOption) LookupHybridUseBenefitResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHybridUseBenefitResult, error) {
			args := v.(LookupHybridUseBenefitArgs)
			r, err := LookupHybridUseBenefit(ctx, &args, opts...)
			var s LookupHybridUseBenefitResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHybridUseBenefitResultOutput)
}

type LookupHybridUseBenefitOutputArgs struct {
	PlanId pulumi.StringInput `pulumi:"planId"`
	Scope  pulumi.StringInput `pulumi:"scope"`
}

func (LookupHybridUseBenefitOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHybridUseBenefitArgs)(nil)).Elem()
}


type LookupHybridUseBenefitResultOutput struct{ *pulumi.OutputState }

func (LookupHybridUseBenefitResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHybridUseBenefitResult)(nil)).Elem()
}

func (o LookupHybridUseBenefitResultOutput) ToLookupHybridUseBenefitResultOutput() LookupHybridUseBenefitResultOutput {
	return o
}

func (o LookupHybridUseBenefitResultOutput) ToLookupHybridUseBenefitResultOutputWithContext(ctx context.Context) LookupHybridUseBenefitResultOutput {
	return o
}

func (o LookupHybridUseBenefitResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridUseBenefitResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o LookupHybridUseBenefitResultOutput) Etag() pulumi.IntOutput {
	return o.ApplyT(func(v LookupHybridUseBenefitResult) int { return v.Etag }).(pulumi.IntOutput)
}

func (o LookupHybridUseBenefitResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridUseBenefitResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupHybridUseBenefitResultOutput) LastUpdatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridUseBenefitResult) string { return v.LastUpdatedDate }).(pulumi.StringOutput)
}

func (o LookupHybridUseBenefitResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridUseBenefitResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupHybridUseBenefitResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridUseBenefitResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupHybridUseBenefitResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupHybridUseBenefitResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupHybridUseBenefitResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridUseBenefitResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHybridUseBenefitResultOutput{})
}
