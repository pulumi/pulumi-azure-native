


package network

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAdminRuleCollection(ctx *pulumi.Context, args *LookupAdminRuleCollectionArgs, opts ...pulumi.InvokeOption) (*LookupAdminRuleCollectionResult, error) {
	var rv LookupAdminRuleCollectionResult
	err := ctx.Invoke("azure-native:network:getAdminRuleCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAdminRuleCollectionArgs struct {
	ConfigurationName  string `pulumi:"configurationName"`
	NetworkManagerName string `pulumi:"networkManagerName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	RuleCollectionName string `pulumi:"ruleCollectionName"`
}


type LookupAdminRuleCollectionResult struct {
	AppliesToGroups   []NetworkManagerSecurityGroupItemResponse `pulumi:"appliesToGroups"`
	Description       *string                                   `pulumi:"description"`
	DisplayName       *string                                   `pulumi:"displayName"`
	Etag              string                                    `pulumi:"etag"`
	Id                string                                    `pulumi:"id"`
	Name              string                                    `pulumi:"name"`
	ProvisioningState string                                    `pulumi:"provisioningState"`
	SystemData        SystemDataResponse                        `pulumi:"systemData"`
	Type              string                                    `pulumi:"type"`
}
