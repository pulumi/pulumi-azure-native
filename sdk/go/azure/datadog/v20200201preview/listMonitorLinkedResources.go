


package v20200201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListMonitorLinkedResources(ctx *pulumi.Context, args *ListMonitorLinkedResourcesArgs, opts ...pulumi.InvokeOption) (*ListMonitorLinkedResourcesResult, error) {
	var rv ListMonitorLinkedResourcesResult
	err := ctx.Invoke("azure-native:datadog/v20200201preview:listMonitorLinkedResources", args, &rv, opts...)
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
