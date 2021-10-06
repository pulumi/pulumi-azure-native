


package v20170201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPatchSchedule(ctx *pulumi.Context, args *LookupPatchScheduleArgs, opts ...pulumi.InvokeOption) (*LookupPatchScheduleResult, error) {
	var rv LookupPatchScheduleResult
	err := ctx.Invoke("azure-native:cache/v20170201:getPatchSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPatchScheduleArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPatchScheduleResult struct {
	Id              string                  `pulumi:"id"`
	Location        string                  `pulumi:"location"`
	Name            string                  `pulumi:"name"`
	ScheduleEntries []ScheduleEntryResponse `pulumi:"scheduleEntries"`
	Type            string                  `pulumi:"type"`
}
