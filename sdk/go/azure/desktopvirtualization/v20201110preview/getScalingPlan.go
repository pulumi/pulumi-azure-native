


package v20201110preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScalingPlan(ctx *pulumi.Context, args *LookupScalingPlanArgs, opts ...pulumi.InvokeOption) (*LookupScalingPlanResult, error) {
	var rv LookupScalingPlanResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20201110preview:getScalingPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScalingPlanArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ScalingPlanName   string `pulumi:"scalingPlanName"`
}


type LookupScalingPlanResult struct {
	Description        *string                            `pulumi:"description"`
	ExclusionTag       *string                            `pulumi:"exclusionTag"`
	FriendlyName       *string                            `pulumi:"friendlyName"`
	HostPoolReferences []ScalingHostPoolReferenceResponse `pulumi:"hostPoolReferences"`
	HostPoolType       *string                            `pulumi:"hostPoolType"`
	Id                 string                             `pulumi:"id"`
	Location           string                             `pulumi:"location"`
	Name               string                             `pulumi:"name"`
	Schedules          []ScalingScheduleResponse          `pulumi:"schedules"`
	Tags               map[string]string                  `pulumi:"tags"`
	TimeZone           *string                            `pulumi:"timeZone"`
	Type               string                             `pulumi:"type"`
}
