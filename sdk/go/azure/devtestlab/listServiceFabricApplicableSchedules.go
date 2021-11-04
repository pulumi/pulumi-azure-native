


package devtestlab

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListServiceFabricApplicableSchedules(ctx *pulumi.Context, args *ListServiceFabricApplicableSchedulesArgs, opts ...pulumi.InvokeOption) (*ListServiceFabricApplicableSchedulesResult, error) {
	var rv ListServiceFabricApplicableSchedulesResult
	err := ctx.Invoke("azure-native:devtestlab:listServiceFabricApplicableSchedules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListServiceFabricApplicableSchedulesArgs struct {
	LabName           string `pulumi:"labName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	UserName          string `pulumi:"userName"`
}


type ListServiceFabricApplicableSchedulesResult struct {
	Id             string            `pulumi:"id"`
	LabVmsShutdown *ScheduleResponse `pulumi:"labVmsShutdown"`
	LabVmsStartup  *ScheduleResponse `pulumi:"labVmsStartup"`
	Location       *string           `pulumi:"location"`
	Name           string            `pulumi:"name"`
	Tags           map[string]string `pulumi:"tags"`
	Type           string            `pulumi:"type"`
}
