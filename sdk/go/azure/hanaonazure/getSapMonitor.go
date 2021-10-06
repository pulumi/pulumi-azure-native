


package hanaonazure

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSapMonitor(ctx *pulumi.Context, args *LookupSapMonitorArgs, opts ...pulumi.InvokeOption) (*LookupSapMonitorResult, error) {
	var rv LookupSapMonitorResult
	err := ctx.Invoke("azure-native:hanaonazure:getSapMonitor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSapMonitorArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SapMonitorName    string `pulumi:"sapMonitorName"`
}


type LookupSapMonitorResult struct {
	EnableCustomerAnalytics        *bool             `pulumi:"enableCustomerAnalytics"`
	Id                             string            `pulumi:"id"`
	Location                       string            `pulumi:"location"`
	LogAnalyticsWorkspaceArmId     *string           `pulumi:"logAnalyticsWorkspaceArmId"`
	LogAnalyticsWorkspaceId        *string           `pulumi:"logAnalyticsWorkspaceId"`
	LogAnalyticsWorkspaceSharedKey *string           `pulumi:"logAnalyticsWorkspaceSharedKey"`
	ManagedResourceGroupName       string            `pulumi:"managedResourceGroupName"`
	MonitorSubnet                  *string           `pulumi:"monitorSubnet"`
	Name                           string            `pulumi:"name"`
	ProvisioningState              string            `pulumi:"provisioningState"`
	SapMonitorCollectorVersion     string            `pulumi:"sapMonitorCollectorVersion"`
	Tags                           map[string]string `pulumi:"tags"`
	Type                           string            `pulumi:"type"`
}
