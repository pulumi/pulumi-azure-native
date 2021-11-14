


package v20180801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAzureFirewall(ctx *pulumi.Context, args *LookupAzureFirewallArgs, opts ...pulumi.InvokeOption) (*LookupAzureFirewallResult, error) {
	var rv LookupAzureFirewallResult
	err := ctx.Invoke("azure-native:network/v20180801:getAzureFirewall", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAzureFirewallArgs struct {
	AzureFirewallName string `pulumi:"azureFirewallName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAzureFirewallResult struct {
	ApplicationRuleCollections []AzureFirewallApplicationRuleCollectionResponse `pulumi:"applicationRuleCollections"`
	Etag                       string                                           `pulumi:"etag"`
	Id                         *string                                          `pulumi:"id"`
	IpConfigurations           []AzureFirewallIPConfigurationResponse           `pulumi:"ipConfigurations"`
	Location                   *string                                          `pulumi:"location"`
	Name                       string                                           `pulumi:"name"`
	NatRuleCollections         []AzureFirewallNatRuleCollectionResponse         `pulumi:"natRuleCollections"`
	NetworkRuleCollections     []AzureFirewallNetworkRuleCollectionResponse     `pulumi:"networkRuleCollections"`
	ProvisioningState          string                                           `pulumi:"provisioningState"`
	Tags                       map[string]string                                `pulumi:"tags"`
	Type                       string                                           `pulumi:"type"`
}
