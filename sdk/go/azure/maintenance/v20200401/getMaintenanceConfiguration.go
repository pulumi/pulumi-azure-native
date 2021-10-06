


package v20200401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMaintenanceConfiguration(ctx *pulumi.Context, args *LookupMaintenanceConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupMaintenanceConfigurationResult, error) {
	var rv LookupMaintenanceConfigurationResult
	err := ctx.Invoke("azure-native:maintenance/v20200401:getMaintenanceConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMaintenanceConfigurationArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupMaintenanceConfigurationResult struct {
	ExtensionProperties map[string]string `pulumi:"extensionProperties"`
	Id                  string            `pulumi:"id"`
	Location            *string           `pulumi:"location"`
	MaintenanceScope    *string           `pulumi:"maintenanceScope"`
	Name                string            `pulumi:"name"`
	Namespace           *string           `pulumi:"namespace"`
	Tags                map[string]string `pulumi:"tags"`
	Type                string            `pulumi:"type"`
}
