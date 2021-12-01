


package v20211001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	var rv LookupUserResult
	err := ctx.Invoke("azure-native:labservices/v20211001preview:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserArgs struct {
	LabName           string `pulumi:"labName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	UserName          string `pulumi:"userName"`
}


type LookupUserResult struct {
	AdditionalUsageQuota *string            `pulumi:"additionalUsageQuota"`
	DisplayName          string             `pulumi:"displayName"`
	Email                string             `pulumi:"email"`
	Id                   string             `pulumi:"id"`
	InvitationSent       string             `pulumi:"invitationSent"`
	InvitationState      string             `pulumi:"invitationState"`
	Name                 string             `pulumi:"name"`
	ProvisioningState    string             `pulumi:"provisioningState"`
	RegistrationState    string             `pulumi:"registrationState"`
	SystemData           SystemDataResponse `pulumi:"systemData"`
	TotalUsage           string             `pulumi:"totalUsage"`
	Type                 string             `pulumi:"type"`
}
