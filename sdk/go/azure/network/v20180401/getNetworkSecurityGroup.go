


package v20180401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkSecurityGroup(ctx *pulumi.Context, args *LookupNetworkSecurityGroupArgs, opts ...pulumi.InvokeOption) (*LookupNetworkSecurityGroupResult, error) {
	var rv LookupNetworkSecurityGroupResult
	err := ctx.Invoke("azure-native:network/v20180401:getNetworkSecurityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkSecurityGroupArgs struct {
	Expand                   *string `pulumi:"expand"`
	NetworkSecurityGroupName string  `pulumi:"networkSecurityGroupName"`
	ResourceGroupName        string  `pulumi:"resourceGroupName"`
}


type LookupNetworkSecurityGroupResult struct {
	DefaultSecurityRules []SecurityRuleResponse     `pulumi:"defaultSecurityRules"`
	Etag                 *string                    `pulumi:"etag"`
	Id                   *string                    `pulumi:"id"`
	Location             *string                    `pulumi:"location"`
	Name                 string                     `pulumi:"name"`
	NetworkInterfaces    []NetworkInterfaceResponse `pulumi:"networkInterfaces"`
	ProvisioningState    *string                    `pulumi:"provisioningState"`
	ResourceGuid         *string                    `pulumi:"resourceGuid"`
	SecurityRules        []SecurityRuleResponse     `pulumi:"securityRules"`
	Subnets              []SubnetResponse           `pulumi:"subnets"`
	Tags                 map[string]string          `pulumi:"tags"`
	Type                 string                     `pulumi:"type"`
}
