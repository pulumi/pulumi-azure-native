


package datadog

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListMonitorHosts(ctx *pulumi.Context, args *ListMonitorHostsArgs, opts ...pulumi.InvokeOption) (*ListMonitorHostsResult, error) {
	var rv ListMonitorHostsResult
	err := ctx.Invoke("azure-native:datadog:listMonitorHosts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMonitorHostsArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListMonitorHostsResult struct {
	NextLink *string               `pulumi:"nextLink"`
	Value    []DatadogHostResponse `pulumi:"value"`
}

func ListMonitorHostsOutput(ctx *pulumi.Context, args ListMonitorHostsOutputArgs, opts ...pulumi.InvokeOption) ListMonitorHostsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListMonitorHostsResult, error) {
			args := v.(ListMonitorHostsArgs)
			r, err := ListMonitorHosts(ctx, &args, opts...)
			var s ListMonitorHostsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListMonitorHostsResultOutput)
}

type ListMonitorHostsOutputArgs struct {
	MonitorName       pulumi.StringInput `pulumi:"monitorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListMonitorHostsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMonitorHostsArgs)(nil)).Elem()
}


type ListMonitorHostsResultOutput struct{ *pulumi.OutputState }

func (ListMonitorHostsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMonitorHostsResult)(nil)).Elem()
}

func (o ListMonitorHostsResultOutput) ToListMonitorHostsResultOutput() ListMonitorHostsResultOutput {
	return o
}

func (o ListMonitorHostsResultOutput) ToListMonitorHostsResultOutputWithContext(ctx context.Context) ListMonitorHostsResultOutput {
	return o
}

func (o ListMonitorHostsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListMonitorHostsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListMonitorHostsResultOutput) Value() DatadogHostResponseArrayOutput {
	return o.ApplyT(func(v ListMonitorHostsResult) []DatadogHostResponse { return v.Value }).(DatadogHostResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListMonitorHostsResultOutput{})
}
