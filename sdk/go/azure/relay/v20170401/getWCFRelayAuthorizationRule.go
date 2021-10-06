


package v20170401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWCFRelayAuthorizationRule(ctx *pulumi.Context, args *LookupWCFRelayAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupWCFRelayAuthorizationRuleResult, error) {
	var rv LookupWCFRelayAuthorizationRuleResult
	err := ctx.Invoke("azure-native:relay/v20170401:getWCFRelayAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWCFRelayAuthorizationRuleArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	NamespaceName         string `pulumi:"namespaceName"`
	RelayName             string `pulumi:"relayName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupWCFRelayAuthorizationRuleResult struct {
	Id     string   `pulumi:"id"`
	Name   string   `pulumi:"name"`
	Rights []string `pulumi:"rights"`
	Type   string   `pulumi:"type"`
}
