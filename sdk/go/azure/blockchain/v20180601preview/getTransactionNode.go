


package v20180601preview

import (
	"context"
	"reflect"

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

func LookupTransactionNodeOutput(ctx *pulumi.Context, args LookupTransactionNodeOutputArgs, opts ...pulumi.InvokeOption) LookupTransactionNodeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTransactionNodeResult, error) {
			args := v.(LookupTransactionNodeArgs)
			r, err := LookupTransactionNode(ctx, &args, opts...)
			var s LookupTransactionNodeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTransactionNodeResultOutput)
}

type LookupTransactionNodeOutputArgs struct {
	BlockchainMemberName pulumi.StringInput `pulumi:"blockchainMemberName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
	TransactionNodeName  pulumi.StringInput `pulumi:"transactionNodeName"`
}

func (LookupTransactionNodeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTransactionNodeArgs)(nil)).Elem()
}


type LookupTransactionNodeResultOutput struct{ *pulumi.OutputState }

func (LookupTransactionNodeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTransactionNodeResult)(nil)).Elem()
}

func (o LookupTransactionNodeResultOutput) ToLookupTransactionNodeResultOutput() LookupTransactionNodeResultOutput {
	return o
}

func (o LookupTransactionNodeResultOutput) ToLookupTransactionNodeResultOutputWithContext(ctx context.Context) LookupTransactionNodeResultOutput {
	return o
}

func (o LookupTransactionNodeResultOutput) Dns() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransactionNodeResult) string { return v.Dns }).(pulumi.StringOutput)
}

func (o LookupTransactionNodeResultOutput) FirewallRules() FirewallRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupTransactionNodeResult) []FirewallRuleResponse { return v.FirewallRules }).(FirewallRuleResponseArrayOutput)
}

func (o LookupTransactionNodeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransactionNodeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTransactionNodeResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTransactionNodeResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupTransactionNodeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransactionNodeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTransactionNodeResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTransactionNodeResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupTransactionNodeResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransactionNodeResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupTransactionNodeResultOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransactionNodeResult) string { return v.PublicKey }).(pulumi.StringOutput)
}

func (o LookupTransactionNodeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransactionNodeResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupTransactionNodeResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransactionNodeResult) string { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTransactionNodeResultOutput{})
}
