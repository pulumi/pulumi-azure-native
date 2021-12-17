


package v20210901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMaintenanceConfiguration(ctx *pulumi.Context, args *LookupMaintenanceConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupMaintenanceConfigurationResult, error) {
	var rv LookupMaintenanceConfigurationResult
	err := ctx.Invoke("azure-native:maintenance/v20210901preview:getMaintenanceConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupMaintenanceConfigurationArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupMaintenanceConfigurationResult struct {
	Duration            *string                          `pulumi:"duration"`
	ExpirationDateTime  *string                          `pulumi:"expirationDateTime"`
	ExtensionProperties map[string]string                `pulumi:"extensionProperties"`
	Id                  string                           `pulumi:"id"`
	InstallPatches      *InputPatchConfigurationResponse `pulumi:"installPatches"`
	Location            *string                          `pulumi:"location"`
	MaintenanceScope    *string                          `pulumi:"maintenanceScope"`
	Name                string                           `pulumi:"name"`
	Namespace           *string                          `pulumi:"namespace"`
	RecurEvery          *string                          `pulumi:"recurEvery"`
	StartDateTime       *string                          `pulumi:"startDateTime"`
	SystemData          SystemDataResponse               `pulumi:"systemData"`
	Tags                map[string]string                `pulumi:"tags"`
	TimeZone            *string                          `pulumi:"timeZone"`
	Type                string                           `pulumi:"type"`
	Visibility          *string                          `pulumi:"visibility"`
}


func (val *LookupMaintenanceConfigurationResult) Defaults() *LookupMaintenanceConfigurationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.InstallPatches = tmp.InstallPatches.Defaults()

	return &tmp
}
