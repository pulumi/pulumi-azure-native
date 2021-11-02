


package v20200701preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListMonitoredResource(ctx *pulumi.Context, args *ListMonitoredResourceArgs, opts ...pulumi.InvokeOption) (*ListMonitoredResourceResult, error) {
	var rv ListMonitoredResourceResult
	err := ctx.Invoke("azure-native:elastic/v20200701preview:listMonitoredResource", args, &rv, opts...)
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
