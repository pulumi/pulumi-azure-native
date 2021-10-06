


package v20210125

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGuestConfigurationAssignment(ctx *pulumi.Context, args *LookupGuestConfigurationAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupGuestConfigurationAssignmentResult, error) {
	var rv LookupGuestConfigurationAssignmentResult
	err := ctx.Invoke("azure-native:guestconfiguration/v20210125:getGuestConfigurationAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGuestConfigurationAssignmentArgs struct {
	GuestConfigurationAssignmentName string `pulumi:"guestConfigurationAssignmentName"`
	ResourceGroupName                string `pulumi:"resourceGroupName"`
	VmName                           string `pulumi:"vmName"`
}


type LookupGuestConfigurationAssignmentResult struct {
	Id         string                                         `pulumi:"id"`
	Location   *string                                        `pulumi:"location"`
	Name       *string                                        `pulumi:"name"`
	Properties GuestConfigurationAssignmentPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                             `pulumi:"systemData"`
	Type       string                                         `pulumi:"type"`
}
