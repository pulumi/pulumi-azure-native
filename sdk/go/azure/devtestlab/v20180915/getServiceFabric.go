


package v20180915

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServiceFabric(ctx *pulumi.Context, args *LookupServiceFabricArgs, opts ...pulumi.InvokeOption) (*LookupServiceFabricResult, error) {
	var rv LookupServiceFabricResult
	err := ctx.Invoke("azure-native:devtestlab/v20180915:getServiceFabric", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceFabricArgs struct {
	Expand            *string `pulumi:"expand"`
	LabName           string  `pulumi:"labName"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	UserName          string  `pulumi:"userName"`
}


type LookupServiceFabricResult struct {
	ApplicableSchedule      ApplicableScheduleResponse `pulumi:"applicableSchedule"`
	EnvironmentId           *string                    `pulumi:"environmentId"`
	ExternalServiceFabricId *string                    `pulumi:"externalServiceFabricId"`
	Id                      string                     `pulumi:"id"`
	Location                *string                    `pulumi:"location"`
	Name                    string                     `pulumi:"name"`
	ProvisioningState       string                     `pulumi:"provisioningState"`
	Tags                    map[string]string          `pulumi:"tags"`
	Type                    string                     `pulumi:"type"`
	UniqueIdentifier        string                     `pulumi:"uniqueIdentifier"`
}
