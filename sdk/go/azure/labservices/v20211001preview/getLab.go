


package v20211001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLab(ctx *pulumi.Context, args *LookupLabArgs, opts ...pulumi.InvokeOption) (*LookupLabResult, error) {
	var rv LookupLabResult
	err := ctx.Invoke("azure-native:labservices/v20211001preview:getLab", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLabArgs struct {
	LabName           string `pulumi:"labName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupLabResult struct {
	AutoShutdownProfile   AutoShutdownProfileResponse   `pulumi:"autoShutdownProfile"`
	ConnectionProfile     ConnectionProfileResponse     `pulumi:"connectionProfile"`
	Description           *string                       `pulumi:"description"`
	Id                    string                        `pulumi:"id"`
	LabPlanId             *string                       `pulumi:"labPlanId"`
	Location              string                        `pulumi:"location"`
	Name                  string                        `pulumi:"name"`
	NetworkProfile        *LabNetworkProfileResponse    `pulumi:"networkProfile"`
	ProvisioningState     string                        `pulumi:"provisioningState"`
	RosterProfile         *RosterProfileResponse        `pulumi:"rosterProfile"`
	SecurityProfile       SecurityProfileResponse       `pulumi:"securityProfile"`
	State                 string                        `pulumi:"state"`
	SystemData            SystemDataResponse            `pulumi:"systemData"`
	Tags                  map[string]string             `pulumi:"tags"`
	Title                 *string                       `pulumi:"title"`
	Type                  string                        `pulumi:"type"`
	VirtualMachineProfile VirtualMachineProfileResponse `pulumi:"virtualMachineProfile"`
}
