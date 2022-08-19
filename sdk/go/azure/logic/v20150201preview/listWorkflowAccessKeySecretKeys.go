


package v20150201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkflowAccessKeySecretKeys(ctx *pulumi.Context, args *ListWorkflowAccessKeySecretKeysArgs, opts ...pulumi.InvokeOption) (*ListWorkflowAccessKeySecretKeysResult, error) {
	var rv ListWorkflowAccessKeySecretKeysResult
	err := ctx.Invoke("azure-native:logic/v20150201preview:listWorkflowAccessKeySecretKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkflowAccessKeySecretKeysArgs struct {
	AccessKeyName     string `pulumi:"accessKeyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkflowName      string `pulumi:"workflowName"`
}

type ListWorkflowAccessKeySecretKeysResult struct {
	PrimarySecretKey   string `pulumi:"primarySecretKey"`
	SecondarySecretKey string `pulumi:"secondarySecretKey"`
}

func ListWorkflowAccessKeySecretKeysOutput(ctx *pulumi.Context, args ListWorkflowAccessKeySecretKeysOutputArgs, opts ...pulumi.InvokeOption) ListWorkflowAccessKeySecretKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWorkflowAccessKeySecretKeysResult, error) {
			args := v.(ListWorkflowAccessKeySecretKeysArgs)
			r, err := ListWorkflowAccessKeySecretKeys(ctx, &args, opts...)
			var s ListWorkflowAccessKeySecretKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWorkflowAccessKeySecretKeysResultOutput)
}

type ListWorkflowAccessKeySecretKeysOutputArgs struct {
	AccessKeyName     pulumi.StringInput `pulumi:"accessKeyName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkflowName      pulumi.StringInput `pulumi:"workflowName"`
}

func (ListWorkflowAccessKeySecretKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkflowAccessKeySecretKeysArgs)(nil)).Elem()
}

type ListWorkflowAccessKeySecretKeysResultOutput struct{ *pulumi.OutputState }

func (ListWorkflowAccessKeySecretKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkflowAccessKeySecretKeysResult)(nil)).Elem()
}

func (o ListWorkflowAccessKeySecretKeysResultOutput) ToListWorkflowAccessKeySecretKeysResultOutput() ListWorkflowAccessKeySecretKeysResultOutput {
	return o
}

func (o ListWorkflowAccessKeySecretKeysResultOutput) ToListWorkflowAccessKeySecretKeysResultOutputWithContext(ctx context.Context) ListWorkflowAccessKeySecretKeysResultOutput {
	return o
}

func (o ListWorkflowAccessKeySecretKeysResultOutput) PrimarySecretKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkflowAccessKeySecretKeysResult) string { return v.PrimarySecretKey }).(pulumi.StringOutput)
}

func (o ListWorkflowAccessKeySecretKeysResultOutput) SecondarySecretKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkflowAccessKeySecretKeysResult) string { return v.SecondarySecretKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkflowAccessKeySecretKeysResultOutput{})
}
