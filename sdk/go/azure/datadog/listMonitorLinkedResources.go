


package datadog

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListMonitorLinkedResources(ctx *pulumi.Context, args *ListMonitorLinkedResourcesArgs, opts ...pulumi.InvokeOption) (*ListMonitorLinkedResourcesResult, error) {
	var rv ListMonitorLinkedResourcesResult
	err := ctx.Invoke("azure-native:datadog:listMonitorLinkedResources", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMonitorLinkedResourcesArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListMonitorLinkedResourcesResult struct {
	NextLink *string                  `pulumi:"nextLink"`
	Value    []LinkedResourceResponse `pulumi:"value"`
}

func ListMonitorLinkedResourcesOutput(ctx *pulumi.Context, args ListMonitorLinkedResourcesOutputArgs, opts ...pulumi.InvokeOption) ListMonitorLinkedResourcesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListMonitorLinkedResourcesResult, error) {
			args := v.(ListMonitorLinkedResourcesArgs)
			r, err := ListMonitorLinkedResources(ctx, &args, opts...)
			var s ListMonitorLinkedResourcesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListMonitorLinkedResourcesResultOutput)
}

type ListMonitorLinkedResourcesOutputArgs struct {
	MonitorName       pulumi.StringInput `pulumi:"monitorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListMonitorLinkedResourcesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMonitorLinkedResourcesArgs)(nil)).Elem()
}


type ListMonitorLinkedResourcesResultOutput struct{ *pulumi.OutputState }

func (ListMonitorLinkedResourcesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMonitorLinkedResourcesResult)(nil)).Elem()
}

func (o ListMonitorLinkedResourcesResultOutput) ToListMonitorLinkedResourcesResultOutput() ListMonitorLinkedResourcesResultOutput {
	return o
}

func (o ListMonitorLinkedResourcesResultOutput) ToListMonitorLinkedResourcesResultOutputWithContext(ctx context.Context) ListMonitorLinkedResourcesResultOutput {
	return o
}

func (o ListMonitorLinkedResourcesResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListMonitorLinkedResourcesResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListMonitorLinkedResourcesResultOutput) Value() LinkedResourceResponseArrayOutput {
	return o.ApplyT(func(v ListMonitorLinkedResourcesResult) []LinkedResourceResponse { return v.Value }).(LinkedResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListMonitorLinkedResourcesResultOutput{})
}
