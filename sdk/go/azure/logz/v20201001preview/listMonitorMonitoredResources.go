


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListMonitorMonitoredResources(ctx *pulumi.Context, args *ListMonitorMonitoredResourcesArgs, opts ...pulumi.InvokeOption) (*ListMonitorMonitoredResourcesResult, error) {
	var rv ListMonitorMonitoredResourcesResult
	err := ctx.Invoke("azure-native:logz/v20201001preview:listMonitorMonitoredResources", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMonitorMonitoredResourcesArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListMonitorMonitoredResourcesResult struct {
	NextLink *string                     `pulumi:"nextLink"`
	Value    []MonitoredResourceResponse `pulumi:"value"`
}

func ListMonitorMonitoredResourcesOutput(ctx *pulumi.Context, args ListMonitorMonitoredResourcesOutputArgs, opts ...pulumi.InvokeOption) ListMonitorMonitoredResourcesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListMonitorMonitoredResourcesResult, error) {
			args := v.(ListMonitorMonitoredResourcesArgs)
			r, err := ListMonitorMonitoredResources(ctx, &args, opts...)
			var s ListMonitorMonitoredResourcesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListMonitorMonitoredResourcesResultOutput)
}

type ListMonitorMonitoredResourcesOutputArgs struct {
	MonitorName       pulumi.StringInput `pulumi:"monitorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListMonitorMonitoredResourcesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMonitorMonitoredResourcesArgs)(nil)).Elem()
}


type ListMonitorMonitoredResourcesResultOutput struct{ *pulumi.OutputState }

func (ListMonitorMonitoredResourcesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMonitorMonitoredResourcesResult)(nil)).Elem()
}

func (o ListMonitorMonitoredResourcesResultOutput) ToListMonitorMonitoredResourcesResultOutput() ListMonitorMonitoredResourcesResultOutput {
	return o
}

func (o ListMonitorMonitoredResourcesResultOutput) ToListMonitorMonitoredResourcesResultOutputWithContext(ctx context.Context) ListMonitorMonitoredResourcesResultOutput {
	return o
}

func (o ListMonitorMonitoredResourcesResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListMonitorMonitoredResourcesResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListMonitorMonitoredResourcesResultOutput) Value() MonitoredResourceResponseArrayOutput {
	return o.ApplyT(func(v ListMonitorMonitoredResourcesResult) []MonitoredResourceResponse { return v.Value }).(MonitoredResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListMonitorMonitoredResourcesResultOutput{})
}
