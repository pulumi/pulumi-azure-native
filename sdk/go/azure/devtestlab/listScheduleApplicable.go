


package devtestlab

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListScheduleApplicable(ctx *pulumi.Context, args *ListScheduleApplicableArgs, opts ...pulumi.InvokeOption) (*ListScheduleApplicableResult, error) {
	var rv ListScheduleApplicableResult
	err := ctx.Invoke("azure-native:devtestlab:listScheduleApplicable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListScheduleApplicableArgs struct {
	LabName           string `pulumi:"labName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListScheduleApplicableResult struct {
	NextLink *string            `pulumi:"nextLink"`
	Value    []ScheduleResponse `pulumi:"value"`
}
