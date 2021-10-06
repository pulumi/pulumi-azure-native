


package maintenance

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfigurationAssignment(ctx *pulumi.Context, args *LookupConfigurationAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationAssignmentResult, error) {
	var rv LookupConfigurationAssignmentResult
	err := ctx.Invoke("azure-native:maintenance:getConfigurationAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationAssignmentArgs struct {
	ConfigurationAssignmentName string `pulumi:"configurationAssignmentName"`
	ProviderName                string `pulumi:"providerName"`
	ResourceGroupName           string `pulumi:"resourceGroupName"`
	ResourceName                string `pulumi:"resourceName"`
	ResourceType                string `pulumi:"resourceType"`
}


type LookupConfigurationAssignmentResult struct {
	Id                         string             `pulumi:"id"`
	Location                   *string            `pulumi:"location"`
	MaintenanceConfigurationId *string            `pulumi:"maintenanceConfigurationId"`
	Name                       string             `pulumi:"name"`
	ResourceId                 *string            `pulumi:"resourceId"`
	SystemData                 SystemDataResponse `pulumi:"systemData"`
	Type                       string             `pulumi:"type"`
}
