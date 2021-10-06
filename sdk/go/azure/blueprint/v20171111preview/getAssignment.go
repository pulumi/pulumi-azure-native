


package v20171111preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAssignment(ctx *pulumi.Context, args *LookupAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupAssignmentResult, error) {
	var rv LookupAssignmentResult
	err := ctx.Invoke("azure-native:blueprint/v20171111preview:getAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssignmentArgs struct {
	AssignmentName string  `pulumi:"assignmentName"`
	SubscriptionId *string `pulumi:"subscriptionId"`
}


type LookupAssignmentResult struct {
	BlueprintId       *string                               `pulumi:"blueprintId"`
	Description       *string                               `pulumi:"description"`
	DisplayName       *string                               `pulumi:"displayName"`
	Id                string                                `pulumi:"id"`
	Identity          ManagedServiceIdentityResponse        `pulumi:"identity"`
	Location          string                                `pulumi:"location"`
	Locks             *AssignmentLockSettingsResponse       `pulumi:"locks"`
	Name              string                                `pulumi:"name"`
	Parameters        map[string]ParameterValueBaseResponse `pulumi:"parameters"`
	ProvisioningState string                                `pulumi:"provisioningState"`
	ResourceGroups    map[string]ResourceGroupValueResponse `pulumi:"resourceGroups"`
	Status            AssignmentStatusResponse              `pulumi:"status"`
	Type              string                                `pulumi:"type"`
}
