


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfigurationStore(ctx *pulumi.Context, args *LookupConfigurationStoreArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationStoreResult, error) {
	var rv LookupConfigurationStoreResult
	err := ctx.Invoke("azure-native:appconfiguration/v20210301preview:getConfigurationStore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationStoreArgs struct {
	ConfigStoreName   string `pulumi:"configStoreName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupConfigurationStoreResult struct {
	CreationDate               string                                       `pulumi:"creationDate"`
	DisableLocalAuth           *bool                                        `pulumi:"disableLocalAuth"`
	Encryption                 *EncryptionPropertiesResponse                `pulumi:"encryption"`
	Endpoint                   string                                       `pulumi:"endpoint"`
	Id                         string                                       `pulumi:"id"`
	Identity                   *ResourceIdentityResponse                    `pulumi:"identity"`
	Location                   string                                       `pulumi:"location"`
	Name                       string                                       `pulumi:"name"`
	PrivateEndpointConnections []PrivateEndpointConnectionReferenceResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                                       `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                                      `pulumi:"publicNetworkAccess"`
	Sku                        SkuResponse                                  `pulumi:"sku"`
	SystemData                 SystemDataResponse                           `pulumi:"systemData"`
	Tags                       map[string]string                            `pulumi:"tags"`
	Type                       string                                       `pulumi:"type"`
}
