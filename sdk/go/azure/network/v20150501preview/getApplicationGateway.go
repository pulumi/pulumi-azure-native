


package v20150501preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplicationGateway(ctx *pulumi.Context, args *LookupApplicationGatewayArgs, opts ...pulumi.InvokeOption) (*LookupApplicationGatewayResult, error) {
	var rv LookupApplicationGatewayResult
	err := ctx.Invoke("azure-native:network/v20150501preview:getApplicationGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationGatewayArgs struct {
	ApplicationGatewayName string `pulumi:"applicationGatewayName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupApplicationGatewayResult struct {
	BackendAddressPools           []ApplicationGatewayBackendAddressPoolResponse      `pulumi:"backendAddressPools"`
	BackendHttpSettingsCollection []ApplicationGatewayBackendHttpSettingsResponse     `pulumi:"backendHttpSettingsCollection"`
	Etag                          *string                                             `pulumi:"etag"`
	FrontendIPConfigurations      []ApplicationGatewayFrontendIPConfigurationResponse `pulumi:"frontendIPConfigurations"`
	FrontendPorts                 []ApplicationGatewayFrontendPortResponse            `pulumi:"frontendPorts"`
	GatewayIPConfigurations       []ApplicationGatewayIPConfigurationResponse         `pulumi:"gatewayIPConfigurations"`
	HttpListeners                 []ApplicationGatewayHttpListenerResponse            `pulumi:"httpListeners"`
	Id                            string                                              `pulumi:"id"`
	Location                      string                                              `pulumi:"location"`
	Name                          string                                              `pulumi:"name"`
	OperationalState              string                                              `pulumi:"operationalState"`
	ProvisioningState             *string                                             `pulumi:"provisioningState"`
	RequestRoutingRules           []ApplicationGatewayRequestRoutingRuleResponse      `pulumi:"requestRoutingRules"`
	ResourceGuid                  *string                                             `pulumi:"resourceGuid"`
	Sku                           *ApplicationGatewaySkuResponse                      `pulumi:"sku"`
	SslCertificates               []ApplicationGatewaySslCertificateResponse          `pulumi:"sslCertificates"`
	Tags                          map[string]string                                   `pulumi:"tags"`
	Type                          string                                              `pulumi:"type"`
}
