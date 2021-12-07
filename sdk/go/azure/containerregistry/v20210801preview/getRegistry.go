


package v20210801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRegistry(ctx *pulumi.Context, args *LookupRegistryArgs, opts ...pulumi.InvokeOption) (*LookupRegistryResult, error) {
	var rv LookupRegistryResult
	err := ctx.Invoke("azure-native:containerregistry/v20210801preview:getRegistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegistryArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupRegistryResult struct {
	AdminUserEnabled           *bool                               `pulumi:"adminUserEnabled"`
	AnonymousPullEnabled       *bool                               `pulumi:"anonymousPullEnabled"`
	CreationDate               string                              `pulumi:"creationDate"`
	DataEndpointEnabled        *bool                               `pulumi:"dataEndpointEnabled"`
	DataEndpointHostNames      []string                            `pulumi:"dataEndpointHostNames"`
	Encryption                 *EncryptionPropertyResponse         `pulumi:"encryption"`
	Id                         string                              `pulumi:"id"`
	Identity                   *IdentityPropertiesResponse         `pulumi:"identity"`
	Location                   string                              `pulumi:"location"`
	LoginServer                string                              `pulumi:"loginServer"`
	Name                       string                              `pulumi:"name"`
	NetworkRuleBypassOptions   *string                             `pulumi:"networkRuleBypassOptions"`
	NetworkRuleSet             *NetworkRuleSetResponse             `pulumi:"networkRuleSet"`
	Policies                   *PoliciesResponse                   `pulumi:"policies"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                             `pulumi:"publicNetworkAccess"`
	Sku                        SkuResponse                         `pulumi:"sku"`
	Status                     StatusResponse                      `pulumi:"status"`
	SystemData                 SystemDataResponse                  `pulumi:"systemData"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Type                       string                              `pulumi:"type"`
	ZoneRedundancy             *string                             `pulumi:"zoneRedundancy"`
}
