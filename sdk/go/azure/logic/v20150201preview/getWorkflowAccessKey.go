


package v20150201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkflowAccessKey(ctx *pulumi.Context, args *LookupWorkflowAccessKeyArgs, opts ...pulumi.InvokeOption) (*LookupWorkflowAccessKeyResult, error) {
	var rv LookupWorkflowAccessKeyResult
	err := ctx.Invoke("azure-native:logic/v20150201preview:getWorkflowAccessKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkflowAccessKeyArgs struct {
	AccessKeyName     string `pulumi:"accessKeyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkflowName      string `pulumi:"workflowName"`
}

type LookupWorkflowAccessKeyResult struct {
	Id        *string `pulumi:"id"`
	Name      string  `pulumi:"name"`
	NotAfter  *string `pulumi:"notAfter"`
	NotBefore *string `pulumi:"notBefore"`
	Type      string  `pulumi:"type"`
}

func LookupWorkflowAccessKeyOutput(ctx *pulumi.Context, args LookupWorkflowAccessKeyOutputArgs, opts ...pulumi.InvokeOption) LookupWorkflowAccessKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkflowAccessKeyResult, error) {
			args := v.(LookupWorkflowAccessKeyArgs)
			r, err := LookupWorkflowAccessKey(ctx, &args, opts...)
			var s LookupWorkflowAccessKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkflowAccessKeyResultOutput)
}

type LookupWorkflowAccessKeyOutputArgs struct {
	AccessKeyName     pulumi.StringInput `pulumi:"accessKeyName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkflowName      pulumi.StringInput `pulumi:"workflowName"`
}

func (LookupWorkflowAccessKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkflowAccessKeyArgs)(nil)).Elem()
}

type LookupWorkflowAccessKeyResultOutput struct{ *pulumi.OutputState }

func (LookupWorkflowAccessKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkflowAccessKeyResult)(nil)).Elem()
}

func (o LookupWorkflowAccessKeyResultOutput) ToLookupWorkflowAccessKeyResultOutput() LookupWorkflowAccessKeyResultOutput {
	return o
}

func (o LookupWorkflowAccessKeyResultOutput) ToLookupWorkflowAccessKeyResultOutputWithContext(ctx context.Context) LookupWorkflowAccessKeyResultOutput {
	return o
}

func (o LookupWorkflowAccessKeyResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowAccessKeyResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupWorkflowAccessKeyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowAccessKeyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkflowAccessKeyResultOutput) NotAfter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowAccessKeyResult) *string { return v.NotAfter }).(pulumi.StringPtrOutput)
}

func (o LookupWorkflowAccessKeyResultOutput) NotBefore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowAccessKeyResult) *string { return v.NotBefore }).(pulumi.StringPtrOutput)
}

func (o LookupWorkflowAccessKeyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowAccessKeyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkflowAccessKeyResultOutput{})
}
