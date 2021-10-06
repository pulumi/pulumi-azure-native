


package v20201001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListMonitorVMHosts(ctx *pulumi.Context, args *ListMonitorVMHostsArgs, opts ...pulumi.InvokeOption) (*ListMonitorVMHostsResult, error) {
	var rv ListMonitorVMHostsResult
	err := ctx.Invoke("azure-native:logz/v20201001:listMonitorVMHosts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMonitorVMHostsArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListMonitorVMHostsResult struct {
	NextLink *string               `pulumi:"nextLink"`
	Value    []VMResourcesResponse `pulumi:"value"`
}
