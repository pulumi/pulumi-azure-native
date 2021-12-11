


package v20211001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDeploymentInfo(ctx *pulumi.Context, args *ListDeploymentInfoArgs, opts ...pulumi.InvokeOption) (*ListDeploymentInfoResult, error) {
	var rv ListDeploymentInfoResult
	err := ctx.Invoke("azure-native:elastic/v20211001preview:listDeploymentInfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDeploymentInfoArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListDeploymentInfoResult struct {
	DiskCapacity   string `pulumi:"diskCapacity"`
	MemoryCapacity string `pulumi:"memoryCapacity"`
	Status         string `pulumi:"status"`
	Version        string `pulumi:"version"`
}
