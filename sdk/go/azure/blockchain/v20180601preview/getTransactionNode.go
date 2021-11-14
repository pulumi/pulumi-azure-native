


package v20180601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTransactionNode(ctx *pulumi.Context, args *LookupTransactionNodeArgs, opts ...pulumi.InvokeOption) (*LookupTransactionNodeResult, error) {
	var rv LookupTransactionNodeResult
	err := ctx.Invoke("azure-native:blockchain/v20180601preview:getTransactionNode", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTransactionNodeArgs struct {
	BlockchainMemberName string `pulumi:"blockchainMemberName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
	TransactionNodeName  string `pulumi:"transactionNodeName"`
}


type LookupTransactionNodeResult struct {
	Dns               string                 `pulumi:"dns"`
	FirewallRules     []FirewallRuleResponse `pulumi:"firewallRules"`
	Id                string                 `pulumi:"id"`
	Location          *string                `pulumi:"location"`
	Name              string                 `pulumi:"name"`
	Password          *string                `pulumi:"password"`
	ProvisioningState string                 `pulumi:"provisioningState"`
	PublicKey         string                 `pulumi:"publicKey"`
	Type              string                 `pulumi:"type"`
	UserName          string                 `pulumi:"userName"`
}
