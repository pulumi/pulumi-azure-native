


package labservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLab(ctx *pulumi.Context, args *LookupLabArgs, opts ...pulumi.InvokeOption) (*LookupLabResult, error) {
	var rv LookupLabResult
	err := ctx.Invoke("azure-native:labservices:getLab", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLabArgs struct {
	Expand            *string `pulumi:"expand"`
	LabAccountName    string  `pulumi:"labAccountName"`
	LabName           string  `pulumi:"labName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupLabResult struct {
	CreatedByObjectId          string                        `pulumi:"createdByObjectId"`
	CreatedByUserPrincipalName string                        `pulumi:"createdByUserPrincipalName"`
	CreatedDate                string                        `pulumi:"createdDate"`
	Id                         string                        `pulumi:"id"`
	InvitationCode             string                        `pulumi:"invitationCode"`
	LatestOperationResult      LatestOperationResultResponse `pulumi:"latestOperationResult"`
	Location                   *string                       `pulumi:"location"`
	MaxUsersInLab              *int                          `pulumi:"maxUsersInLab"`
	Name                       string                        `pulumi:"name"`
	ProvisioningState          *string                       `pulumi:"provisioningState"`
	Tags                       map[string]string             `pulumi:"tags"`
	Type                       string                        `pulumi:"type"`
	UniqueIdentifier           *string                       `pulumi:"uniqueIdentifier"`
	UsageQuota                 *string                       `pulumi:"usageQuota"`
	UserAccessMode             *string                       `pulumi:"userAccessMode"`
	UserQuota                  int                           `pulumi:"userQuota"`
}
