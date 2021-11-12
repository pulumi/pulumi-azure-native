


package managednetwork

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScopeAssignment(ctx *pulumi.Context, args *LookupScopeAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupScopeAssignmentResult, error) {
	var rv LookupScopeAssignmentResult
	err := ctx.Invoke("azure-native:managednetwork:getScopeAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScopeAssignmentArgs struct {
	Scope               string `pulumi:"scope"`
	ScopeAssignmentName string `pulumi:"scopeAssignmentName"`
}


type LookupScopeAssignmentResult struct {
	AssignedManagedNetwork *string `pulumi:"assignedManagedNetwork"`
	Etag                   string  `pulumi:"etag"`
	Id                     string  `pulumi:"id"`
	Location               *string `pulumi:"location"`
	Name                   string  `pulumi:"name"`
	ProvisioningState      string  `pulumi:"provisioningState"`
	Type                   string  `pulumi:"type"`
}
