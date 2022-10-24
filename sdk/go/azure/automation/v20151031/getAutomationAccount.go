


package v20151031

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAutomationAccount(ctx *pulumi.Context, args *LookupAutomationAccountArgs, opts ...pulumi.InvokeOption) (*LookupAutomationAccountResult, error) {
	var rv LookupAutomationAccountResult
	err := ctx.Invoke("azure-native:automation/v20151031:getAutomationAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAutomationAccountArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupAutomationAccountResult struct {
	CreationTime     string            `pulumi:"creationTime"`
	Description      *string           `pulumi:"description"`
	Etag             *string           `pulumi:"etag"`
	Id               string            `pulumi:"id"`
	LastModifiedBy   *string           `pulumi:"lastModifiedBy"`
	LastModifiedTime string            `pulumi:"lastModifiedTime"`
	Location         *string           `pulumi:"location"`
	Name             string            `pulumi:"name"`
	Sku              *SkuResponse      `pulumi:"sku"`
	State            string            `pulumi:"state"`
	Tags             map[string]string `pulumi:"tags"`
	Type             string            `pulumi:"type"`
}

func LookupAutomationAccountOutput(ctx *pulumi.Context, args LookupAutomationAccountOutputArgs, opts ...pulumi.InvokeOption) LookupAutomationAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAutomationAccountResult, error) {
			args := v.(LookupAutomationAccountArgs)
			r, err := LookupAutomationAccount(ctx, &args, opts...)
			var s LookupAutomationAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAutomationAccountResultOutput)
}

type LookupAutomationAccountOutputArgs struct {
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAutomationAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutomationAccountArgs)(nil)).Elem()
}


type LookupAutomationAccountResultOutput struct{ *pulumi.OutputState }

func (LookupAutomationAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutomationAccountResult)(nil)).Elem()
}

func (o LookupAutomationAccountResultOutput) ToLookupAutomationAccountResultOutput() LookupAutomationAccountResultOutput {
	return o
}

func (o LookupAutomationAccountResultOutput) ToLookupAutomationAccountResultOutputWithContext(ctx context.Context) LookupAutomationAccountResultOutput {
	return o
}

func (o LookupAutomationAccountResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutomationAccountResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupAutomationAccountResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutomationAccountResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupAutomationAccountResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutomationAccountResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupAutomationAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutomationAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAutomationAccountResultOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutomationAccountResult) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o LookupAutomationAccountResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutomationAccountResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o LookupAutomationAccountResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutomationAccountResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupAutomationAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutomationAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAutomationAccountResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupAutomationAccountResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupAutomationAccountResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutomationAccountResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupAutomationAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAutomationAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAutomationAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutomationAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAutomationAccountResultOutput{})
}
