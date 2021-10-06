


package v20201001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListMonitorMonitoredResources(ctx *pulumi.Context, args *ListMonitorMonitoredResourcesArgs, opts ...pulumi.InvokeOption) (*ListMonitorMonitoredResourcesResult, error) {
	var rv ListMonitorMonitoredResourcesResult
	err := ctx.Invoke("azure-native:logz/v20201001:listMonitorMonitoredResources", args, &rv, opts...)
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
