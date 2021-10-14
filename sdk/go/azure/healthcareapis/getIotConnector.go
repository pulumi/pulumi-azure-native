


package healthcareapis

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIotConnector(ctx *pulumi.Context, args *LookupIotConnectorArgs, opts ...pulumi.InvokeOption) (*LookupIotConnectorResult, error) {
	var rv LookupIotConnectorResult
	err := ctx.Invoke("azure-native:healthcareapis:getIotConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIotConnectorArgs struct {
	IotConnectorName  string `pulumi:"iotConnectorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupIotConnectorResult struct {
	DeviceMapping                  *IotMappingPropertiesResponse                      `pulumi:"deviceMapping"`
	Etag                           *string                                            `pulumi:"etag"`
	Id                             string                                             `pulumi:"id"`
	Identity                       *ServiceManagedIdentityResponseIdentity            `pulumi:"identity"`
	IngestionEndpointConfiguration *IotEventHubIngestionEndpointConfigurationResponse `pulumi:"ingestionEndpointConfiguration"`
	Location                       *string                                            `pulumi:"location"`
	Name                           string                                             `pulumi:"name"`
	ProvisioningState              string                                             `pulumi:"provisioningState"`
	SystemData                     SystemDataResponse                                 `pulumi:"systemData"`
	Tags                           map[string]string                                  `pulumi:"tags"`
	Type                           string                                             `pulumi:"type"`
}
