


package v20210901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfigurationAssignmentParent(ctx *pulumi.Context, args *LookupConfigurationAssignmentParentArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationAssignmentParentResult, error) {
	var rv LookupConfigurationAssignmentParentResult
	err := ctx.Invoke("azure-native:maintenance/v20210901preview:getConfigurationAssignmentParent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationAssignmentParentArgs struct {
	ConfigurationAssignmentName string `pulumi:"configurationAssignmentName"`
	ProviderName                string `pulumi:"providerName"`
	ResourceGroupName           string `pulumi:"resourceGroupName"`
	ResourceName                string `pulumi:"resourceName"`
	ResourceParentName          string `pulumi:"resourceParentName"`
	ResourceParentType          string `pulumi:"resourceParentType"`
	ResourceType                string `pulumi:"resourceType"`
}


type LookupConfigurationAssignmentParentResult struct {
	Id                         string             `pulumi:"id"`
	Location                   *string            `pulumi:"location"`
	MaintenanceConfigurationId *string            `pulumi:"maintenanceConfigurationId"`
	Name                       string             `pulumi:"name"`
	ResourceId                 *string            `pulumi:"resourceId"`
	SystemData                 SystemDataResponse `pulumi:"systemData"`
	Type                       string             `pulumi:"type"`
}
