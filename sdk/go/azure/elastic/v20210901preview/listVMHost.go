


package v20210901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListVMHost(ctx *pulumi.Context, args *ListVMHostArgs, opts ...pulumi.InvokeOption) (*ListVMHostResult, error) {
	var rv ListVMHostResult
	err := ctx.Invoke("azure-native:elastic/v20210901preview:listVMHost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListVMHostArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListVMHostResult struct {
	NextLink *string               `pulumi:"nextLink"`
	Value    []VMResourcesResponse `pulumi:"value"`
}
