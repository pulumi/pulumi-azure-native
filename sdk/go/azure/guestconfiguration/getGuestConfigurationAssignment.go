


package guestconfiguration

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGuestConfigurationAssignment(ctx *pulumi.Context, args *LookupGuestConfigurationAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupGuestConfigurationAssignmentResult, error) {
	var rv LookupGuestConfigurationAssignmentResult
	err := ctx.Invoke("azure-native:guestconfiguration:getGuestConfigurationAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
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
	Type       string                                         `pulumi:"type"`
}


func (val *LookupGuestConfigurationAssignmentResult) Defaults() *LookupGuestConfigurationAssignmentResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}
