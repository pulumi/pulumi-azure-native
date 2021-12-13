


package logz

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSubAccountVMHosts(ctx *pulumi.Context, args *ListSubAccountVMHostsArgs, opts ...pulumi.InvokeOption) (*ListSubAccountVMHostsResult, error) {
	var rv ListSubAccountVMHostsResult
	err := ctx.Invoke("azure-native:logz:listSubAccountVMHosts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSubAccountVMHostsArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SubAccountName    string `pulumi:"subAccountName"`
}


type ListSubAccountVMHostsResult struct {
	NextLink *string               `pulumi:"nextLink"`
	Value    []VMResourcesResponse `pulumi:"value"`
}
