


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListMonitorApiKeys(ctx *pulumi.Context, args *ListMonitorApiKeysArgs, opts ...pulumi.InvokeOption) (*ListMonitorApiKeysResult, error) {
	var rv ListMonitorApiKeysResult
	err := ctx.Invoke("azure-native:datadog/v20210301:listMonitorApiKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMonitorApiKeysArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListMonitorApiKeysResult struct {
	NextLink *string                 `pulumi:"nextLink"`
	Value    []DatadogApiKeyResponse `pulumi:"value"`
}

func ListMonitorApiKeysOutput(ctx *pulumi.Context, args ListMonitorApiKeysOutputArgs, opts ...pulumi.InvokeOption) ListMonitorApiKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListMonitorApiKeysResult, error) {
			args := v.(ListMonitorApiKeysArgs)
			r, err := ListMonitorApiKeys(ctx, &args, opts...)
			var s ListMonitorApiKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListMonitorApiKeysResultOutput)
}

type ListMonitorApiKeysOutputArgs struct {
	MonitorName       pulumi.StringInput `pulumi:"monitorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListMonitorApiKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMonitorApiKeysArgs)(nil)).Elem()
}


type ListMonitorApiKeysResultOutput struct{ *pulumi.OutputState }

func (ListMonitorApiKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMonitorApiKeysResult)(nil)).Elem()
}

func (o ListMonitorApiKeysResultOutput) ToListMonitorApiKeysResultOutput() ListMonitorApiKeysResultOutput {
	return o
}

func (o ListMonitorApiKeysResultOutput) ToListMonitorApiKeysResultOutputWithContext(ctx context.Context) ListMonitorApiKeysResultOutput {
	return o
}

func (o ListMonitorApiKeysResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListMonitorApiKeysResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListMonitorApiKeysResultOutput) Value() DatadogApiKeyResponseArrayOutput {
	return o.ApplyT(func(v ListMonitorApiKeysResult) []DatadogApiKeyResponse { return v.Value }).(DatadogApiKeyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListMonitorApiKeysResultOutput{})
}
