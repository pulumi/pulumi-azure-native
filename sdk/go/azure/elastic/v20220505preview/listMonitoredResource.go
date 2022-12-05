


package v20220505preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListMonitoredResource(ctx *pulumi.Context, args *ListMonitoredResourceArgs, opts ...pulumi.InvokeOption) (*ListMonitoredResourceResult, error) {
	var rv ListMonitoredResourceResult
	err := ctx.Invoke("azure-native:elastic/v20220505preview:listMonitoredResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMonitoredResourceArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListMonitoredResourceResult struct {
	NextLink *string                     `pulumi:"nextLink"`
	Value    []MonitoredResourceResponse `pulumi:"value"`
}

func ListMonitoredResourceOutput(ctx *pulumi.Context, args ListMonitoredResourceOutputArgs, opts ...pulumi.InvokeOption) ListMonitoredResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListMonitoredResourceResult, error) {
			args := v.(ListMonitoredResourceArgs)
			r, err := ListMonitoredResource(ctx, &args, opts...)
			var s ListMonitoredResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListMonitoredResourceResultOutput)
}

type ListMonitoredResourceOutputArgs struct {
	MonitorName       pulumi.StringInput `pulumi:"monitorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListMonitoredResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMonitoredResourceArgs)(nil)).Elem()
}


type ListMonitoredResourceResultOutput struct{ *pulumi.OutputState }

func (ListMonitoredResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMonitoredResourceResult)(nil)).Elem()
}

func (o ListMonitoredResourceResultOutput) ToListMonitoredResourceResultOutput() ListMonitoredResourceResultOutput {
	return o
}

func (o ListMonitoredResourceResultOutput) ToListMonitoredResourceResultOutputWithContext(ctx context.Context) ListMonitoredResourceResultOutput {
	return o
}

func (o ListMonitoredResourceResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListMonitoredResourceResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListMonitoredResourceResultOutput) Value() MonitoredResourceResponseArrayOutput {
	return o.ApplyT(func(v ListMonitoredResourceResult) []MonitoredResourceResponse { return v.Value }).(MonitoredResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListMonitoredResourceResultOutput{})
}
