


package v20201001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSubAccountMonitoredResources(ctx *pulumi.Context, args *ListSubAccountMonitoredResourcesArgs, opts ...pulumi.InvokeOption) (*ListSubAccountMonitoredResourcesResult, error) {
	var rv ListSubAccountMonitoredResourcesResult
	err := ctx.Invoke("azure-native:logz/v20201001:listSubAccountMonitoredResources", args, &rv, opts...)
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
