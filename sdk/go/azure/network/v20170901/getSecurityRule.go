


package v20170901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSecurityRule(ctx *pulumi.Context, args *LookupSecurityRuleArgs, opts ...pulumi.InvokeOption) (*LookupSecurityRuleResult, error) {
	var rv LookupSecurityRuleResult
	err := ctx.Invoke("azure-native:network/v20170901:getSecurityRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityRuleArgs struct {
	NetworkSecurityGroupName string `pulumi:"networkSecurityGroupName"`
	ResourceGroupName        string `pulumi:"resourceGroupName"`
	SecurityRuleName         string `pulumi:"securityRuleName"`
}


type LookupSecurityRuleResult struct {
	Access                               string                             `pulumi:"access"`
	Description                          *string                            `pulumi:"description"`
	DestinationAddressPrefix             *string                            `pulumi:"destinationAddressPrefix"`
	DestinationAddressPrefixes           []string                           `pulumi:"destinationAddressPrefixes"`
	DestinationApplicationSecurityGroups []ApplicationSecurityGroupResponse `pulumi:"destinationApplicationSecurityGroups"`
	DestinationPortRange                 *string                            `pulumi:"destinationPortRange"`
	DestinationPortRanges                []string                           `pulumi:"destinationPortRanges"`
	Direction                            string                             `pulumi:"direction"`
	Etag                                 *string                            `pulumi:"etag"`
	Id                                   *string                            `pulumi:"id"`
	Name                                 *string                            `pulumi:"name"`
	Priority                             *int                               `pulumi:"priority"`
	Protocol                             string                             `pulumi:"protocol"`
	ProvisioningState                    *string                            `pulumi:"provisioningState"`
	SourceAddressPrefix                  *string                            `pulumi:"sourceAddressPrefix"`
	SourceAddressPrefixes                []string                           `pulumi:"sourceAddressPrefixes"`
	SourceApplicationSecurityGroups      []ApplicationSecurityGroupResponse `pulumi:"sourceApplicationSecurityGroups"`
	SourcePortRange                      *string                            `pulumi:"sourcePortRange"`
	SourcePortRanges                     []string                           `pulumi:"sourcePortRanges"`
}
