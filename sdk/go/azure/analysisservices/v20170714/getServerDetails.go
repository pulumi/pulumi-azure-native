


package v20170714

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerDetails(ctx *pulumi.Context, args *LookupServerDetailsArgs, opts ...pulumi.InvokeOption) (*LookupServerDetailsResult, error) {
	var rv LookupServerDetailsResult
	err := ctx.Invoke("azure-native:analysisservices/v20170714:getServerDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerDetailsArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupServerDetailsResult struct {
	AsAdministrators       *ServerAdministratorsResponse `pulumi:"asAdministrators"`
	BackupBlobContainerUri *string                       `pulumi:"backupBlobContainerUri"`
	GatewayDetails         *GatewayDetailsResponse       `pulumi:"gatewayDetails"`
	Id                     string                        `pulumi:"id"`
	Location               string                        `pulumi:"location"`
	ManagedMode            *int                          `pulumi:"managedMode"`
	Name                   string                        `pulumi:"name"`
	ProvisioningState      string                        `pulumi:"provisioningState"`
	ServerFullName         string                        `pulumi:"serverFullName"`
	ServerMonitorMode      *int                          `pulumi:"serverMonitorMode"`
	Sku                    ResourceSkuResponse           `pulumi:"sku"`
	State                  string                        `pulumi:"state"`
	Tags                   map[string]string             `pulumi:"tags"`
	Type                   string                        `pulumi:"type"`
}
