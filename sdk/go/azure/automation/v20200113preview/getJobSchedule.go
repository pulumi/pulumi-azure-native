


package v20200113preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJobSchedule(ctx *pulumi.Context, args *LookupJobScheduleArgs, opts ...pulumi.InvokeOption) (*LookupJobScheduleResult, error) {
	var rv LookupJobScheduleResult
	err := ctx.Invoke("azure-native:automation/v20200113preview:getJobSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobScheduleArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	JobScheduleId         string `pulumi:"jobScheduleId"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupJobScheduleResult struct {
	Id            string                               `pulumi:"id"`
	JobScheduleId *string                              `pulumi:"jobScheduleId"`
	Name          string                               `pulumi:"name"`
	Parameters    map[string]string                    `pulumi:"parameters"`
	RunOn         *string                              `pulumi:"runOn"`
	Runbook       *RunbookAssociationPropertyResponse  `pulumi:"runbook"`
	Schedule      *ScheduleAssociationPropertyResponse `pulumi:"schedule"`
	Type          string                               `pulumi:"type"`
}
