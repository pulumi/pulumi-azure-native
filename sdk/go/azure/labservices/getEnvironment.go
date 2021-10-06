


package labservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEnvironment(ctx *pulumi.Context, args *LookupEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentResult, error) {
	var rv LookupEnvironmentResult
	err := ctx.Invoke("azure-native:labservices:getEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentArgs struct {
	EnvironmentName        string  `pulumi:"environmentName"`
	EnvironmentSettingName string  `pulumi:"environmentSettingName"`
	Expand                 *string `pulumi:"expand"`
	LabAccountName         string  `pulumi:"labAccountName"`
	LabName                string  `pulumi:"labName"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
}


type LookupEnvironmentResult struct {
	ClaimedByUserName        string                        `pulumi:"claimedByUserName"`
	ClaimedByUserObjectId    string                        `pulumi:"claimedByUserObjectId"`
	ClaimedByUserPrincipalId string                        `pulumi:"claimedByUserPrincipalId"`
	Id                       string                        `pulumi:"id"`
	IsClaimed                bool                          `pulumi:"isClaimed"`
	LastKnownPowerState      string                        `pulumi:"lastKnownPowerState"`
	LatestOperationResult    LatestOperationResultResponse `pulumi:"latestOperationResult"`
	Location                 *string                       `pulumi:"location"`
	Name                     string                        `pulumi:"name"`
	NetworkInterface         NetworkInterfaceResponse      `pulumi:"networkInterface"`
	PasswordLastReset        string                        `pulumi:"passwordLastReset"`
	ProvisioningState        *string                       `pulumi:"provisioningState"`
	ResourceSets             *ResourceSetResponse          `pulumi:"resourceSets"`
	Tags                     map[string]string             `pulumi:"tags"`
	TotalUsage               string                        `pulumi:"totalUsage"`
	Type                     string                        `pulumi:"type"`
	UniqueIdentifier         *string                       `pulumi:"uniqueIdentifier"`
}
