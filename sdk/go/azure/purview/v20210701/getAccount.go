


package v20210701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("azure-native:purview/v20210701:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAccountResult struct {
	CloudConnectors            *CloudConnectorsResponse                  `pulumi:"cloudConnectors"`
	CreatedAt                  string                                    `pulumi:"createdAt"`
	CreatedBy                  string                                    `pulumi:"createdBy"`
	CreatedByObjectId          string                                    `pulumi:"createdByObjectId"`
	Endpoints                  AccountPropertiesResponseEndpoints        `pulumi:"endpoints"`
	FriendlyName               string                                    `pulumi:"friendlyName"`
	Id                         string                                    `pulumi:"id"`
	Identity                   *IdentityResponse                         `pulumi:"identity"`
	Location                   *string                                   `pulumi:"location"`
	ManagedResourceGroupName   *string                                   `pulumi:"managedResourceGroupName"`
	ManagedResources           AccountPropertiesResponseManagedResources `pulumi:"managedResources"`
	Name                       string                                    `pulumi:"name"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse       `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                                    `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                                   `pulumi:"publicNetworkAccess"`
	Sku                        AccountResponseSku                        `pulumi:"sku"`
	SystemData                 TrackedResourceResponseSystemData         `pulumi:"systemData"`
	Tags                       map[string]string                         `pulumi:"tags"`
	Type                       string                                    `pulumi:"type"`
}
