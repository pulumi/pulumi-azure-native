


package v20200201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListMonitorHosts(ctx *pulumi.Context, args *ListMonitorHostsArgs, opts ...pulumi.InvokeOption) (*ListMonitorHostsResult, error) {
	var rv ListMonitorHostsResult
	err := ctx.Invoke("azure-native:datadog/v20200201preview:listMonitorHosts", args, &rv, opts...)
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
