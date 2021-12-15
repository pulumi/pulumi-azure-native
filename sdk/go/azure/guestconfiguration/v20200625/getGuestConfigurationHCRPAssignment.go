


package v20200625

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGuestConfigurationHCRPAssignment(ctx *pulumi.Context, args *LookupGuestConfigurationHCRPAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupGuestConfigurationHCRPAssignmentResult, error) {
	var rv LookupGuestConfigurationHCRPAssignmentResult
	err := ctx.Invoke("azure-native:guestconfiguration/v20200625:getGuestConfigurationHCRPAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupGuestConfigurationHCRPAssignmentArgs struct {
	GuestConfigurationAssignmentName string `pulumi:"guestConfigurationAssignmentName"`
	MachineName                      string `pulumi:"machineName"`
	ResourceGroupName                string `pulumi:"resourceGroupName"`
}


type LookupGuestConfigurationHCRPAssignmentResult struct {
	Id         string                                         `pulumi:"id"`
	Location   *string                                        `pulumi:"location"`
	Name       *string                                        `pulumi:"name"`
	Properties GuestConfigurationAssignmentPropertiesResponse `pulumi:"properties"`
	Type       string                                         `pulumi:"type"`
}


func (val *LookupGuestConfigurationHCRPAssignmentResult) Defaults() *LookupGuestConfigurationHCRPAssignmentResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}
