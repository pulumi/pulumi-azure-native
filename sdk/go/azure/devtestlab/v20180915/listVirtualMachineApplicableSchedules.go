


package v20180915

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListVirtualMachineApplicableSchedules(ctx *pulumi.Context, args *ListVirtualMachineApplicableSchedulesArgs, opts ...pulumi.InvokeOption) (*ListVirtualMachineApplicableSchedulesResult, error) {
	var rv ListVirtualMachineApplicableSchedulesResult
	err := ctx.Invoke("azure-native:devtestlab/v20180915:listVirtualMachineApplicableSchedules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type ListVirtualMachineApplicableSchedulesArgs struct {
	LabName           string `pulumi:"labName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListVirtualMachineApplicableSchedulesResult struct {
	Id             string            `pulumi:"id"`
	LabVmsShutdown *ScheduleResponse `pulumi:"labVmsShutdown"`
	LabVmsStartup  *ScheduleResponse `pulumi:"labVmsStartup"`
	Location       *string           `pulumi:"location"`
	Name           string            `pulumi:"name"`
	Tags           map[string]string `pulumi:"tags"`
	Type           string            `pulumi:"type"`
}


func (val *ListVirtualMachineApplicableSchedulesResult) Defaults() *ListVirtualMachineApplicableSchedulesResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.LabVmsShutdown = tmp.LabVmsShutdown.Defaults()

	tmp.LabVmsStartup = tmp.LabVmsStartup.Defaults()

	return &tmp
}
