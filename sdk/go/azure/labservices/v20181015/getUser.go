


package v20181015

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	var rv LookupUserResult
	err := ctx.Invoke("azure-native:labservices/v20181015:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserArgs struct {
	Expand            *string `pulumi:"expand"`
	LabAccountName    string  `pulumi:"labAccountName"`
	LabName           string  `pulumi:"labName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	UserName          string  `pulumi:"userName"`
}


type LookupUserResult struct {
	Email                 string                        `pulumi:"email"`
	FamilyName            string                        `pulumi:"familyName"`
	GivenName             string                        `pulumi:"givenName"`
	Id                    string                        `pulumi:"id"`
	LatestOperationResult LatestOperationResultResponse `pulumi:"latestOperationResult"`
	Location              *string                       `pulumi:"location"`
	Name                  string                        `pulumi:"name"`
	ProvisioningState     *string                       `pulumi:"provisioningState"`
	Tags                  map[string]string             `pulumi:"tags"`
	TenantId              string                        `pulumi:"tenantId"`
	TotalUsage            string                        `pulumi:"totalUsage"`
	Type                  string                        `pulumi:"type"`
	UniqueIdentifier      *string                       `pulumi:"uniqueIdentifier"`
}
