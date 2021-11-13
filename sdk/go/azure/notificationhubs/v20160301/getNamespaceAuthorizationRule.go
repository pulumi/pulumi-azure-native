


package v20160301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNamespaceAuthorizationRule(ctx *pulumi.Context, args *LookupNamespaceAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceAuthorizationRuleResult, error) {
	var rv LookupNamespaceAuthorizationRuleResult
	err := ctx.Invoke("azure-native:notificationhubs/v20160301:getNamespaceAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamespaceAuthorizationRuleArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	NamespaceName         string `pulumi:"namespaceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupNamespaceAuthorizationRuleResult struct {
	Id       string            `pulumi:"id"`
	Location string            `pulumi:"location"`
	Name     string            `pulumi:"name"`
	Rights   []string          `pulumi:"rights"`
	Sku      *SkuResponse      `pulumi:"sku"`
	Tags     map[string]string `pulumi:"tags"`
	Type     string            `pulumi:"type"`
}
