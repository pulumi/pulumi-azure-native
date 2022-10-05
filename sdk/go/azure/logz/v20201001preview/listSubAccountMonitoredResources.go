


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSubAccountMonitoredResources(ctx *pulumi.Context, args *ListSubAccountMonitoredResourcesArgs, opts ...pulumi.InvokeOption) (*ListSubAccountMonitoredResourcesResult, error) {
	var rv ListSubAccountMonitoredResourcesResult
	err := ctx.Invoke("azure-native:logz/v20201001preview:listSubAccountMonitoredResources", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSubAccountMonitoredResourcesArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SubAccountName    string `pulumi:"subAccountName"`
}


type ListSubAccountMonitoredResourcesResult struct {
	NextLink *string                     `pulumi:"nextLink"`
	Value    []MonitoredResourceResponse `pulumi:"value"`
}

func ListSubAccountMonitoredResourcesOutput(ctx *pulumi.Context, args ListSubAccountMonitoredResourcesOutputArgs, opts ...pulumi.InvokeOption) ListSubAccountMonitoredResourcesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSubAccountMonitoredResourcesResult, error) {
			args := v.(ListSubAccountMonitoredResourcesArgs)
			r, err := ListSubAccountMonitoredResources(ctx, &args, opts...)
			var s ListSubAccountMonitoredResourcesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSubAccountMonitoredResourcesResultOutput)
}

type ListSubAccountMonitoredResourcesOutputArgs struct {
	MonitorName       pulumi.StringInput `pulumi:"monitorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SubAccountName    pulumi.StringInput `pulumi:"subAccountName"`
}

func (ListSubAccountMonitoredResourcesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSubAccountMonitoredResourcesArgs)(nil)).Elem()
}


type ListSubAccountMonitoredResourcesResultOutput struct{ *pulumi.OutputState }

func (ListSubAccountMonitoredResourcesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSubAccountMonitoredResourcesResult)(nil)).Elem()
}

func (o ListSubAccountMonitoredResourcesResultOutput) ToListSubAccountMonitoredResourcesResultOutput() ListSubAccountMonitoredResourcesResultOutput {
	return o
}

func (o ListSubAccountMonitoredResourcesResultOutput) ToListSubAccountMonitoredResourcesResultOutputWithContext(ctx context.Context) ListSubAccountMonitoredResourcesResultOutput {
	return o
}

func (o ListSubAccountMonitoredResourcesResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSubAccountMonitoredResourcesResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListSubAccountMonitoredResourcesResultOutput) Value() MonitoredResourceResponseArrayOutput {
	return o.ApplyT(func(v ListSubAccountMonitoredResourcesResult) []MonitoredResourceResponse { return v.Value }).(MonitoredResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSubAccountMonitoredResourcesResultOutput{})
}
