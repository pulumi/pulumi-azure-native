


package automanage

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfigurationProfileAssignment(ctx *pulumi.Context, args *LookupConfigurationProfileAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationProfileAssignmentResult, error) {
	var rv LookupConfigurationProfileAssignmentResult
	err := ctx.Invoke("azure-native:automanage:getConfigurationProfileAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationProfileAssignmentArgs struct {
	ConfigurationProfileAssignmentName string `pulumi:"configurationProfileAssignmentName"`
	ResourceGroupName                  string `pulumi:"resourceGroupName"`
	VmName                             string `pulumi:"vmName"`
}


type LookupConfigurationProfileAssignmentResult struct {
	Id         string                                           `pulumi:"id"`
	Name       string                                           `pulumi:"name"`
	Properties ConfigurationProfileAssignmentPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                               `pulumi:"systemData"`
	Type       string                                           `pulumi:"type"`
}
