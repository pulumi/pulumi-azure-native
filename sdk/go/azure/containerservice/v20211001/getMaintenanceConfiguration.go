


package v20211001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMaintenanceConfiguration(ctx *pulumi.Context, args *LookupMaintenanceConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupMaintenanceConfigurationResult, error) {
	var rv LookupMaintenanceConfigurationResult
	err := ctx.Invoke("azure-native:containerservice/v20211001:getMaintenanceConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMaintenanceConfigurationArgs struct {
	ConfigName        string `pulumi:"configName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupMaintenanceConfigurationResult struct {
	Id             string               `pulumi:"id"`
	Name           string               `pulumi:"name"`
	NotAllowedTime []TimeSpanResponse   `pulumi:"notAllowedTime"`
	SystemData     SystemDataResponse   `pulumi:"systemData"`
	TimeInWeek     []TimeInWeekResponse `pulumi:"timeInWeek"`
	Type           string               `pulumi:"type"`
}
