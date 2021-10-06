


package v20200601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiManagementService(ctx *pulumi.Context, args *LookupApiManagementServiceArgs, opts ...pulumi.InvokeOption) (*LookupApiManagementServiceResult, error) {
	var rv LookupApiManagementServiceResult
	err := ctx.Invoke("azure-native:apimanagement/v20200601preview:getApiManagementService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiManagementServiceArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupApiManagementServiceResult struct {
	AdditionalLocations         []AdditionalLocationResponse              `pulumi:"additionalLocations"`
	ApiVersionConstraint        *ApiVersionConstraintResponse             `pulumi:"apiVersionConstraint"`
	Certificates                []CertificateConfigurationResponse        `pulumi:"certificates"`
	CreatedAtUtc                string                                    `pulumi:"createdAtUtc"`
	CustomProperties            map[string]string                         `pulumi:"customProperties"`
	DeveloperPortalUrl          string                                    `pulumi:"developerPortalUrl"`
	DisableGateway              *bool                                     `pulumi:"disableGateway"`
	EnableClientCertificate     *bool                                     `pulumi:"enableClientCertificate"`
	Etag                        string                                    `pulumi:"etag"`
	GatewayRegionalUrl          string                                    `pulumi:"gatewayRegionalUrl"`
	GatewayUrl                  string                                    `pulumi:"gatewayUrl"`
	HostnameConfigurations      []HostnameConfigurationResponse           `pulumi:"hostnameConfigurations"`
	Id                          string                                    `pulumi:"id"`
	Identity                    *ApiManagementServiceIdentityResponse     `pulumi:"identity"`
	Location                    string                                    `pulumi:"location"`
	ManagementApiUrl            string                                    `pulumi:"managementApiUrl"`
	Name                        string                                    `pulumi:"name"`
	NotificationSenderEmail     *string                                   `pulumi:"notificationSenderEmail"`
	PortalUrl                   string                                    `pulumi:"portalUrl"`
	PrivateIPAddresses          []string                                  `pulumi:"privateIPAddresses"`
	ProvisioningState           string                                    `pulumi:"provisioningState"`
	PublicIPAddresses           []string                                  `pulumi:"publicIPAddresses"`
	PublisherEmail              string                                    `pulumi:"publisherEmail"`
	PublisherName               string                                    `pulumi:"publisherName"`
	Restore                     *bool                                     `pulumi:"restore"`
	ScmUrl                      string                                    `pulumi:"scmUrl"`
	Sku                         ApiManagementServiceSkuPropertiesResponse `pulumi:"sku"`
	Tags                        map[string]string                         `pulumi:"tags"`
	TargetProvisioningState     string                                    `pulumi:"targetProvisioningState"`
	Type                        string                                    `pulumi:"type"`
	VirtualNetworkConfiguration *VirtualNetworkConfigurationResponse      `pulumi:"virtualNetworkConfiguration"`
	VirtualNetworkType          *string                                   `pulumi:"virtualNetworkType"`
	Zones                       []string                                  `pulumi:"zones"`
}
