


package analysisservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerDetails(ctx *pulumi.Context, args *LookupServerDetailsArgs, opts ...pulumi.InvokeOption) (*LookupServerDetailsResult, error) {
	var rv LookupServerDetailsResult
	err := ctx.Invoke("azure-native:analysisservices:getServerDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupServerDetailsArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupServerDetailsResult struct {
	AsAdministrators        *ServerAdministratorsResponse `pulumi:"asAdministrators"`
	BackupBlobContainerUri  *string                       `pulumi:"backupBlobContainerUri"`
	GatewayDetails          *GatewayDetailsResponse       `pulumi:"gatewayDetails"`
	Id                      string                        `pulumi:"id"`
	IpV4FirewallSettings    *IPv4FirewallSettingsResponse `pulumi:"ipV4FirewallSettings"`
	Location                string                        `pulumi:"location"`
	ManagedMode             *int                          `pulumi:"managedMode"`
	Name                    string                        `pulumi:"name"`
	ProvisioningState       string                        `pulumi:"provisioningState"`
	QuerypoolConnectionMode *string                       `pulumi:"querypoolConnectionMode"`
	ServerFullName          string                        `pulumi:"serverFullName"`
	ServerMonitorMode       *int                          `pulumi:"serverMonitorMode"`
	Sku                     ResourceSkuResponse           `pulumi:"sku"`
	State                   string                        `pulumi:"state"`
	Tags                    map[string]string             `pulumi:"tags"`
	Type                    string                        `pulumi:"type"`
}


func (val *LookupServerDetailsResult) Defaults() *LookupServerDetailsResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ManagedMode) {
		managedMode_ := 1
		tmp.ManagedMode = &managedMode_
	}
	if isZero(tmp.QuerypoolConnectionMode) {
		querypoolConnectionMode_ := "All"
		tmp.QuerypoolConnectionMode = &querypoolConnectionMode_
	}
	if isZero(tmp.ServerMonitorMode) {
		serverMonitorMode_ := 1
		tmp.ServerMonitorMode = &serverMonitorMode_
	}
	tmp.Sku = *tmp.Sku.Defaults()

	return &tmp
}
