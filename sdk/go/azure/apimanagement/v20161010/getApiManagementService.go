


package v20161010

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiManagementService(ctx *pulumi.Context, args *LookupApiManagementServiceArgs, opts ...pulumi.InvokeOption) (*LookupApiManagementServiceResult, error) {
	var rv LookupApiManagementServiceResult
	err := ctx.Invoke("azure-native:apimanagement/v20161010:getApiManagementService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupApiManagementServiceArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupApiManagementServiceResult struct {
	AdditionalLocations     []AdditionalRegionResponse                `pulumi:"additionalLocations"`
	AddresserEmail          *string                                   `pulumi:"addresserEmail"`
	CreatedAtUtc            string                                    `pulumi:"createdAtUtc"`
	CustomProperties        map[string]string                         `pulumi:"customProperties"`
	Etag                    string                                    `pulumi:"etag"`
	HostnameConfigurations  []HostnameConfigurationResponse           `pulumi:"hostnameConfigurations"`
	Id                      string                                    `pulumi:"id"`
	Location                string                                    `pulumi:"location"`
	ManagementApiUrl        string                                    `pulumi:"managementApiUrl"`
	Name                    *string                                   `pulumi:"name"`
	PortalUrl               string                                    `pulumi:"portalUrl"`
	ProvisioningState       string                                    `pulumi:"provisioningState"`
	PublisherEmail          string                                    `pulumi:"publisherEmail"`
	PublisherName           string                                    `pulumi:"publisherName"`
	RuntimeUrl              string                                    `pulumi:"runtimeUrl"`
	ScmUrl                  string                                    `pulumi:"scmUrl"`
	Sku                     ApiManagementServiceSkuPropertiesResponse `pulumi:"sku"`
	StaticIPs               []string                                  `pulumi:"staticIPs"`
	Tags                    map[string]string                         `pulumi:"tags"`
	TargetProvisioningState string                                    `pulumi:"targetProvisioningState"`
	Type                    string                                    `pulumi:"type"`
	VpnType                 *string                                   `pulumi:"vpnType"`
	Vpnconfiguration        *VirtualNetworkConfigurationResponse      `pulumi:"vpnconfiguration"`
}


func (val *LookupApiManagementServiceResult) Defaults() *LookupApiManagementServiceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Sku = *tmp.Sku.Defaults()

	if isZero(tmp.VpnType) {
		vpnType_ := "None"
		tmp.VpnType = &vpnType_
	}
	return &tmp
}
