


package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationAccount(ctx *pulumi.Context, args *LookupIntegrationAccountArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountResult, error) {
	var rv LookupIntegrationAccountResult
	err := ctx.Invoke("azure-native:logic/v20160601:getIntegrationAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationAccountArgs struct {
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupIntegrationAccountResult struct {
	Id       string                         `pulumi:"id"`
	Location *string                        `pulumi:"location"`
	Name     string                         `pulumi:"name"`
	Sku      *IntegrationAccountSkuResponse `pulumi:"sku"`
	Tags     map[string]string              `pulumi:"tags"`
	Type     string                         `pulumi:"type"`
}

func LookupIntegrationAccountOutput(ctx *pulumi.Context, args LookupIntegrationAccountOutputArgs, opts ...pulumi.InvokeOption) LookupIntegrationAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIntegrationAccountResult, error) {
			args := v.(LookupIntegrationAccountArgs)
			r, err := LookupIntegrationAccount(ctx, &args, opts...)
			var s LookupIntegrationAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIntegrationAccountResultOutput)
}

type LookupIntegrationAccountOutputArgs struct {
	IntegrationAccountName pulumi.StringInput `pulumi:"integrationAccountName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIntegrationAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountArgs)(nil)).Elem()
}


type LookupIntegrationAccountResultOutput struct{ *pulumi.OutputState }

func (LookupIntegrationAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountResult)(nil)).Elem()
}

func (o LookupIntegrationAccountResultOutput) ToLookupIntegrationAccountResultOutput() LookupIntegrationAccountResultOutput {
	return o
}

func (o LookupIntegrationAccountResultOutput) ToLookupIntegrationAccountResultOutputWithContext(ctx context.Context) LookupIntegrationAccountResultOutput {
	return o
}

func (o LookupIntegrationAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountResultOutput) Sku() IntegrationAccountSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountResult) *IntegrationAccountSkuResponse { return v.Sku }).(IntegrationAccountSkuResponsePtrOutput)
}

func (o LookupIntegrationAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIntegrationAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupIntegrationAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIntegrationAccountResultOutput{})
}
