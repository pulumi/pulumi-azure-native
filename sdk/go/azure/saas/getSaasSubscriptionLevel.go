


package saas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSaasSubscriptionLevel(ctx *pulumi.Context, args *LookupSaasSubscriptionLevelArgs, opts ...pulumi.InvokeOption) (*LookupSaasSubscriptionLevelResult, error) {
	var rv LookupSaasSubscriptionLevelResult
	err := ctx.Invoke("azure-native:saas:getSaasSubscriptionLevel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSaasSubscriptionLevelArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupSaasSubscriptionLevelResult struct {
	Id         string                         `pulumi:"id"`
	Name       string                         `pulumi:"name"`
	Properties SaasResourceResponseProperties `pulumi:"properties"`
	Tags       map[string]string              `pulumi:"tags"`
	Type       string                         `pulumi:"type"`
}

func LookupSaasSubscriptionLevelOutput(ctx *pulumi.Context, args LookupSaasSubscriptionLevelOutputArgs, opts ...pulumi.InvokeOption) LookupSaasSubscriptionLevelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSaasSubscriptionLevelResult, error) {
			args := v.(LookupSaasSubscriptionLevelArgs)
			r, err := LookupSaasSubscriptionLevel(ctx, &args, opts...)
			var s LookupSaasSubscriptionLevelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSaasSubscriptionLevelResultOutput)
}

type LookupSaasSubscriptionLevelOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupSaasSubscriptionLevelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSaasSubscriptionLevelArgs)(nil)).Elem()
}


type LookupSaasSubscriptionLevelResultOutput struct{ *pulumi.OutputState }

func (LookupSaasSubscriptionLevelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSaasSubscriptionLevelResult)(nil)).Elem()
}

func (o LookupSaasSubscriptionLevelResultOutput) ToLookupSaasSubscriptionLevelResultOutput() LookupSaasSubscriptionLevelResultOutput {
	return o
}

func (o LookupSaasSubscriptionLevelResultOutput) ToLookupSaasSubscriptionLevelResultOutputWithContext(ctx context.Context) LookupSaasSubscriptionLevelResultOutput {
	return o
}

func (o LookupSaasSubscriptionLevelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSaasSubscriptionLevelResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSaasSubscriptionLevelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSaasSubscriptionLevelResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSaasSubscriptionLevelResultOutput) Properties() SaasResourceResponsePropertiesOutput {
	return o.ApplyT(func(v LookupSaasSubscriptionLevelResult) SaasResourceResponseProperties { return v.Properties }).(SaasResourceResponsePropertiesOutput)
}

func (o LookupSaasSubscriptionLevelResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSaasSubscriptionLevelResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSaasSubscriptionLevelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSaasSubscriptionLevelResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSaasSubscriptionLevelResultOutput{})
}
